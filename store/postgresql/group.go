/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package postgresql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/polarismesh/polaris/common/log"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/common/utils"
	"github.com/polarismesh/polaris/store"
	"go.uber.org/zap"
)

const (
	// IDAttribute is the name of the attribute that stores the ID of the object.
	IDAttribute string = "id"

	// NameAttribute will be used as the name of the attribute that stores the name of the object.
	NameAttribute string = "name"

	// FlagAttribute will be used as the name of the attribute that stores the flag of the object.
	FlagAttribute string = "flag"

	// GroupIDAttribute will be used as the name of the attribute that stores the group ID of the object.
	GroupIDAttribute string = "group_id"
)

var (
	groupAttribute map[string]string = map[string]string{
		"name":  "ug.name",
		"id":    "ug.id",
		"owner": "ug.owner",
	}
)

type groupStore struct {
	master *BaseDB
	slave  *BaseDB
}

// AddGroup 创建一个用户组
func (u *groupStore) AddGroup(group *model.UserGroupDetail) error {
	if group.ID == "" || group.Name == "" || group.Token == "" {
		return store.NewStatusError(store.EmptyParamsErr, fmt.Sprintf(
			"add usergroup missing some params, groupId is %s, name is %s", group.ID, group.Name))
	}

	err := RetryTransaction("addGroup", func() error {
		return u.addGroup(group)
	})

	return store.Error(err)
}

func (u *groupStore) addGroup(group *model.UserGroupDetail) error {
	tx, err := u.master.Begin()
	if err != nil {
		return err
	}

	defer func() { _ = tx.Rollback() }()

	// 先清理无效数据
	if err := cleanInValidGroup(tx, group.Name, group.Owner); err != nil {
		return store.Error(err)
	}

	addSql := "INSERT INTO user_group (id, name, owner, token, token_enable, comment, " +
		"flag, ctime, mtime) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	stmt, err := tx.Prepare(addSql)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec([]interface{}{
		group.ID,
		group.Name,
		group.Owner,
		group.Token,
		1,
		group.Comment,
		0,
		GetCurrentTimeFormat(),
		GetCurrentTimeFormat(),
	}...); err != nil {
		log.Errorf("[Store][Group] add usergroup err: %s", err.Error())
		return err
	}

	if err := u.addGroupRelation(tx, group.ID, group.ToUserIdSlice()); err != nil {
		log.Errorf("[Store][Group] add usergroup relation err: %s", err.Error())
		return err
	}

	if err := createDefaultStrategy(tx, model.PrincipalGroup, group.ID, group.Name, group.Owner); err != nil {
		log.Errorf("[Store][Group] add usergroup default strategy err: %s", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("[Store][Group] add usergroup tx commit err: %s", err.Error())
		return err
	}
	return nil
}

// UpdateGroup 更新用户组
func (u *groupStore) UpdateGroup(group *model.ModifyUserGroup) error {
	if group.ID == "" {
		return store.NewStatusError(store.EmptyParamsErr, fmt.Sprintf(
			"update usergroup missing some params, groupId is %s", group.ID))
	}

	err := RetryTransaction("updateGroup", func() error {
		return u.updateGroup(group)
	})

	return store.Error(err)
}

