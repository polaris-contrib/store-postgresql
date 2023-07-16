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
	"github.com/polarismesh/polaris/store"
)

type configFileReleaseHistoryStore struct {
	db *BaseDB
}

// CreateConfigFileReleaseHistory 创建配置文件发布历史记录
func (rh *configFileReleaseHistoryStore) CreateConfigFileReleaseHistory(tx store.Tx,
	fileReleaseHistory *model.ConfigFileReleaseHistory) error {
	s := "insert into config_file_release_history(name, namespace, \"group\", file_name, content, comment, " +
		" md5, type, status, format, tags, " +
		"create_time, create_by, modify_time, modify_by) values " +
		"($1,$2,$3,$4,$5,$6,$7,$8, $9,$10,$11,$12,$13,$14,$15)"
	var err error
	if tx != nil {
		stmt, err := tx.GetDelegateTx().(*BaseTx).Prepare(s)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(fileReleaseHistory.Name, fileReleaseHistory.Namespace,
			fileReleaseHistory.Group, fileReleaseHistory.FileName, fileReleaseHistory.Content,
			fileReleaseHistory.Comment, fileReleaseHistory.Md5, fileReleaseHistory.Type,
			fileReleaseHistory.Status, fileReleaseHistory.Format, fileReleaseHistory.Tags,
			GetCurrentTimeFormat(), fileReleaseHistory.CreateBy, GetCurrentTimeFormat(),
			fileReleaseHistory.ModifyBy)
	} else {
		stmt, err := rh.db.Prepare(s)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(fileReleaseHistory.Name, fileReleaseHistory.Namespace,
			fileReleaseHistory.Group, fileReleaseHistory.FileName, fileReleaseHistory.Content,
			fileReleaseHistory.Comment, fileReleaseHistory.Md5, fileReleaseHistory.Type,
			fileReleaseHistory.Status, fileReleaseHistory.Format, fileReleaseHistory.Tags,
			GetCurrentTimeFormat(), fileReleaseHistory.CreateBy, GetCurrentTimeFormat(),
			fileReleaseHistory.ModifyBy)
	}
	if err != nil {
		return store.Error(err)
	}
	return nil
}

// QueryConfigFileReleaseHistories 获取配置文件的发布历史记录
func (rh *configFileReleaseHistoryStore) QueryConfigFileReleaseHistories(namespace, group, fileName string,
	offset, limit uint32, endId uint64) (uint32, []*model.ConfigFileReleaseHistory, error) {
	countSql := "select count(*) from config_file_release_history where "
	querySql := rh.genSelectSql() + " where "
	var idx = 1

	var queryParams []interface{}
	if namespace != "" {
		countSql += fmt.Sprintf(" namespace = $%d and ", idx)
		querySql += fmt.Sprintf(" namespace = $%d and ", idx)
		idx++
		queryParams = append(queryParams, namespace)
	}
	if endId > 0 {
		countSql += fmt.Sprintf(" id < $%d and ", idx)
		querySql += fmt.Sprintf(" id < $%d and ", idx)
		idx++
		queryParams = append(queryParams, endId)
	}

	countSql += fmt.Sprintf("group like $%d and file_name like $%d", idx, idx+1)
	querySql += fmt.Sprintf("group like $%d and file_name like $%d order by id desc limit $%d offset $%d",
		idx+2, idx+3, idx+4, idx+5)
	queryParams = append(queryParams, "%"+group+"%")
	queryParams = append(queryParams, "%"+fileName+"%")

	var count uint32
	err := rh.db.QueryRow(countSql, queryParams...).Scan(&count)
	if err != nil {
		return 0, nil, err
	}

	queryParams = append(queryParams, limit)
	queryParams = append(queryParams, offset)
	rows, err := rh.db.Query(querySql, queryParams...)
	if err != nil {
		return 0, nil, err
	}

	fileReleaseHistories, err := rh.transferRows(rows)
	if err != nil {
		return 0, nil, err
	}

	return count, fileReleaseHistories, nil
}

func (rh *configFileReleaseHistoryStore) GetLatestConfigFileReleaseHistory(namespace, group,
	fileName string) (*model.ConfigFileReleaseHistory, error) {
	s := rh.genSelectSql() + "where namespace = $1 and \"group\" = $2 and file_name = $3 order by id desc limit 1 offset 0"
	rows, err := rh.db.Query(s, namespace, group, fileName)
	if err != nil {
		return nil, err
	}

	fileReleaseHistories, err := rh.transferRows(rows)
	if err != nil {
		return nil, err
	}

	if len(fileReleaseHistories) == 0 {
		return nil, nil
	}

	return fileReleaseHistories[0], nil
}

func (rh *configFileReleaseHistoryStore) genSelectSql() string {
	return "select id, name, namespace, group, file_name, content, comment, md5, format, tags, type, " +
		" status, create_time, create_by, modify_time, " +
		"modify_by from config_file_release_history "
}

func (rh *configFileReleaseHistoryStore) transferRows(rows *sql.Rows) ([]*model.ConfigFileReleaseHistory, error) {
	if rows == nil {
		return nil, nil
	}
	defer rows.Close()

	var fileReleaseHistories []*model.ConfigFileReleaseHistory

	for rows.Next() {
		fileReleaseHistory := &model.ConfigFileReleaseHistory{}
		err := rows.Scan(&fileReleaseHistory.Id, &fileReleaseHistory.Name, &fileReleaseHistory.Namespace,
			&fileReleaseHistory.Group, &fileReleaseHistory.FileName, &fileReleaseHistory.Content,
			&fileReleaseHistory.Comment, &fileReleaseHistory.Md5, &fileReleaseHistory.Format,
			&fileReleaseHistory.Tags, &fileReleaseHistory.Type, &fileReleaseHistory.Status,
			&fileReleaseHistory.CreateTime, &fileReleaseHistory.CreateBy, &fileReleaseHistory.ModifyTime,
			&fileReleaseHistory.ModifyBy)
		if err != nil {
			return nil, err
		}
		fileReleaseHistories = append(fileReleaseHistories, fileReleaseHistory)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return fileReleaseHistories, nil
}
