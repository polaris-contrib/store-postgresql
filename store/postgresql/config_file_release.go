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
	"time"

	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
)

var _ store.ConfigFileReleaseStore = (*configFileReleaseStore)(nil)

type configFileReleaseStore struct {
	db    *BaseDB
	slave *BaseDB
}

// CreateConfigFileRelease 新建配置文件发布
func (cfr *configFileReleaseStore) CreateConfigFileRelease(tx store.Tx,
	fileRelease *model.ConfigFileRelease) (*model.ConfigFileRelease, error) {
	s := "insert into config_file_release(name, namespace, \"group\", file_name, content, comment, md5, version, " +
		" create_time, create_by, modify_time, modify_by) values" +
		"($1,$2,$3,$4,$5,$6,$7,$8, $9,$10,$11,$12)"
	var err error
	if tx != nil {
		stmt, err := tx.GetDelegateTx().(*BaseTx).Prepare(s)
		if err != nil {
			return nil, store.Error(err)
		}
		_, err = stmt.Exec(fileRelease.Name, fileRelease.Namespace, fileRelease.Group,
			fileRelease.FileName, fileRelease.Content, fileRelease.Comment, fileRelease.Md5,
			fileRelease.Version, GetCurrentTimeFormat(), fileRelease.CreateBy,
			GetCurrentTimeFormat(), fileRelease.ModifyBy)
	} else {
		stmt, err := cfr.db.Prepare(s)
		if err != nil {
			return nil, store.Error(err)
		}
		_, err = stmt.Exec(fileRelease.Name, fileRelease.Namespace, fileRelease.Group,
			fileRelease.FileName, fileRelease.Content, fileRelease.Comment, fileRelease.Md5,
			fileRelease.Version, GetCurrentTimeFormat(), fileRelease.CreateBy,
			GetCurrentTimeFormat(), fileRelease.ModifyBy)
	}
	if err != nil {
		return nil, store.Error(err)
	}
	return cfr.GetConfigFileRelease(tx, fileRelease.Namespace, fileRelease.Group, fileRelease.FileName)
}

// UpdateConfigFileRelease 更新配置文件发布
func (cfr *configFileReleaseStore) UpdateConfigFileRelease(tx store.Tx,
	fileRelease *model.ConfigFileRelease) (*model.ConfigFileRelease, error) {
	s := "update config_file_release set name = $1 , content = $2, comment = $3, md5 = $4, version = $5, flag = 0, " +
		" modify_time = $6, modify_by = $7 where namespace = $8 and \"group\" = $9 and file_name = $10"
	var err error
	if tx != nil {
		stmt, err := tx.GetDelegateTx().(*BaseTx).Prepare(s)
		if err != nil {
			return nil, store.Error(err)
		}
		_, err = stmt.Exec(fileRelease.Name, fileRelease.Content, fileRelease.Comment,
			fileRelease.Md5, fileRelease.Version, GetCurrentTimeFormat(), fileRelease.ModifyBy,
			fileRelease.Namespace, fileRelease.Group, fileRelease.FileName)
	} else {
		stmt, err := cfr.db.Prepare(s)
		if err != nil {
			return nil, store.Error(err)
		}
		_, err = stmt.Exec(fileRelease.Name, fileRelease.Content, fileRelease.Comment,
			fileRelease.Md5, fileRelease.Version, GetCurrentTimeFormat(), fileRelease.ModifyBy,
			fileRelease.Namespace, fileRelease.Group, fileRelease.FileName)
	}
	if err != nil {
		return nil, store.Error(err)
	}
	return cfr.GetConfigFileRelease(tx, fileRelease.Namespace, fileRelease.Group, fileRelease.FileName)
}

// GetConfigFileRelease 获取配置文件发布，只返回 flag=0 的记录
func (cfr *configFileReleaseStore) GetConfigFileRelease(tx store.Tx, namespace,
	group, fileName string) (*model.ConfigFileRelease, error) {
	return cfr.getConfigFileReleaseByFlag(tx, namespace, group, fileName, false)
}