func (u *groupStore) updateGroup(group *model.ModifyUserGroup) error {
	tx, err := u.master.Begin()
	if err != nil {
		return err
	}

	defer func() { _ = tx.Rollback() }()

	tokenEnable := 1
	if !group.TokenEnable {
		tokenEnable = 0
	}

	// 更新用户-用户组关联数据
	if len(group.AddUserIds) != 0 {
		if err := u.addGroupRelation(tx, group.ID, group.AddUserIds); err != nil {
			log.Errorf("[Store][Group] add usergroup relation err: %s", err.Error())
			return err
		}
	}

	if len(group.RemoveUserIds) != 0 {
		if err := u.removeGroupRelation(tx, group.ID, group.RemoveUserIds); err != nil {
			log.Errorf("[Store][Group] remove usergroup relation err: %s", err.Error())
			return err
		}
	}

	modifySql := "UPDATE user_group SET token = $1, comment = $2, token_enable = $3, mtime = $4 " +
		" WHERE id = $5 AND flag = 0"
	stmt, err := tx.Prepare(modifySql)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec([]interface{}{
		group.Token,
		group.Comment,
		tokenEnable,
		GetCurrentTimeFormat(),
		group.ID,
	}...); err != nil {
		log.Errorf("[Store][Group] update usergroup main err: %s", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("[Store][Group] update usergroup tx commit err: %s", err.Error())
		return err
	}

	return nil
}

// DeleteGroup 删除用户组
func (u *groupStore) DeleteGroup(group *model.UserGroupDetail) error {
	if group.ID == "" || group.Name == "" {
		return store.NewStatusError(store.EmptyParamsErr, fmt.Sprintf(
			"delete usergroup missing some params, groupId is %s", group.ID))
	}

	err := RetryTransaction("deleteUserGroup", func() error {
		return u.deleteUserGroup(group)
	})

	return store.Error(err)
}

func (u *groupStore) deleteUserGroup(group *model.UserGroupDetail) error {
	tx, err := u.master.Begin()
	if err != nil {
		return err
	}

	defer func() { _ = tx.Rollback() }()

	stmt, err := tx.Prepare("DELETE FROM user_group_relation WHERE group_id = $1")
	if err != nil {
		return err
	}
	if _, err = stmt.Exec([]interface{}{
		group.ID,
	}...); err != nil {
		log.Errorf("[Store][Group] clean usergroup relation err: %s", err.Error())
		return err
	}

	stmt, err = tx.Prepare("UPDATE user_group SET flag = 1, mtime = $1 WHERE id = $2")
	if err != nil {
		return err
	}
	if _, err = stmt.Exec([]interface{}{
		GetCurrentTimeFormat(),
		group.ID,
	}...); err != nil {
		log.Errorf("[Store][Group] remove usergroup err: %s", err.Error())
		return err
	}

	if err := cleanLinkStrategy(tx, model.PrincipalGroup, group.ID, group.Owner); err != nil {
		log.Errorf("[Store][Group] clean usergroup default strategy err: %s", err.Error())
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Errorf("[Store][Group] delete usergroupr tx commit err: %s", err.Error())
		return err
	}
	return nil
}

// GetGroup 根据用户组ID获取用户组
func (u *groupStore) GetGroup(groupId string) (*model.UserGroupDetail, error) {
	if groupId == "" {
		return nil, store.NewStatusError(store.EmptyParamsErr, fmt.Sprintf(
			"get usergroup missing some params, groupId is %s", groupId))
	}

	getSql := "SELECT ug.id, ug.name, ug.owner, ug.comment, ug.token, " +
		"ug.token_enable, ug.ctime, ug.mtime " +
		"FROM user_group ug WHERE ug.flag = 0 AND ug.id = $1"
	row := u.master.QueryRow(getSql, groupId)

	group := &model.UserGroupDetail{
		UserGroup: &model.UserGroup{},
	}
	var (
		tokenEnable int
	)

	if err := row.Scan(&group.ID, &group.Name, &group.Owner, &group.Comment, &group.Token,
		&tokenEnable, &group.CreateTime, &group.ModifyTime); err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, store.Error(err)
		}
	}
	uids, err := u.getGroupLinkUserIds(group.ID)
	if err != nil {
		return nil, store.Error(err)
	}

	group.UserIds = uids
	group.TokenEnable = tokenEnable == 1

	return group, nil
}

// GetGroupByName 根据 owner、name 获取用户组
func (u *groupStore) GetGroupByName(name, owner string) (*model.UserGroup, error) {
	if name == "" || owner == "" {
		return nil, store.NewStatusError(store.EmptyParamsErr, fmt.Sprintf(
			"get usergroup missing some params, name=%s, owner=%s", name, owner))
	}

	getSql := "SELECT ug.id, ug.name, ug.owner, ug.comment, ug.token, " +
		"ug.ctime, ug.mtime FROM user_group ug " +
		"WHERE ug.flag = 0 AND ug.name = $1 AND ug.owner = $2"
	row := u.master.QueryRow(getSql, name, owner)

	group := new(model.UserGroup)

	if err := row.Scan(&group.ID, &group.Name, &group.Owner, &group.Comment, &group.Token,
		&group.CreateTime, &group.ModifyTime); err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, store.Error(err)
		}
	}

	return group, nil
}

// GetGroups 根据不同的请求情况进行不同的用户组列表查询
func (u *groupStore) GetGroups(filters map[string]string, offset uint32, limit uint32) (uint32,
	[]*model.UserGroup, error) {

	// 如果本次请求参数携带了 user_id，那么就是查询这个用户所关联的所有用户组
	if _, ok := filters["user_id"]; ok {
		return u.listGroupByUser(filters, offset, limit)
	}
	// 正常查询用户组信息
	return u.listSimpleGroups(filters, offset, limit)
}

