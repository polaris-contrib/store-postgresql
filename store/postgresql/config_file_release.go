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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/polarismesh/polaris/common/utils"
	"strconv"
	"time"

	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
)

var _ store.ConfigFileReleaseStore = (*configFileReleaseStore)(nil)

var (
	ErrTxIsNil = errors.New("tx is nil")
)

type configFileReleaseStore struct {
	master *BaseDB
	slave  *BaseDB
}

// CreateConfigFileReleaseTx 新建配置文件发布
func (cfr *configFileReleaseStore) CreateConfigFileReleaseTx(tx store.Tx, data *model.ConfigFileRelease) error {
	if tx == nil {
		return ErrTxIsNil
	}
	dbTx := tx.GetDelegateTx().(*BaseTx)

	// 使用 defer 处理回滚逻辑
	var err error
	defer func() {
		if err != nil {
			_ = dbTx.Rollback()
		}
		if r := recover(); r != nil {
			_ = dbTx.Rollback()
			panic(r)
		}
	}()

	// 使用 PostgreSQL 的语法进行查询
	args := []interface{}{data.Namespace, data.Group, data.Name}
	_, err = dbTx.Exec("SELECT id FROM config_file WHERE namespace = $1 AND \"group\" = $2 AND name = $3 FOR UPDATE", args...)
	if err != nil {
		return store.Error(err)
	}

	// 删除旧的配置文件发布
	clean := "DELETE FROM config_file_release WHERE namespace = $1 AND \"group\" = $2 AND file_name = $3 AND name = $4 AND flag = 1"
	if _, err := dbTx.Exec(clean, data.Namespace, data.Group, data.FileName, data.Name); err != nil {
		return store.Error(err)
	}

	// 使旧版本无效
	maxVersion, err := cfr.inactiveConfigFileRelease(dbTx, data)
	if err != nil {
		return store.Error(err)
	}

	// 插入新的配置文件发布
	s := `
		INSERT INTO config_file_release (
			name, namespace, "group", file_name, content, comment, md5, 
			version, create_time, create_by, modify_time, modify_by, 
			active, tags, description, release_type
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, 
				$8, CURRENT_TIMESTAMP, $9, CURRENT_TIMESTAMP, $10, 
				1, $11, $12, $13)
	`
	args = []interface{}{
		data.Name, data.Namespace, data.Group,
		data.FileName, data.Content, data.Comment, data.Md5, maxVersion + 1,
		data.CreateBy, data.ModifyBy, utils.MustJson(data.Metadata), data.ReleaseDescription, data.ReleaseType,
	}

	if _, err = dbTx.Exec(s, args...); err != nil {
		return store.Error(err)
	}

	if err = dbTx.Commit(); err != nil {
		return store.Error(err)
	}

	return nil
}

// GetConfigFileRelease 获取配置文件发布，只返回 flag=0 的记录
func (cfr *configFileReleaseStore) GetConfigFileRelease(req *model.ConfigFileReleaseKey) (*model.ConfigFileRelease, error) {
	tx, err := cfr.master.Begin()
	if err != nil {
		return nil, store.Error(err)
	}
	defer func() {
		_ = tx.Rollback()
	}()
	return cfr.GetConfigFileReleaseTx(NewSqlDBTx(tx), req)
}

// GetConfigFileReleaseTx 在已开启的事务中获取配置文件发布内容，只获取 flag=0 的记录
func (cfr *configFileReleaseStore) GetConfigFileReleaseTx(tx store.Tx,
	req *model.ConfigFileReleaseKey) (*model.ConfigFileRelease, error) {
	if tx == nil {
		return nil, ErrTxIsNil
	}

	dbTx := tx.GetDelegateTx().(*BaseTx)
	// 使用 PostgreSQL 语法构建查询
	querySql := cfr.baseQuerySql() + " WHERE namespace = $1 AND \"group\" = $2 AND file_name = $3 AND name = $4 AND flag = 0"

	var (
		rows *sql.Rows
		err  error
	)

	// 使用 PostgreSQL 的占位符
	rows, err = dbTx.Query(querySql, req.Namespace, req.Group, req.FileName, req.Name)
	if err != nil {
		return nil, err
	}

	fileRelease, err := cfr.transferRows(rows)
	if err != nil {
		return nil, err
	}
	if len(fileRelease) > 0 {
		return fileRelease[0], nil
	}
	return nil, nil
}

