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

// RoutingConfigStore的实现
type routingConfigStore struct {
	master *BaseDB
	slave  *BaseDB
}

// CreateRoutingConfig 新建RoutingConfig
func (rs *routingConfigStore) CreateRoutingConfig(conf *model.RoutingConfig) error {
	if conf.ID == "" || conf.Revision == "" {
		log.Errorf("[Store][database] create routing config missing service id or revision")
		return store.NewStatusError(store.EmptyParamsErr, "missing service id or revision")
	}
	if conf.InBounds == "" || conf.OutBounds == "" {
		log.Errorf("[Store][database] create routing config missing params")
		return store.NewStatusError(store.EmptyParamsErr, "missing some params")
	}

	return RetryTransaction("createRoutingConfig", func() error {
		return rs.master.processWithTransaction("createRoutingConfig", func(tx *BaseTx) error {
			// 新建之前，先清理老数据
			if err := cleanRoutingConfig(tx, conf.ID); err != nil {
				return store.Error(err)
			}

			// 服务配置的创建由外层进行服务的保护，这里不需要加锁
			str := `insert into routing_config(id, in_bounds, out_bounds, revision, ctime, mtime)
			values($1,$2,$3,$4,$5,$6)`
			stmt, err := tx.Prepare(str)
			if err != nil {
				return store.Error(err)
			}
			if _, err = stmt.Exec(conf.ID, conf.InBounds, conf.OutBounds, conf.Revision, GetCurrentTimeFormat(), GetCurrentTimeFormat()); err != nil {
				log.Errorf("[Store][database] create routing(%+v) err: %s", conf, err.Error())
				return store.Error(err)
			}

			if err := tx.Commit(); err != nil {
				log.Errorf("[Store][database] fail to create routing commit tx, rule(%+v) commit tx err: %s",
					conf, err.Error())
				return err
			}
			return nil
		})
	})
}

// UpdateRoutingConfig 更新
func (rs *routingConfigStore) UpdateRoutingConfig(conf *model.RoutingConfig) error {
	if conf.ID == "" || conf.Revision == "" {
		log.Errorf("[Store][database] update routing config missing service id or revision")
		return store.NewStatusError(store.EmptyParamsErr, "missing service id or revision")
	}
	if conf.InBounds == "" || conf.OutBounds == "" {
		log.Errorf("[Store][database] update routing config missing params")
		return store.NewStatusError(store.EmptyParamsErr, "missing some params")
	}
	return RetryTransaction("updateRoutingConfig", func() error {
		return rs.master.processWithTransaction("updateRoutingConfig", func(tx *BaseTx) error {
			str := `update routing_config set in_bounds = $1, out_bounds = $2, revision = $3, mtime = $4 where id = $5`
			stmt, err := tx.Prepare(str)
			if err != nil {
				return store.Error(err)
			}
			if _, err = stmt.Exec(conf.InBounds, conf.OutBounds, conf.Revision, GetCurrentTimeFormat(), conf.ID); err != nil {
				log.Errorf("[Store][database] update routing config(%+v) exec err: %s", conf, err.Error())
				return store.Error(err)
			}

			if err := tx.Commit(); err != nil {
				log.Errorf("[Store][database] fail to update routing commit tx, rule(%+v) commit tx err: %s",
					conf, err.Error())
				return err
			}
			return nil
		})
	})
}

// DeleteRoutingConfig 删除
func (rs *routingConfigStore) DeleteRoutingConfig(serviceID string) error {
	if serviceID == "" {
		log.Errorf("[Store][database] delete routing config missing service id")
		return store.NewStatusError(store.EmptyParamsErr, "missing service id")
	}
	return RetryTransaction("deleteRoutingConfig", func() error {
		return rs.master.processWithTransaction("deleteRoutingConfig", func(tx *BaseTx) error {
			stmt, err := tx.Prepare(`update routing_config set flag = 1, mtime = $1 where id = $2`)
			if err != nil {
				return store.Error(err)
			}
			if _, err = stmt.Exec(GetCurrentTimeFormat(), serviceID); err != nil {
				log.Errorf("[Store][database] delete routing config(%s) err: %s", serviceID, err.Error())
				return store.Error(err)
			}

			if err := tx.Commit(); err != nil {
				log.Errorf("[Store][database] fail to delete routing commit tx, rule(%s) commit tx err: %s",
					serviceID, err.Error())
				return err
			}
			return nil
		})
	})
}