// listSimpleGroups 正常的用户组查询
func (u *groupStore) listSimpleGroups(filters map[string]string, offset uint32, limit uint32) (uint32,
	[]*model.UserGroup, error) {

	query := make(map[string]string)
	if _, ok := filters["id"]; ok {
		query["id"] = filters["id"]
	}
	if _, ok := filters["name"]; ok {
		query["name"] = filters["name"]
	}
	filters = query

	countSql := "SELECT COUNT(*) FROM user_group ug WHERE ug.flag = 0 "
	getSql := "SELECT ug.id, ug.name, ug.owner, ug.comment, ug.token, " +
		"ug.token_enable, ug.ctime, ug.mtime, ug.flag " +
		"FROM user_group ug WHERE ug.flag = 0"

	args := make([]interface{}, 0)
	idx := 1

	if len(filters) != 0 {
		for k, v := range filters {
			getSql += " AND "
			countSql += " AND "
			if newK, ok := groupAttribute[k]; ok {
				k = newK
			}
			if utils.IsPrefixWildName(v) {
				getSql += fmt.Sprintf(" "+k+" like $%d ", idx)
				countSql += fmt.Sprintf(" "+k+" like $%d ", idx)
				args = append(args, "%"+v[:len(v)-1]+"%")
			} else {
				getSql += " " + k + fmt.Sprintf(" = $%d ", idx)
				countSql += " " + k + fmt.Sprintf(" = $%d ", idx)
				args = append(args, v)
			}
			idx++
		}
	}

	count, err := queryEntryCount(u.master, countSql, args)
	if err != nil {
		return 0, nil, err
	}

	getSql += fmt.Sprintf(" ORDER BY ug.mtime LIMIT $%d OFFSET $%d", idx, idx+1)
	args = append(args, limit, offset)

	groups, err := u.collectGroupsFromRows(u.master.Query, getSql, args)
	if err != nil {
		return 0, nil, err
	}

	return count, groups, nil
}

// listGroupByUser 查询某个用户下所关联的用户组信息
func (u *groupStore) listGroupByUser(filters map[string]string, offset uint32, limit uint32) (uint32,
	[]*model.UserGroup, error) {
	countSql := "SELECT COUNT(*) FROM user_group_relation ul LEFT JOIN user_group ug ON " +
		" ul.group_id = ug.id WHERE ug.flag = 0 "
	getSql := "SELECT ug.id, ug.name, ug.owner, ug.comment, ug.token, ug.token_enable, ug.ctime, " +
		" ug.mtime, ug.flag " +
		" FROM user_group_relation ul LEFT JOIN user_group ug ON ul.group_id = ug.id WHERE ug.flag = 0 "

	args := make([]interface{}, 0)
	idx := 1

	if len(filters) != 0 {
		for k, v := range filters {
			getSql += " AND "
			countSql += " AND "
			if newK, ok := userLinkGroupAttributeMapping[k]; ok {
				k = newK
			}
			if utils.IsPrefixWildName(v) {
				getSql += " " + k + fmt.Sprintf(" like $%d ", idx)
				countSql += " " + k + fmt.Sprintf(" like $%d ", idx)
				args = append(args, "%"+v[:len(v)-1]+"%")
			} else if k == "ug.owner" {
				getSql += fmt.Sprintf(" (ug.owner = $%d OR ul.user_id = $%d ) ", idx, idx+1)
				countSql += fmt.Sprintf(" (ug.owner = $%d OR ul.user_id = $%d ) ", idx, idx+1)
				idx += 1
				args = append(args, v, v)
			} else {
				getSql += " " + k + fmt.Sprintf(" = $%d ", idx)
				countSql += fmt.Sprintf(" "+k+" = $%d ", idx)
				args = append(args, v)
			}
			idx++
		}
	}

	count, err := queryEntryCount(u.master, countSql, args)
	if err != nil {
		return 0, nil, err
	}

	getSql += fmt.Sprintf(" GROUP BY ug.id ORDER BY ug.mtime LIMIT $%d OFFSET $%d", idx, idx+1)
	args = append(args, limit, offset)

	groups, err := u.collectGroupsFromRows(u.master.Query, getSql, args)
	if err != nil {
		return 0, nil, err
	}

	return count, groups, nil
}