func (cfr *configFileReleaseStore) DeleteConfigFileReleaseTx(tx store.Tx, data *model.ConfigFileReleaseKey) error {
	if tx == nil {
		return ErrTxIsNil
	}

	dbTx := tx.GetDelegateTx().(*BaseTx)

	// 使用 defer 处理回滚逻辑
	var err error
	defer func() {
		if err != nil {
			_ = dbTx.Rollback()
		}
		if r := recover(); r != nil {
			_ = dbTx.Rollback()
			panic(r)
		}
	}()

	s := "UPDATE config_file_release SET flag = 1, modify_time = CURRENT_TIMESTAMP " +
		"WHERE namespace = $1 AND \"group\" = $2 AND file_name = $3 AND name = $4"
	_, err = dbTx.Exec(s, data.Namespace, data.Group, data.FileName, data.Name)
	if err != nil {
		return store.Error(err)
	}

	if err = dbTx.Commit(); err != nil {
		return store.Error(err)
	}

	return nil
}

func (cfr *configFileReleaseStore) CleanConfigFileReleasesTx(tx store.Tx,
	namespace, group, fileName string) error {
	if tx == nil {
		return ErrTxIsNil
	}

	dbTx := tx.GetDelegateTx().(*BaseTx)

	// 使用 defer 处理回滚逻辑
	var err error
	defer func() {
		if err != nil {
			_ = dbTx.Rollback()
		}
		if r := recover(); r != nil {
			_ = dbTx.Rollback()
			panic(r)
		}
	}()

	s := "UPDATE config_file_release SET flag = 1, modify_time = CURRENT_TIMESTAMP WHERE namespace = $1 " +
		"AND \"group\" = $2 AND file_name = $3"
	_, err = dbTx.Exec(s, namespace, group, fileName)
	if err != nil {
		return store.Error(err)
	}

	if err = dbTx.Commit(); err != nil {
		return store.Error(err)
	}

	return nil
}

// CleanDeletedConfigFileRelease 清理配置发布历史
func (cfr *configFileReleaseStore) CleanDeletedConfigFileRelease(endTime time.Time, limit uint64) error {
	delSql := `
		DELETE FROM config_file_release 
		WHERE ctid IN (
			SELECT ctid FROM config_file_release 
			WHERE modify_time < $1 AND flag = 1 
			LIMIT $2
		)`
	_, err := cfr.master.Exec(delSql, endTime, limit)
	return err
}

func (cfr *configFileReleaseStore) GetConfigFileActiveRelease(file *model.ConfigFileKey) (*model.ConfigFileRelease, error) {
	tx, err := cfr.master.Begin()
	if err != nil {
		return nil, store.Error(err)
	}
	defer func() {
		_ = tx.Rollback()
	}()
	return cfr.GetConfigFileActiveReleaseTx(NewSqlDBTx(tx), file)
}

// GetConfigFileActiveReleaseTx retrieves the active release of a configuration file within a transaction.
func (cfr *configFileReleaseStore) GetConfigFileActiveReleaseTx(tx store.Tx,
	file *model.ConfigFileKey) (*model.ConfigFileRelease, error) {
	if tx == nil {
		return nil, ErrTxIsNil
	}

	dbTx := tx.GetDelegateTx().(*BaseTx)
	querySql := cfr.baseQuerySql() + "WHERE namespace = $1 AND \"group\" = $2 AND " +
		"file_name = $3 AND active = 1 AND release_type = $4 AND flag = 0 "

	var (
		rows *sql.Rows
		err  error
	)

	rows, err = dbTx.Query(querySql, file.Namespace, file.Group, file.Name, model.ReleaseTypeFull)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Close the rows after processing

	fileRelease, err := cfr.transferRows(rows)
	if err != nil {
		return nil, err
	}
	if len(fileRelease) > 1 {
		return nil, errors.New("multiple active file releases found")
	}
	if len(fileRelease) > 0 {
		return fileRelease[0], nil
	}
	return nil, nil
}

