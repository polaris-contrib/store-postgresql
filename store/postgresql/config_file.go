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
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/common/utils"
	"github.com/polarismesh/polaris/store"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

var (
	configFileStoreFieldMapping = map[string]map[string]string{
		"config_file": {
			"group":     "`group`",
			"file_name": "name",
			"namespace": "namespace",
		},
		"config_file_release": {
			"group":        "`group`",
			"file_name":    "file_name",
			"release_name": "name",
		},
		"config_file_group": {},
	}
)

var _ store.ConfigFileStore = (*configFileStore)(nil)

type configFileStore struct {
	master *BaseDB
	slave  *BaseDB
}

// LockConfigFile 加锁配置文件
func (cf *configFileStore) LockConfigFile(tx store.Tx, file *model.ConfigFileKey) (*model.ConfigFile, error) {
	if tx == nil {
		return nil, ErrTxIsNil
	}

	dbTx := tx.GetDelegateTx().(*BaseTx)
	args := []interface{}{file.Namespace, file.Group, file.Name}
	lockSql := cf.baseSelectConfigFileSql() +
		` WHERE namespace = $1 AND "group" = $2 AND name = $3 AND flag = 0 FOR UPDATE`

	rows, err := dbTx.Query(lockSql, args...)
	if err != nil {
		return nil, store.Error(err)
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

// CreateConfigFileTx 创建配置文件
func (cf *configFileStore) CreateConfigFileTx(tx store.Tx, file *model.ConfigFile) error {
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

	deleteSql := `DELETE FROM config_file WHERE namespace = $1 AND "group" = $2 AND name = $3 AND flag = 1`
	if _, err = dbTx.Exec(deleteSql, file.Namespace, file.Group, file.Name); err != nil {
		return store.Error(err)
	}

	createSql := `INSERT INTO config_file (
		name, namespace, "group", content, comment, format, create_time, create_by, modify_time, modify_by
	) VALUES (
		$1, $2, $3, $4, $5, $6, current_timestamp, $7, current_timestamp, $8
	)`
	if _, err = dbTx.Exec(createSql, file.Name, file.Namespace, file.Group, file.Content, file.Comment, file.Format,
		file.CreateBy, file.ModifyBy); err != nil {
		return store.Error(err)
	}

	if err = cf.batchCleanTags(dbTx, file); err != nil {
		return store.Error(err)
	}

	if err = cf.batchAddTags(dbTx, file); err != nil {
		return store.Error(err)
	}

	if err = dbTx.Commit(); err != nil {
		return store.Error(err)
	}

	return nil
}

func (cf *configFileStore) batchAddTags(tx *BaseTx, file *model.ConfigFile) error {
	if len(file.Metadata) == 0 {
		return nil
	}

	// 添加配置标签
	insertSql := `INSERT INTO config_file_tag (
		"key", "value", namespace, "group", file_name, create_time, create_by, modify_time, modify_by
	) VALUES `
	valuesSql := []string{}
	var args []interface{}
	for k, v := range file.Metadata {
		valuesSql = append(valuesSql, "($1, $2, $3, $4, $5, current_timestamp, $6, current_timestamp, $7)")
		args = append(args, k, v, file.Namespace, file.Group, file.Name, file.CreateBy, file.ModifyBy)
	}
	insertSql = insertSql + strings.Join(valuesSql, ",")
	_, err := tx.Exec(insertSql, args...)
	return store.Error(err)
}

func (cf *configFileStore) batchCleanTags(tx *BaseTx, file *model.ConfigFile) error {
	// 删除配置标签
	cleanSql := `DELETE FROM config_file_tag WHERE namespace = $1 AND "group" = $2 AND file_name = $3`
	args := []interface{}{file.Namespace, file.Group, file.Name}
	_, err := tx.Exec(cleanSql, args...)
	return store.Error(err)
}

func (cf *configFileStore) loadFileTags(tx *BaseTx, file *model.ConfigFile) error {
	querySql := `SELECT "key", "value" FROM config_file_tag WHERE namespace = $1 AND 
		"group" = $2 AND file_name = $3`

	rows, err := tx.Query(querySql, file.Namespace, file.Group, file.Name)
	if err != nil {
		return err
	}
	if rows == nil {
		return nil
	}
	defer rows.Close()

	file.Metadata = make(map[string]string)
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return err
		}
		file.Metadata[key] = value
	}
	return nil
}