// DeleteRoutingConfigTx 删除
func (rs *routingConfigStore) DeleteRoutingConfigTx(tx store.Tx, serviceID string) error {
	if tx == nil {
		return errors.New("transaction is nil")
	}

	if serviceID == "" {
		log.Errorf("[Store][database] delete routing config missing service id")
		return store.NewStatusError(store.EmptyParamsErr, "missing service id")
	}

	dbTx := tx.GetDelegateTx().(*BaseTx)

	stmt, err := dbTx.Prepare(`update routing_config set flag = 1, mtime = $1 where id = $2`)
	if err != nil {
		return store.Error(err)
	}
	if _, err = stmt.Exec(GetCurrentTimeFormat(), serviceID); err != nil {
		log.Errorf("[Store][database] delete routing config(%s) err: %s", serviceID, err.Error())
		return store.Error(err)
	}
	return nil
}

// GetRoutingConfigsForCache 缓存增量拉取
func (rs *routingConfigStore) GetRoutingConfigsForCache(mtime time.Time, firstUpdate bool) ([]*model.RoutingConfig, error) {
	str := `select id, in_bounds, out_bounds, revision,flag, ctime, mtime  
			from routing_config where mtime > $1`
	if firstUpdate {
		str += " and flag != 1"
	}
	rows, err := rs.slave.Query(str, mtime)
	if err != nil {
		log.Errorf("[Store][database] query routing configs with mtime err: %s", err.Error())
		return nil, err
	}
	out, err := fetchRoutingConfigRows(rows)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// GetRoutingConfigWithService 根据服务名+namespace获取对应的配置
func (rs *routingConfigStore) GetRoutingConfigWithService(name string, namespace string) (*model.RoutingConfig, error) {
	// 只查询到flag=0的数据
	str := `select routing_config.id, in_bounds, out_bounds, revision, flag, ctime, mtime  
			from (select id from service where name = $1 and namespace = $2) as service, routing_config 
			where service.id = routing_config.id and routing_config.flag = 0`
	rows, err := rs.master.Query(str, name, namespace)
	if err != nil {
		log.Errorf("[Store][database] query routing config with service(%s, %s) err: %s",
			name, namespace, err.Error())
		return nil, err
	}

	out, err := fetchRoutingConfigRows(rows)
	if err != nil {
		return nil, err
	}

	if len(out) == 0 {
		return nil, nil
	}

	return out[0], nil
}

// GetRoutingConfigWithID 根据服务ID获取对应的配置
func (rs *routingConfigStore) GetRoutingConfigWithID(id string) (*model.RoutingConfig, error) {
	str := `select routing_config.id, in_bounds, out_bounds, revision, flag, ctime, mtime
			from routing_config 
			where id = $1 and flag = 0`
	rows, err := rs.master.Query(str, id)
	if err != nil {
		log.Errorf("[Store][database] query routing with id(%s) err: %s", id, err.Error())
		return nil, err
	}

	out, err := fetchRoutingConfigRows(rows)
	if err != nil {
		return nil, err
	}

	if len(out) == 0 {
		return nil, nil
	}

	return out[0], nil
}

// GetRoutingConfigs 获取路由配置列表
func (rs *routingConfigStore) GetRoutingConfigs(filter map[string]string,
	offset uint32, limit uint32) (uint32, []*model.ExtendRoutingConfig, error) {

	var index = 1
	filterStr, args, index1 := genFilterRoutingConfigSQL(filter, index)
	index = index1
	countStr := genQueryRoutingConfigCountSQL() + filterStr
	var total uint32
	err := rs.master.QueryRow(countStr, args...).Scan(&total)
	switch {
	case err == sql.ErrNoRows:
		return 0, nil, nil
	case err != nil:
		log.Errorf("[Store][database] get routing config query count err: %s", err.Error())
		return 0, nil, err
	default:
	}

	str := genQueryRoutingConfigSQL() + filterStr + fmt.Sprintf(" order by routing_config.mtime desc limit $%d offset $%d", index, index+1)
	args = append(args, limit, offset)
	rows, err := rs.master.Query(str, args...)
	if err != nil {
		log.Errorf("[Store][database] get routing configs query err: %s", err.Error())
		return 0, nil, err
	}
	defer rows.Close()

	var out []*model.ExtendRoutingConfig
	for rows.Next() {
		var tmp model.ExtendRoutingConfig
		tmp.Config = &model.RoutingConfig{}
		err := rows.Scan(&tmp.ServiceName, &tmp.NamespaceName, &tmp.Config.ID,
			&tmp.Config.InBounds, &tmp.Config.OutBounds, &tmp.Config.CreateTime, &tmp.Config.ModifyTime)
		if err != nil {
			log.Errorf("[Store][database] query routing configs rows scan err: %s", err.Error())
			return 0, nil, err
		}

		out = append(out, &tmp)
	}
	if err := rows.Err(); err != nil {
		log.Errorf("[Store][database] query routing configs rows next err: %s", err.Error())
		return 0, nil, err
	}

	return total, out, nil
}

// cleanRoutingConfig 从数据库彻底清理路由配置
func cleanRoutingConfig(tx *BaseTx, serviceID string) error {
	stmt, err := tx.Prepare(`delete from routing_config where id = $1 and flag = 1`)
	if err != nil {
		return err
	}
	if _, err = stmt.Exec(serviceID); err != nil {
		log.Errorf("[Store][database] clean routing config(%s) err: %s", serviceID, err.Error())
		return err
	}

	return nil
}

// fetchRoutingConfigRows 读取数据库的数据，并且释放rows
func fetchRoutingConfigRows(rows *sql.Rows) ([]*model.RoutingConfig, error) {
	defer rows.Close()
	var out []*model.RoutingConfig
	for rows.Next() {
		var entry model.RoutingConfig
		var flag int
		err := rows.Scan(&entry.ID, &entry.InBounds, &entry.OutBounds, &entry.Revision,
			&flag, &entry.CreateTime, &entry.ModifyTime)
		if err != nil {
			log.Errorf("[database][store] fetch routing config scan err: %s", err.Error())
			return nil, err
		}

		entry.Valid = true
		if flag == 1 {
			entry.Valid = false
		}

		out = append(out, &entry)
	}
	if err := rows.Err(); err != nil {
		log.Errorf("[database][store] fetch routing config next err: %s", err.Error())
		return nil, err
	}

	return out, nil
}

// genQueryRoutingConfigSQL 查询路由配置的语句
func genQueryRoutingConfigSQL() string {
	str := `select name, namespace, routing_config.id, in_bounds, out_bounds,
			routing_config.ctime, routing_config.mtime  
			from routing_config, service 
			where routing_config.id = service.id 
			and routing_config.flag = 0`
	return str
}

// genQueryRoutingConfigCountSQL 获取路由配置指定过滤条件下的总条目数
func genQueryRoutingConfigCountSQL() string {
	str := `select count(*) from routing_config, service
			where routing_config.id = service.id 
			and routing_config.flag = 0`
	return str
}

// genFilterRoutingConfigSQL 生成过滤语句
func genFilterRoutingConfigSQL(filters map[string]string, index int) (string, []interface{}, int) {
	str := ""
	args := make([]interface{}, 0, len(filters))
	for key, value := range filters {
		str += fmt.Sprintf(" and %s = $%d ", key, index)
		index++
		args = append(args, value)
	}

	return str, args, index
}