// GetConfigFileBetaReleaseTx retrieves the beta release of a configuration file within a transaction.
func (cfr *configFileReleaseStore) GetConfigFileBetaReleaseTx(tx store.Tx,
	file *model.ConfigFileKey) (*model.ConfigFileRelease, error) {
	if tx == nil {
		return nil, ErrTxIsNil
	}

	dbTx := tx.GetDelegateTx().(*BaseTx)
	querySql := cfr.baseQuerySql() + "WHERE namespace = $1 AND \"group\" = $2 AND " +
		"file_name = $3 AND active = 1 AND release_type = $4 AND flag = 0 "

	var (
		rows *sql.Rows
		err  error
	)

	rows, err = dbTx.Query(querySql, file.Namespace, file.Group, file.Name, model.ReleaseTypeGray)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Ensure rows are closed after processing

	fileRelease, err := cfr.transferRows(rows)
	if err != nil {
		return nil, err
	}
	if len(fileRelease) > 1 {
		return nil, errors.New("multiple active file releases found")
	}
	if len(fileRelease) > 0 {
		return fileRelease[0], nil
	}
	return nil, nil
}

// ActiveConfigFileReleaseTx activates the specified config file release within a transaction.
func (cfr *configFileReleaseStore) ActiveConfigFileReleaseTx(tx store.Tx, release *model.ConfigFileRelease) error {
	if tx == nil {
		return ErrTxIsNil
	}

	dbTx := tx.GetDelegateTx().(*BaseTx)

	var err error
	defer func() {
		if err != nil {
			_ = dbTx.Rollback()
		}
		if r := recover(); r != nil {
			_ = dbTx.Rollback()
			panic(r)
		}
	}()

	maxVersion, err := cfr.inactiveConfigFileRelease(dbTx, release)
	if err != nil {
		return err
	}

	args := []interface{}{
		maxVersion + 1, // Version incremented by 1
		release.ReleaseType,
		release.Namespace,
		release.Group,
		release.FileName,
		release.Name,
	}

	// Update the specified release record, setting its active status, version, and modify time
	updateSql := "UPDATE config_file_release SET active = 1, version = $1, modify_time = CURRENT_TIMESTAMP, release_type = $2 " +
		"WHERE namespace = $3 AND \"group\" = $4 AND file_name = $5 AND name = $6"
	if _, err := dbTx.Exec(updateSql, args...); err != nil {
		return store.Error(err)
	}

	if err = dbTx.Commit(); err != nil {
		return store.Error(err)
	}

	return nil
}

// InactiveConfigFileReleaseTx deactivates the specified config file release within a transaction.
func (cfr *configFileReleaseStore) InactiveConfigFileReleaseTx(tx store.Tx, release *model.ConfigFileRelease) error {
	if tx == nil {
		return ErrTxIsNil
	}
	dbTx := tx.GetDelegateTx().(*BaseTx)

	var err error
	defer func() {
		if err != nil {
			_ = dbTx.Rollback()
		}
		if r := recover(); r != nil {
			_ = dbTx.Rollback()
			panic(r)
		}
	}()

	args := []interface{}{
		release.Namespace,
		release.Group,
		release.FileName,
		release.Name,
		release.ReleaseType,
	}

	// Deactivate the corresponding release version
	if _, err := dbTx.Exec("UPDATE config_file_release SET active = 0, modify_time = CURRENT_TIMESTAMP "+
		"WHERE namespace = $1 AND \"group\" = $2 AND file_name = $3 AND name = $4 AND release_type = $5", args...); err != nil {
		return store.Error(err)
	}

	if err = dbTx.Commit(); err != nil {
		return store.Error(err)
	}

	return nil
}

func (cfr *configFileReleaseStore) inactiveConfigFileRelease(tx *BaseTx,
	release *model.ConfigFileRelease) (uint64, error) {
	if tx == nil {
		return 0, ErrTxIsNil
	}

	args := []interface{}{release.Namespace, release.Group, release.FileName, release.ReleaseType}
	// Deactivate all records with active == true
	if _, err := tx.Exec("UPDATE config_file_release SET active = 0, modify_time = CURRENT_TIMESTAMP "+
		"WHERE namespace = $1 AND \"group\" = $2 AND file_name = $3 AND active = 1 AND release_type = $4", args...); err != nil {
		return 0, err
	}
	return cfr.selectMaxVersion(tx, release)
}

func (cfr *configFileReleaseStore) selectMaxVersion(tx *BaseTx, release *model.ConfigFileRelease) (uint64, error) {
	if tx == nil {
		return 0, ErrTxIsNil
	}

	args := []interface{}{release.Namespace, release.Group, release.FileName}
	// Generate the latest version information
	row := tx.QueryRow("SELECT COALESCE(MAX(version), 0) FROM config_file_release WHERE namespace = $1 AND "+
		"\"group\" = $2 AND file_name = $3", args...)
	var maxVersion uint64
	if err := row.Scan(&maxVersion); err != nil {
		return 0, err
	}
	return maxVersion, nil
}