func (cf *configFileStore) CountConfigFiles(namespace, group string) (uint64, error) {
	metricsSql := `SELECT count(*) FROM config_file WHERE flag = 0 AND namespace = $1 AND "group" = $2`
	row := cf.slave.QueryRow(metricsSql, namespace, group)
	var total uint64
	if err := row.Scan(&total); err != nil {
		return 0, store.Error(err)
	}
	return total, nil
}

// GetConfigFile 获取配置文件
func (cf *configFileStore) GetConfigFile(namespace, group, name string) (*model.ConfigFile, error) {
	tx, err := cf.master.Begin()
	if err != nil {
		return nil, store.Error(err)
	}
	defer func() {
		_ = tx.Rollback()
	}()

	return cf.GetConfigFileTx(NewSqlDBTx(tx), namespace, group, name)
}

func (cf *configFileStore) GetConfigFileTx(tx store.Tx,
	namespace, group, name string) (*model.ConfigFile, error) {
	if tx == nil {
		return nil, ErrTxIsNil
	}

	dbTx := tx.GetDelegateTx().(*BaseTx)
	querySql := cf.baseSelectConfigFileSql() + `WHERE namespace = $1 AND "group" = $2 AND name = $3 AND flag = 0`
	rows, err := dbTx.Query(querySql, namespace, group, name)
	if err != nil {
		return nil, store.Error(err)
	}
	files, err := cf.transferRows(rows)
	if err != nil {
		return nil, store.Error(err)
	}
	if len(files) == 0 {
		return nil, nil
	}
	if err := cf.loadFileTags(dbTx, files[0]); err != nil {
		return nil, store.Error(err)
	}
	return files[0], nil
}

// UpdateConfigFileTx 更新配置文件
func (cf *configFileStore) UpdateConfigFileTx(tx store.Tx, file *model.ConfigFile) error {
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

	updateSql := `UPDATE config_file SET content = $1, comment = $2, format = $3, 
		modify_time = current_timestamp, modify_by = $4 
		WHERE namespace = $5 AND "group" = $6 AND name = $7`
	_, err = dbTx.Exec(updateSql, file.Content, file.Comment, file.Format,
		file.ModifyBy, file.Namespace, file.Group, file.Name)
	if err != nil {
		return store.Error(err)
	}

	if err = cf.batchCleanTags(dbTx, file); err != nil {
		return store.Error(err)
	}

	if err = cf.batchAddTags(dbTx, file); err != nil {
		return store.Error(err)
	}

	if err = dbTx.Commit(); err != nil {
		return store.Error(err)
	}

	return nil
}

func (cf *configFileStore) DeleteConfigFileTx(tx store.Tx, namespace, group, name string) error {
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

	deleteSql := `UPDATE config_file SET flag = 1 WHERE namespace = $1 AND "group" = $2 AND name = $3`
	if _, err = dbTx.Exec(deleteSql, namespace, group, name); err != nil {
		return store.Error(err)
	}

	if err = dbTx.Commit(); err != nil {
		return store.Error(err)
	}

	return nil
}