func (cfr *configFileReleaseStore) GetConfigFileReleaseWithAllFlag(tx store.Tx, namespace,
	group, fileName string) (*model.ConfigFileRelease, error) {
	return cfr.getConfigFileReleaseByFlag(tx, namespace, group, fileName, true)
}

func (cfr *configFileReleaseStore) getConfigFileReleaseByFlag(tx store.Tx, namespace, group,
	fileName string, withAllFlag bool) (*model.ConfigFileRelease, error) {
	querySql := cfr.baseQuerySql() + "where namespace = $1 and \"group\" = $2 and file_name = $3 and flag = 0"

	if withAllFlag {
		querySql = cfr.baseQuerySql() + "where namespace = $1 and \"group\" = $2 and file_name = $3"
	}

	var (
		rows *sql.Rows
		err  error
	)

	if tx != nil {
		rows, err = tx.GetDelegateTx().(*BaseTx).Query(querySql, namespace, group, fileName)
	} else {
		rows, err = cfr.db.Query(querySql, namespace, group, fileName)
	}
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

func (cfr *configFileReleaseStore) DeleteConfigFileRelease(tx store.Tx, namespace, group,
	fileName, deleteBy string) error {
	s := "update config_file_release set flag = 1, modify_time = $1, modify_by = $2, version = version + 1, " +
		" md5='' where namespace = $3 and \"group\" = $4 and file_name = $5"
	var err error
	if tx != nil {
		stmt, err := tx.GetDelegateTx().(*BaseTx).Prepare(s)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(GetCurrentTimeFormat(), deleteBy, namespace, group, fileName)
	} else {
		stmt, err := cfr.db.Prepare(s)
		if err != nil {
			return store.Error(err)
		}
		_, err = stmt.Exec(GetCurrentTimeFormat(), deleteBy, namespace, group, fileName)
	}
	if err != nil {
		return store.Error(err)
	}
	return nil
}

// FindConfigFileReleaseByModifyTimeAfter 获取最后更新时间大于某个时间点的发布，注意包含 flag = 1 的，为了能够获取被删除的 release
func (cfr *configFileReleaseStore) FindConfigFileReleaseByModifyTimeAfter(
	modifyTime time.Time) ([]*model.ConfigFileRelease, error) {
	s := cfr.baseQuerySql() + " where modify_time > $1"
	rows, err := cfr.slave.Query(s, modifyTime)
	if err != nil {
		return nil, err
	}
	releases, err := cfr.transferRows(rows)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

func (cfr *configFileReleaseStore) CountConfigFileReleaseEachGroup() (map[string]map[string]int64, error) {
	metricsSql := "SELECT namespace, \"group\", count(file_name) FROM config_file_release " +
		" WHERE flag = 0 GROUP by namespace, \"group\""
	rows, err := cfr.slave.Query(metricsSql)
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

func (cfr *configFileReleaseStore) baseQuerySql() string {
	return "select id, name, namespace, \"group\", file_name, content, comment, md5, version, " +
		" create_time, create_by, modify_time, modify_by, " +
		" flag from config_file_release "
}

func (cfr *configFileReleaseStore) transferRows(rows *sql.Rows) ([]*model.ConfigFileRelease, error) {
	if rows == nil {
		return nil, nil
	}
	defer rows.Close()

	var fileReleases []*model.ConfigFileRelease

	for rows.Next() {
		fileRelease := &model.ConfigFileRelease{}
		err := rows.Scan(&fileRelease.Id, &fileRelease.Name, &fileRelease.Namespace,
			&fileRelease.Group, &fileRelease.FileName, &fileRelease.Content, &fileRelease.Comment,
			&fileRelease.Md5, &fileRelease.Version, &fileRelease.CreateTime, &fileRelease.CreateBy,
			&fileRelease.ModifyTime, &fileRelease.ModifyBy, &fileRelease.Flag)
		if err != nil {
			return nil, err
		}

		fileReleases = append(fileReleases, fileRelease)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return fileReleases, nil
}