// GetMoreReleaseFile 获取最近更新的配置文件发布, 此方法用于 cache 增量更新，需要注意 modifyTime 应为数据库时间戳
func (cfr *configFileReleaseStore) GetMoreReleaseFile(firstUpdate bool, modifyTime time.Time) ([]*model.ConfigFileRelease, error) {
	if firstUpdate {
		modifyTime = time.Time{} // 设定为零时间
	}

	// 使用 PostgreSQL 的时间比较
	s := cfr.baseQuerySql() + " WHERE modify_time > $1"
	rows, err := cfr.slave.Query(s, modifyTime) // 直接传入 time.Time 类型
	if err != nil {
		return nil, err
	}

	releases, err := cfr.transferRows(rows)
	if err != nil {
		return nil, err
	}
	return releases, nil
}

// CountConfigReleases 获取一个配置文件组下的文件数量
func (cfr *configFileReleaseStore) CountConfigReleases(namespace, group string, onlyActive bool) (uint64, error) {
	// 初始化查询语句
	metricsSql := "SELECT count(file_name) FROM config_file_release WHERE flag = 0 " +
		" AND namespace = $1 AND \"group\" = $2" // PostgreSQL 中使用双引号包裹列名

	if onlyActive {
		metricsSql += " AND active = 1" // 仅在需要时添加 active 条件
	}

	// 执行查询
	row := cfr.master.QueryRow(metricsSql, namespace, group)
	var total uint64
	if err := row.Scan(&total); err != nil {
		return 0, store.Error(err)
	}
	return total, nil
}

func (cfr *configFileReleaseStore) baseQuerySql() string {
	return "SELECT id, name, namespace, \"group\", file_name, content, COALESCE(comment, '') AS comment, " +
		" md5, version, EXTRACT(EPOCH FROM create_time) AS create_time, COALESCE(create_by, '') AS create_by, " +
		" EXTRACT(EPOCH FROM modify_time) AS modify_time, COALESCE(modify_by, '') AS modify_by, " +
		" flag, COALESCE(tags, '') AS tags, active, COALESCE(description, '') AS description, " +
		" COALESCE(release_type, '') AS release_type FROM config_file_release "
}

func (cfr *configFileReleaseStore) transferRows(rows *sql.Rows) ([]*model.ConfigFileRelease, error) {
	if rows == nil {
		return nil, nil
	}
	defer func() {
		_ = rows.Close()
	}()

	var fileReleases []*model.ConfigFileRelease

	for rows.Next() {
		fileRelease := model.NewConfigFileRelease()
		var (
			ctimeStr, mtimeStr string
			active             int64
			tags               string
		)
		err := rows.Scan(&fileRelease.Id, &fileRelease.Name, &fileRelease.Namespace, &fileRelease.Group,
			&fileRelease.FileName, &fileRelease.Content,
			&fileRelease.Comment, &fileRelease.Md5, &fileRelease.Version, &ctimeStr, &fileRelease.CreateBy,
			&mtimeStr, &fileRelease.ModifyBy, &fileRelease.Flag, &tags, &active, &fileRelease.ReleaseDescription,
			&fileRelease.ReleaseType)
		if err != nil {
			return nil, err
		}
		// 将字符串转换为int64
		ctimeFloat, err := strconv.ParseFloat(ctimeStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse create_time: %v", err)
		}
		mtimeFloat, err := strconv.ParseFloat(mtimeStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse modify_time: %v", err)
		}
		fileRelease.CreateTime = time.Unix(int64(ctimeFloat), 0)
		fileRelease.ModifyTime = time.Unix(int64(mtimeFloat), 0)
		fileRelease.Active = active == 1           // 转换 active 字段
		fileRelease.Valid = fileRelease.Flag == 0  // 检查有效性
		fileRelease.Metadata = map[string]string{} // 初始化 Metadata
		if err := json.Unmarshal([]byte(tags), &fileRelease.Metadata); err != nil {
			return nil, err // 如果解析失败，返回错误
		}
		fileReleases = append(fileReleases, fileRelease) // 添加到结果集合中
	}

	if err := rows.Err(); err != nil {
		return nil, err // 检查遍历行时的错误
	}

	return fileReleases, nil // 返回结果
}