// QueryConfigFiles 翻页查询配置文件，group、name可为模糊匹配
func (cf *configFileStore) QueryConfigFiles(filter map[string]string, offset, limit uint32) (uint32, []*model.ConfigFile, error) {

	countSql := `SELECT COUNT(*) FROM config_file WHERE flag = 0 `
	querySql := cf.baseSelectConfigFileSql() + ` WHERE flag = 0 `

	args := make([]interface{}, 0, len(filter))
	searchQuery := make([]string, 0, len(filter))

	for k, v := range filter {
		if v, ok := configFileStoreFieldMapping["config_file"][k]; ok {
			k = v
		}
		if utils.IsWildName(v) {
			searchQuery = append(searchQuery, k+" LIKE $"+strconv.Itoa(len(args)+1))
		} else {
			searchQuery = append(searchQuery, k+" = $"+strconv.Itoa(len(args)+1))
		}
		args = append(args, utils.ParseWildNameForSql(v))
	}

	if len(searchQuery) > 0 {
		countSql += " AND "
		querySql += " AND "
	}
	countSql += strings.Join(searchQuery, " AND ")

	var count uint32
	err := cf.master.QueryRow(countSql, args...).Scan(&count)
	if err != nil {
		log.Error("[Config][Storage] query config files", zap.String("count-sql", countSql), zap.Error(err))
		return 0, nil, store.Error(err)
	}

	// 使用 OFFSET 和 LIMIT
	querySql += strings.Join(searchQuery, " AND ") + " ORDER BY id DESC LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)

	args = append(args, limit, offset)
	rows, err := cf.master.Query(querySql, args...)
	if err != nil {
		log.Error("[Config][Storage] query config files", zap.String("query-sql", countSql), zap.Error(err))
		return 0, nil, store.Error(err)
	}

	files, err := cf.transferRows(rows)
	if err != nil {
		return 0, nil, store.Error(err)
	}

	err = cf.slave.processWithTransaction("batch-load-file-tags", func(tx *BaseTx) error {
		for i := range files {
			item := files[i]
			if err := cf.loadFileTags(tx, item); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return 0, nil, store.Error(err)
	}

	return count, files, nil
}

func (cf *configFileStore) CountConfigFileEachGroup() (map[string]map[string]int64, error) {
	metricsSql := `SELECT namespace, "group", COUNT(name) 
		FROM config_file 
		WHERE flag = 0 
		GROUP BY namespace, "group"`

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
			namespace string
			group     string
			cnt       int64
		)

		if err := rows.Scan(&namespace, &group, &cnt); err != nil {
			return nil, err
		}
		if _, ok := ret[namespace]; !ok {
			ret[namespace] = map[string]int64{}
		}
		ret[namespace][group] = cnt
	}

	return ret, nil
}

func (cf *configFileStore) baseSelectConfigFileSql() string {
	return `SELECT id, name, namespace, "group", content, 
		COALESCE(comment, '') AS comment, 
		format, 
		EXTRACT(EPOCH FROM create_time) AS create_time, 
		COALESCE(create_by, '') AS create_by, 
		EXTRACT(EPOCH FROM modify_time) AS modify_time, 
		COALESCE(modify_by, '') AS modify_by 
		FROM config_file `
}

func (cf *configFileStore) hardDeleteConfigFile(namespace, group, name string) error {
	deleteSql := `DELETE FROM config_file WHERE namespace = $1 AND "group" = $2 AND name = $3 AND flag = 1`
	_, err := cf.master.Exec(deleteSql, namespace, group, name)
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

	var (
		files = make([]*model.ConfigFile, 0, 32)
	)

	for rows.Next() {
		file := &model.ConfigFile{
			Metadata: map[string]string{},
		}

		var ctimeStr, mtimeStr string
		if err := rows.Scan(&file.Id, &file.Name, &file.Namespace, &file.Group, &file.Content, &file.Comment,
			&file.Format, &ctimeStr, &file.CreateBy, &mtimeStr, &file.ModifyBy); err != nil {
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

		file.CreateTime = time.Unix(int64(ctimeFloat), 0)
		file.ModifyTime = time.Unix(int64(mtimeFloat), 0)

		files = append(files, file)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return files, nil
}
