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

	"github.com/polarismesh/polaris/common/log"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
)

var _ store.ConfigFileStore = (*configFileStore)(nil)

type configFileStore struct {
	master *BaseDB
	slave  *BaseDB
}

// CreateConfigFile 创建配置文件
func (cf *configFileStore) CreateConfigFile(tx store.Tx, file *model.ConfigFile) (*model.ConfigFile, error) {
	err := cf.hardDeleteConfigFile(file.Namespace, file.Group, file.Name)
	if err != nil {
		return nil, err
	}
	createSql := "insert into config_file(name,namespace,group,content,comment,format,create_time, " +
		"create_by,modify_time,modify_by) values " +
		"($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"
	if tx != nil {
		stmt, err := tx.GetDelegateTx().(*BaseTx).Prepare(createSql)
		if err != nil {
			return nil, store.Error(err)
		}
		_, err = stmt.Exec(file.Name, file.Namespace, file.Group, file.Content, file.Comment,
			file.Format, GetCurrentTimeFormat(), file.CreateBy, GetCurrentTimeFormat(),
			file.ModifyBy)
	} else {
		stmt, err := cf.master.Prepare(createSql)
		if err != nil {
			return nil, store.Error(err)
		}
		_, err = stmt.Exec(file.Name, file.Namespace, file.Group, file.Content, file.Comment,
			file.Format, GetCurrentTimeFormat(), file.CreateBy, GetCurrentTimeFormat(),
			file.ModifyBy)
	}
	if err != nil {
		return nil, store.Error(err)
	}
	return cf.GetConfigFile(tx, file.Namespace, file.Group, file.Name)
}

// GetConfigFile 获取配置文件
func (cf *configFileStore) GetConfigFile(tx store.Tx, namespace, group, name string) (*model.ConfigFile, error) {
	querySql := cf.baseSelectConfigFileSql() + "where namespace = $1 and `group` = $2 and name = $3 and flag = 0"
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.GetDelegateTx().(*BaseTx).Query(querySql, namespace, group, name)
	} else {
		rows, err = cf.master.Query(querySql, namespace, group, name)
	}
	if err != nil {
		return nil, err
	}
	files, err := cf.transferRows(rows)
	if err != nil {
		return nil, err
	}
	if len(files) > 0 {
		return files[0], nil
	}

	return nil, nil
}

func (cf *configFileStore) QueryConfigFilesByGroup(namespace, group string,
	offset, limit uint32) (uint32, []*model.ConfigFile, error) {
	var (
		countSql = "select count(*) from config_file where namespace = $1 and group = $2 and flag = 0"
		count    uint32
		err      = cf.master.QueryRow(countSql, namespace, group).Scan(&count)
	)

	if err != nil {
		return 0, nil, err
	}

	querySql := cf.baseSelectConfigFileSql() + "where namespace = $1 and `group` = $2 and flag = 0 order by id " +
		" desc limit $3 offset $4"
	rows, err := cf.master.Query(querySql, namespace, group, limit, offset)
	if err != nil {
		return 0, nil, err
	}

	files, err := cf.transferRows(rows)
	if err != nil {
		return 0, nil, err
	}

	return count, files, nil
}

// QueryConfigFiles 翻页查询配置文件，group、name可为模糊匹配
func (cf *configFileStore) QueryConfigFiles(namespace, group, name string,
	offset, limit uint32) (uint32, []*model.ConfigFile, error) {
	// 全部 namespace
	if namespace == "" {
		group = "%" + group + "%"
		name = "%" + name + "%"
		countSql := "select count(*) from config_file where group like $1 and name like $2 and flag = 0"

		var count uint32
		err := cf.master.QueryRow(countSql, group, name).Scan(&count)
		if err != nil {
			return 0, nil, err
		}

		querySql := cf.baseSelectConfigFileSql() + "where group like $1 and name like $2 and flag = 0 " +
			" order by id desc limit $3 offset $4"
		rows, err := cf.master.Query(querySql, group, name, limit, offset)
		if err != nil {
			return 0, nil, err
		}

		files, err := cf.transferRows(rows)
		if err != nil {
			return 0, nil, err
		}

		return count, files, nil
	}

	// 特定 namespace
	group = "%" + group + "%"
	name = "%" + name + "%"
	countSql := "select count(*) from config_file where namespace = $1 and group like $2 and name like $3 and flag = 0"

	var count uint32
	err := cf.master.QueryRow(countSql, namespace, group, name).Scan(&count)
	if err != nil {
		return 0, nil, err
	}

	querySql := cf.baseSelectConfigFileSql() + "where namespace = $1 and group like $2 and name like $3 " +
		" and flag = 0 order by id desc limit $4 offset $5"
	rows, err := cf.master.Query(querySql, namespace, group, name, offset, limit)
	if err != nil {
		return 0, nil, err
	}

	files, err := cf.transferRows(rows)
	if err != nil {
		return 0, nil, err
	}

	return count, files, nil
}

