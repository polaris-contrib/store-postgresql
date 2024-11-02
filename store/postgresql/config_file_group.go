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
	"fmt"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/common/utils"
	"github.com/polarismesh/polaris/store"
	"strconv"
	"time"
)

type configFileGroupStore struct {
	master *BaseDB
	slave  *BaseDB
}

// CreateConfigFileGroup 创建配置文件组
func (fg *configFileGroupStore) CreateConfigFileGroup(
	fileGroup *model.ConfigFileGroup) (*model.ConfigFileGroup, error) {
	err := fg.master.processWithTransaction("", func(tx *BaseTx) error {
		// 删除标志为 1 的记录
		if _, err := tx.Exec("DELETE FROM config_file_group WHERE flag = 1 AND namespace = $1 AND name = $2",
			fileGroup.Namespace, fileGroup.Name); err != nil {
			return store.Error(err)
		}

		// 插入新的配置文件组
		createSql := `
			INSERT INTO config_file_group (name, namespace, comment, create_time, create_by,
				modify_time, modify_by, owner, business, department, metadata)
			VALUES ($1, $2, $3, current_timestamp, $4, current_timestamp, $5, $6, $7, $8, $9)
		`
		args := []interface{}{
			fileGroup.Name, fileGroup.Namespace, fileGroup.Comment,
			fileGroup.CreateBy, fileGroup.ModifyBy, fileGroup.Owner, fileGroup.Business,
			fileGroup.Department, utils.MustJson(fileGroup.Metadata),
		}
		if _, err := tx.Exec(createSql, args...); err != nil {
			return store.Error(err)
		}
		return tx.Commit()
	})
	if err != nil {
		return nil, store.Error(err)
	}

	return fg.GetConfigFileGroup(fileGroup.Namespace, fileGroup.Name)
}

// UpdateConfigFileGroup 更新配置文件组信息
func (fg *configFileGroupStore) UpdateConfigFileGroup(fileGroup *model.ConfigFileGroup) error {
	updateSql := `
		UPDATE config_file_group 
		SET comment = $1, modify_time = CURRENT_TIMESTAMP, modify_by = $2, 
			business = $3, department = $4, metadata = $5 
		WHERE namespace = $6 AND name = $7
	`

	args := []interface{}{
		fileGroup.Comment,
		fileGroup.ModifyBy,
		fileGroup.Business,
		fileGroup.Department,
		utils.MustJson(fileGroup.Metadata),
		fileGroup.Namespace,
		fileGroup.Name,
	}

	if _, err := fg.master.Exec(updateSql, args...); err != nil {
		return store.Error(err)
	}
	return nil
}

// GetConfigFileGroup 获取配置文件组
func (fg *configFileGroupStore) GetConfigFileGroup(namespace, name string) (*model.ConfigFileGroup, error) {
	querySql := fg.genConfigFileGroupSelectSql() + " WHERE namespace=$1 AND name=$2 AND flag = 0"
	rows, err := fg.master.Query(querySql, namespace, name)
	if err != nil {
		return nil, store.Error(err)
	}
	cfgs, err := fg.transferRows(rows)
	if err != nil {
		return nil, err
	}
	if len(cfgs) > 0 {
		return cfgs[0], nil
	}
	return nil, nil
}

// DeleteConfigFileGroup 删除配置文件组
func (fg *configFileGroupStore) DeleteConfigFileGroup(namespace, name string) error {
	deleteSql := "UPDATE config_file_group SET flag = 1 WHERE namespace = $1 and name = $2"

	log.Infof("[Config][Storage] delete config file group(%s, %s)", namespace, name)
	if _, err := fg.master.Exec(deleteSql, namespace, name); err != nil {
		return err
	}

	return nil
}

func (fg *configFileGroupStore) GetMoreConfigGroup(firstUpdate bool, mtime time.Time) ([]*model.ConfigFileGroup, error) {
	if firstUpdate {
		mtime = time.Unix(0, 1)
	}
	loadSql := `
SELECT 
	id, 
	name, 
	namespace, 
	COALESCE(comment, '') AS comment, 
	EXTRACT(EPOCH FROM create_time) AS create_time, 
	COALESCE(create_by, '') AS create_by, 
	EXTRACT(EPOCH FROM modify_time) AS modify_time, 
	COALESCE(modify_by, '') AS modify_by, 
	COALESCE(owner, '') AS owner, 
	COALESCE(business, '') AS business, 
	COALESCE(department, '') AS department, 
	COALESCE(metadata, '{}') AS metadata, 
	flag 
FROM config_file_group 
WHERE modify_time >= $1
`

	rows, err := fg.slave.Query(loadSql, mtime)
	if err != nil {
		return nil, err
	}
	return fg.transferRows(rows)
}

func (fg *configFileGroupStore) CountConfigGroups(namespace string) (uint64, error) {
	metricsSql := "SELECT count(*) FROM config_file_group WHERE flag = 0 AND namespace = $1"
	row := fg.master.QueryRow(metricsSql, namespace)
	var total uint64
	if err := row.Scan(&total); err != nil {
		return 0, store.Error(err)
	}
	return total, nil
}

func (fg *configFileGroupStore) genConfigFileGroupSelectSql() string {
	return `
SELECT 
	id, 
	name, 
	namespace, 
	COALESCE(comment, '') AS comment, 
	EXTRACT(EPOCH FROM create_time) AS create_time, 
	COALESCE(create_by, '') AS create_by, 
	EXTRACT(EPOCH FROM modify_time) AS modify_time, 
	COALESCE(modify_by, '') AS modify_by, 
	COALESCE(owner, '') AS owner, 
	COALESCE(business, '') AS business, 
	COALESCE(department, '') AS department, 
	COALESCE(metadata, '{}') AS metadata, 
	flag 
FROM config_file_group
`
}

func (fg *configFileGroupStore) transferRows(rows *sql.Rows) ([]*model.ConfigFileGroup, error) {
	if rows == nil {
		return nil, nil
	}
	defer rows.Close()

	var fileGroups []*model.ConfigFileGroup

	for rows.Next() {
		fileGroup := &model.ConfigFileGroup{}
		var (
			//ctime, mtime int64
			ctimeStr, mtimeStr string
			flag               int64
			metadata           string
		)

		// 读取数据
		err := rows.Scan(
			&fileGroup.Id,
			&fileGroup.Name,
			&fileGroup.Namespace,
			&fileGroup.Comment,
			&ctimeStr,
			&fileGroup.CreateBy,
			&mtimeStr,
			&fileGroup.ModifyBy,
			&fileGroup.Owner,
			&fileGroup.Business,
			&fileGroup.Department,
			&metadata,
			&flag,
		)
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
		fileGroup.CreateTime = time.Unix(int64(ctimeFloat), 0)
		fileGroup.ModifyTime = time.Unix(int64(mtimeFloat), 0)

		// 处理 Metadata
		fileGroup.Metadata = make(map[string]string)
		if err := json.Unmarshal([]byte(metadata), &fileGroup.Metadata); err != nil {
			return nil, err
		}

		// 设置 Valid 状态
		fileGroup.Valid = flag == 0

		// 将文件组添加到结果集中
		fileGroups = append(fileGroups, fileGroup)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return fileGroups, nil
}