// collectGroupsFromRows 查询用户组列表
func (u *groupStore) collectGroupsFromRows(handler QueryHandler, querySql string,
	args []interface{}) ([]*model.UserGroup, error) {
	rows, err := u.master.Query(querySql, args...)
	if err != nil {
		log.Error("[Store][Group] list group", zap.String("query sql", querySql), zap.Any("args", args))
		return nil, err
	}
	defer rows.Close()

	groups := make([]*model.UserGroup, 0)
	for rows.Next() {
		group, err := fetchRown2UserGroup(rows)
		if err != nil {
			log.Errorf("[Store][Group] list group by user fetch rows scan err: %s", err.Error())
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

// GetGroupsForCache .
func (u *groupStore) GetGroupsForCache(mtime time.Time, firstUpdate bool) ([]*model.UserGroupDetail, error) {
	tx, err := u.slave.Begin()
	if err != nil {
		return nil, store.Error(err)
	}

	defer func() { _ = tx.Commit() }()

	args := make([]interface{}, 0)
	querySql := "SELECT id, name, owner, comment, token, token_enable, ctime, mtime, " +
		" flag FROM user_group "
	if !firstUpdate {
		querySql += " WHERE mtime >= $1"
		args = append(args, mtime)
	}

	rows, err := tx.Query(querySql, args...)
	if err != nil {
		return nil, store.Error(err)
	}
	defer rows.Close()

	ret := make([]*model.UserGroupDetail, 0)
	for rows.Next() {
		detail := &model.UserGroupDetail{
			UserIds: make(map[string]struct{}, 0),
		}
		group, err := fetchRown2UserGroup(rows)
		if err != nil {
			return nil, store.Error(err)
		}
		uids, err := u.getGroupLinkUserIds(group.ID)
		if err != nil {
			return nil, store.Error(err)
		}

		detail.UserIds = uids
		detail.UserGroup = group

		ret = append(ret, detail)
	}

	return ret, nil
}

func (u *groupStore) addGroupRelation(tx *BaseTx, groupId string, userIds []string) error {
	if groupId == "" {
		return store.NewStatusError(store.EmptyParamsErr, fmt.Sprintf(
			"add user relation missing some params, groupid is %s", groupId))
	}
	if len(userIds) > utils.MaxBatchSize {
		return store.NewStatusError(store.InvalidUserIDSlice, fmt.Sprintf(
			"user id slice is invalid, len=%d", len(userIds)))
	}

	for i := range userIds {
		uid := userIds[i]
		addSql := "INSERT INTO user_group_relation (group_id, user_id) VALUE ($1,$2)"
		args := []interface{}{groupId, uid}
		stmt, err := tx.Prepare(addSql)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(args...)
		if err != nil {
			err = store.Error(err)
			// 之前的用户已经存在，直接忽略
			if store.Code(err) == store.DuplicateEntryErr {
				continue
			}
			return err
		}
	}
	return nil
}

func (u *groupStore) removeGroupRelation(tx *BaseTx, groupId string, userIds []string) error {
	if groupId == "" {
		return store.NewStatusError(store.EmptyParamsErr, fmt.Sprintf(
			"delete user relation missing some params, groupid is %s", groupId))
	}
	if len(userIds) > utils.MaxBatchSize {
		return store.NewStatusError(store.InvalidUserIDSlice, fmt.Sprintf(
			"user id slice is invalid, len=%d", len(userIds)))
	}

	for i := range userIds {
		uid := userIds[i]
		addSql := "DELETE FROM user_group_relation WHERE group_id = $1 AND user_id = $2"
		args := []interface{}{groupId, uid}
		stmt, err := tx.Prepare(addSql)
		if err != nil {
			return err
		}
		if _, err := stmt.Exec(args...); err != nil {
			return err
		}
	}

	return nil
}

func (u *groupStore) getGroupLinkUserIds(groupId string) (map[string]struct{}, error) {

	ids := make(map[string]struct{})

	// 拉取该分组下的所有 user
	idRows, err := u.slave.Query("SELECT user_id FROM user u JOIN user_group_relation ug ON "+
		" u.id = ug.user_id WHERE ug.group_id = $1", groupId)
	if err != nil {
		return nil, err
	}
	defer idRows.Close()
	for idRows.Next() {
		var uid string
		if err := idRows.Scan(&uid); err != nil {
			return nil, err
		}
		ids[uid] = struct{}{}
	}

	return ids, nil
}

func fetchRown2UserGroup(rows *sql.Rows) (*model.UserGroup, error) {
	var flag, tokenEnable int
	group := new(model.UserGroup)
	if err := rows.Scan(&group.ID, &group.Name, &group.Owner, &group.Comment, &group.Token,
		&tokenEnable, &group.CreateTime, &group.ModifyTime, &flag); err != nil {
		return nil, err
	}

	group.Valid = flag == 0
	group.TokenEnable = tokenEnable == 1

	return group, nil
}

// cleanInValidUserGroup 清理无效的用户组数据
func cleanInValidGroup(tx *BaseTx, name, owner string) error {
	log.Infof("[Store][User] clean usergroup(%s)", name)

	str := "delete from user_group where name = $1 and flag = 1"
	stmt, err := tx.Prepare(str)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(name); err != nil {
		log.Errorf("[Store][User] clean usergroup(%s) err: %s", name, err.Error())
		return err
	}

	return nil
}