// UpdateConfigFile 更新配置文件
func (cf *configFileStore) UpdateConfigFile(tx store.Tx, file *model.ConfigFile) (*model.ConfigFile, error) {
	updateSql := "update config_file set content = $1 , comment = $2, format = $3, modify_time = $4, " +
		" modify_by = $5 where namespace = $6 and group = $7 and name = $8"
	var err error
	if tx != nil {
		stmt, err := tx.GetDelegateTx().(*BaseTx).Prepare(updateSql)
		if err != nil {
			return nil, store.Error(err)
		}
		_, err = stmt.Exec(file.Content, file.Comment, file.Format, GetCurrentTimeFormat(),
			file.ModifyBy, file.Namespace, file.Group, file.Name)
	} else {
		stmt, err := cf.master.Prepare(updateSql)
		if err != nil {
			return nil, store.Error(err)
		}
		_, err = stmt.Exec(file.Content, file.Comment, file.Format, GetCurrentTimeFormat(),
			file.ModifyBy, file.Namespace, file.Group, file.Name)
	}
	if err != nil {
		return nil, store.Error(err)
	}
	return cf.GetConfigFile(tx, file.Namespace, file.Group, file.Name)
}

// DeleteConfigFile 删除配置文件
func (cf *configFileStore) DeleteConfigFile(tx store.Tx, namespace, group, name string) error {
	deleteSql := "update config_file set flag = 1 where namespace = $1 and group = $2 and name = $3"
	var err error
	if tx != nil {
		stmt, err := tx.GetDelegateTx().(*BaseTx).Prepare(deleteSql)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(namespace, group, name)
	} else {
		stmt, err := cf.master.Prepare(deleteSql)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(namespace, group, name)
	}
	if err != nil {
		return store.Error(err)
	}
	return nil
}

func (cf *configFileStore) CountByConfigFileGroup(namespace, group string) (uint64, error) {
	countSql := "select count(*) from config_file where namespace = $1 and group = $2 and flag = 0"
	var count uint64
	err := cf.master.QueryRow(countSql, namespace, group).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (cf *configFileStore) CountConfigFileEachGroup() (map[string]map[string]int64, error) {
	metricsSql := "SELECT namespace, group, count(name) FROM config_file WHERE flag = 0 GROUP by namespace, group"
	rows, err := cf.slave.Query(metricsSql)
	if err != nil {
		return nil, store.Error(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	ret := map[string]map[string]int64{}
	for rows.Next() {
		var (
			namespce string
			group    string
			cnt      int64
		)

		if err := rows.Scan(&namespce, &group, &cnt); err != nil {
			return nil, err
		}
		if _, ok := ret[namespce]; !ok {
			ret[namespce] = map[string]int64{}
		}
		ret[namespce][group] = cnt
	}

	return ret, nil
}

func (cf *configFileStore) baseSelectConfigFileSql() string {
	return "select id, name,namespace,group,content,comment,format, create_time, " +
		" create_by,modify_time,modify_by from config_file "
}

func (cf *configFileStore) hardDeleteConfigFile(namespace, group, name string) error {
	log.Infof("[Config][Storage] delete config file. namespace = %s, group = %s, name = %s", namespace, group, name)

	deleteSql := "delete from config_file where namespace = $1 and group = $2 and name = $3 and flag = 1"
	stmt, err := cf.master.Prepare(deleteSql)
	if err != nil {
		return store.Error(err)
	}
	_, err = stmt.Exec(namespace, group, name)
	if err != nil {
		return store.Error(err)
	}

	return nil
}

func (cf *configFileStore) transferRows(rows *sql.Rows) ([]*model.ConfigFile, error) {
	if rows == nil {
		return nil, nil
	}
	defer rows.Close()

	var files []*model.ConfigFile

	for rows.Next() {
		file := &model.ConfigFile{}
		err := rows.Scan(&file.Id, &file.Name, &file.Namespace, &file.Group, &file.Content, &file.Comment,
			&file.Format, &file.CreateTime, &file.CreateBy, &file.ModifyTime, &file.ModifyBy)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}
