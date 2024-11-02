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
	"github.com/polarismesh/polaris/common/utils"
	"strconv"
	"time"

	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
)

type configFileReleaseHistoryStore struct {
	master *BaseDB
	slave  *BaseDB
}

// CreateConfigFileReleaseHistory 创建配置文件发布历史记录
func (rh *configFileReleaseHistoryStore) CreateConfigFileReleaseHistory(history *model.ConfigFileReleaseHistory) error {
	s := "insert into config_file_release_history(name, namespace, \"group\", file_name, content, comment, " +
		" md5, type, status, format, tags, " +
		"create_time, create_by, modify_time, modify_by, version, reason, description) values " +
		"($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)"
	stmt, err := rh.master.Prepare(s)
	if err != nil {
		return store.Error(err)
	}

	if _, err := stmt.Exec(history.Name, history.Namespace,
		history.Group, history.FileName, history.Content, history.Comment,
		history.Md5, history.Type, history.Status, history.Format,
		utils.MustJson(history.Metadata), history.CreateBy, history.ModifyBy,
		history.Version, history.Reason, history.ReleaseDescription); err != nil {
		return store.Error(err)
	}

	return nil
}

// QueryConfigFileReleaseHistories 获取配置文件的发布历史记录
func (rh *configFileReleaseHistoryStore) QueryConfigFileReleaseHistories(filter map[string]string,
	offset, limit uint32) (uint32, []*model.ConfigFileReleaseHistory, error) {
	countSql := "select count(*) from config_file_release_history where "
	querySql := rh.genSelectSql() + " where "

	var idx = 1
	namespace := filter["namespace"]
	group := filter["group"]
	fileName := filter["name"]
	endId, _ := strconv.ParseUint(filter["endId"], 10, 64)

	var queryParams []interface{}
	if namespace != "" {
		countSql += fmt.Sprintf(" namespace = $%d AND ", idx)
		querySql += fmt.Sprintf(" namespace = $%d AND ", idx)
		idx++
		queryParams = append(queryParams, namespace)
	}
	if endId > 0 {
		countSql += fmt.Sprintf(" id < $%d AND ", idx)
		querySql += fmt.Sprintf(" id < $%d AND ", idx)
		idx++
		queryParams = append(queryParams, endId)
	}

	countSql += fmt.Sprintf("\"group\" like $%d and file_name like $%d", idx, idx+1)
	querySql += fmt.Sprintf("\"group\" like $%d and file_name like $%d order by id desc limit $%d offset $%d",
		idx, idx+1, idx+2, idx+3)
	queryParams = append(queryParams, "%"+group+"%")
	queryParams = append(queryParams, "%"+fileName+"%")

	var count uint32
	err := rh.master.QueryRow(countSql, queryParams...).Scan(&count)
	if err != nil {
		return 0, nil, err
	}

	queryParams = append(queryParams, limit)
	queryParams = append(queryParams, offset)
	rows, err := rh.master.Query(querySql, queryParams...)
	if err != nil {
		return 0, nil, err
	}

	fileReleaseHistories, err := rh.transferRows(rows)
	if err != nil {
		return 0, nil, err
	}

	return count, fileReleaseHistories, nil
}

// CleanConfigFileReleaseHistory 清理配置发布历史
func (rh *configFileReleaseHistoryStore) CleanConfigFileReleaseHistory(endTime time.Time, limit uint64) error {
	delSql := "DELETE FROM config_file_release_history WHERE create_time < $1 LIMIT $2"
	_, err := rh.master.Exec(delSql, endTime, limit)
	return err
}

func (rh *configFileReleaseHistoryStore) genSelectSql() string {
	return "select id, name, namespace, \"group\", file_name, content, comment, md5, format, tags, type, " +
		" status, create_time, create_by, modify_time, " +
		"modify_by, reason, description, version from config_file_release_history "
}

func (rh *configFileReleaseHistoryStore) transferRows(rows *sql.Rows) ([]*model.ConfigFileReleaseHistory, error) {
	if rows == nil {
		return nil, nil
	}
	defer rows.Close()

	records := make([]*model.ConfigFileReleaseHistory, 0, 16)

	for rows.Next() {
		item := &model.ConfigFileReleaseHistory{}
		var (
			ctimeStr, mtimeStr string
			tags               string
		)
		err := rows.Scan(&item.Id, &item.Name, &item.Namespace, &item.Group, &item.FileName, &item.Content,
			&item.Comment, &item.Md5, &item.Format, &tags, &item.Type, &item.Status, &ctimeStr, &item.CreateBy,
			&mtimeStr, &item.ModifyBy, &item.Reason, &item.ReleaseDescription, &item.Version)
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

		item.CreateTime = time.Unix(int64(ctimeFloat), 0)
		item.ModifyTime = time.Unix(int64(mtimeFloat), 0)
		item.Metadata = map[string]string{}
		_ = json.Unmarshal([]byte(tags), &item.Metadata)

		records = append(records, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
