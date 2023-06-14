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
	"errors"
	"fmt"
	"time"

	"github.com/polarismesh/polaris/common/log"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
)

// namespaceStore 实现了NamespaceStore
type namespaceStore struct {
	master *BaseDB
	slave  *BaseDB
}

// AddNamespace 添加命名空间
func (ns *namespaceStore) AddNamespace(namespace *model.Namespace) error {
	if namespace.Name == "" {
		return errors.New("store add namespace name is empty")
	}

	return RetryTransaction("addNamespace", func() error {
		return ns.master.processWithTransaction("addNamespace", func(tx *BaseTx) error {
			// 先删除无效数据，再添加新数据
			if err := cleanNamespace(tx, namespace.Name); err != nil {
				return err
			}

			str := fmt.Sprintf("insert into namespace(name, comment, token, owner, ctime, mtime) values($1, $2, $3, $4, $5, $6)")
			stmt, err := tx.Prepare(str)
			if err != nil {
				log.Errorf("[Store][database] insert prepare[%v] commit tx err: %s", namespace, err.Error())
				return store.Error(err)
			}

			if _, err := stmt.Exec(namespace.Name, namespace.Comment, namespace.Token, namespace.Owner, GetCurrentTimeFormat(), GetCurrentTimeFormat()); err != nil {
				return store.Error(err)
			}

			if err := tx.Commit(); err != nil {
				log.Errorf("[Store][database] batch insert instance commit tx err: %s", err.Error())
				return err
			}

			return nil
		})
	})
}

// UpdateNamespace 更新命名空间，目前只更新owner
func (ns *namespaceStore) UpdateNamespace(namespace *model.Namespace) error {
	if namespace.Name == "" {
		return errors.New("store update namespace name is empty")
	}

	return RetryTransaction("updateNamespace", func() error {
		return ns.master.processWithTransaction("updateNamespace", func(tx *BaseTx) error {
			stmt, err := tx.Prepare("update namespace set owner = $1, comment = $2, mtime = $3 where name = $4")
			if err != nil {
				return store.Error(err)
			}

			if _, err := stmt.Exec(namespace.Owner, namespace.Comment, GetCurrentTimeFormat(), namespace.Name); err != nil {
				return store.Error(err)
			}

			if err := tx.Commit(); err != nil {
				log.Errorf("[Store][database] batch delete instance commit tx err: %s", err.Error())
				return err
			}

			return nil
		})
	})
}

// UpdateNamespaceToken 更新命名空间token
func (ns *namespaceStore) UpdateNamespaceToken(name string, token string) error {
	if name == "" || token == "" {
		return fmt.Errorf("store update namespace token some param are empty, name is %s, token is %s", name, token)
	}
	return RetryTransaction("updateNamespaceToken", func() error {
		return ns.master.processWithTransaction("updateNamespaceToken", func(tx *BaseTx) error {
			/*str := fmt.Sprintf("update namespace set token = '%s', mtime = '%s' where name = '%s'",
			token, GetCurrentTimeFormat(), name)*/

			str := "update namespace set token = $1, mtime = $2 where name = $3"
			if _, err := tx.Exec(str, token, GetCurrentTimeFormat(), name); err != nil {
				return store.Error(err)
			}

			if err := tx.Commit(); err != nil {
				log.Errorf("[Store][database] batch delete instance commit tx err: %s", err.Error())
				return err
			}

			return nil
		})
	})
}

// GetNamespace 根据名字获取命名空间详情，只返回有效的
func (ns *namespaceStore) GetNamespace(name string) (*model.Namespace, error) {
	namespace, err := ns.getNamespace(name)
	if err != nil {
		return nil, err
	}

	if namespace != nil && !namespace.Valid {
		return nil, nil
	}

	return namespace, nil
}

// GetNamespaces 根据过滤条件查询命名空间及数目
func (ns *namespaceStore) GetNamespaces(filter map[string][]string, offset, limit int) ([]*model.Namespace, uint32, error) {
	// 只查询有效数据
	filter["flag"] = []string{"0"}

	num, err := ns.getNamespacesCount(filter)
	if err != nil {
		return nil, 0, err
	}

	out, err := ns.getNamespaces(filter, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return out, num, nil
}

// GetMoreNamespaces 根据mtime获取命名空间
func (ns *namespaceStore) GetMoreNamespaces(mtime time.Time) ([]*model.Namespace, error) {
	str := genNamespaceSelectSQL() + " where mtime >= $1"
	rows, err := ns.slave.Query(str, mtime)
	if err != nil {
		log.Errorf("[Store][database] get more namespace query err: %s", err.Error())
		return nil, err
	}

	return namespaceFetchRows(rows)
}

// getNamespacesCount 根据相关条件查询对应命名空间数目
func (ns *namespaceStore) getNamespacesCount(filter map[string][]string) (uint32, error) {
	str := `select count(*) from namespace `
	str, args := genNamespaceWhereSQLAndArgs(str, filter, nil, 0, 1)

	var count uint32

	err := ns.master.QueryRow(str, args...).Scan(&count)
	switch {
	case err == sql.ErrNoRows:
		log.Errorf("[Store][database] no row with this namespace filter")
		return count, err
	case err != nil:
		log.Errorf("[Store][database] get namespace count by filter err: %s", err.Error())
		return count, err
	default:
		return count, err
	}
}

// getNamespaces 根据相关条件查询对应命名空间
func (ns *namespaceStore) getNamespaces(filter map[string][]string, offset, limit int) ([]*model.Namespace, error) {
	str := genNamespaceSelectSQL()
	order := &Order{"mtime", "desc"}
	str, args := genNamespaceWhereSQLAndArgs(str, filter, order, offset, limit)

	rows, err := ns.master.Query(str, args...)
	if err != nil {
		log.Errorf("[Store][database] get namespaces by filter query err: %s", err.Error())
		return nil, err
	}

	return namespaceFetchRows(rows)
}

// getNamespace 获取namespace的内部函数，从数据库中拉取数据
func (ns *namespaceStore) getNamespace(name string) (*model.Namespace, error) {
	if name == "" {
		return nil, errors.New("store get namespace name is empty")
	}

	str := genNamespaceSelectSQL() + " where name = $1"
	rows, err := ns.master.Query(str, name)
	if err != nil {
		log.Errorf("[Store][database] get namespace query err: %s", err.Error())
		return nil, err
	}

	out, err := namespaceFetchRows(rows)
	if err != nil {
		return nil, err
	}

	if len(out) == 0 {
		return nil, nil
	}

	return out[0], nil
}

// namespaceFetchRows 取出rows的数据
func namespaceFetchRows(rows *sql.Rows) ([]*model.Namespace, error) {
	if rows == nil {
		return nil, nil
	}
	defer rows.Close()

	var (
		out  []*model.Namespace
		flag int
	)

	for rows.Next() {
		space := &model.Namespace{}
		err := rows.Scan(
			&space.Name,
			&space.Comment,
			&space.Token,
			&space.Owner,
			&flag,
			&space.CreateTime,
			&space.ModifyTime)
		if err != nil {
			log.Errorf("[Store][database] fetch namespace rows scan err: %s", err.Error())
			return nil, err
		}

		space.Valid = true
		if flag == 1 {
			space.Valid = false
		}

		out = append(out, space)
	}
	if err := rows.Err(); err != nil {
		log.Errorf("[Store][database] fetch namespace rows next err: %s", err.Error())
		return nil, err
	}

	return out, nil
}

// genNamespaceSelectSQL 生成namespace的查询语句
func genNamespaceSelectSQL() string {
	str := `select name, comment, token, owner, flag, ctime, mtime 
			from namespace `
	return str
}

// cleanNamespace clean真实的数据，只有flag=1的数据才可以清除
func cleanNamespace(tx *BaseTx, name string) error {
	str := "delete from namespace where name = $1 and flag = 1"
	stmt, err := tx.Prepare(str)
	if err != nil {
		log.Errorf("[Store][database] clean Prepare namespace(%s) err: %s", name, err.Error())
		return err
	}

	// 必须打印日志说明
	log.Infof("[Store][database] clean namespace(%s)", name)

	if _, err := stmt.Exec(name); err != nil {
		log.Infof("[Store][database] clean namespace(%s) err: %s", name, err.Error())
		return err
	}

	return nil
}
