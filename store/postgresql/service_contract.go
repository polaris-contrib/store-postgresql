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
	"fmt"
	"strconv"
	"time"

	"github.com/polarismesh/specification/source/go/api/v1/service_manage"
	"go.uber.org/zap"

	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
)

var (
	// contractAvaiableFilter 允许查询的字段
	contractAvaiableFilter = map[string]struct{}{
		"name":      {},
		"service":   {},
		"namespace": {},
		"version":   {},
		"protocol":  {},
	}
)

type serviceContractStore struct {
	master *BaseDB
	slave  *BaseDB
}

// CreateServiceContract 创建服务契约
func (s *serviceContractStore) CreateServiceContract(contract *model.ServiceContract) error {
	addSql := "INSERT INTO service_contract(id, name, namespace, service, protocol, version, revision, flag, content, ctime, mtime) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, 0, $8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"

	_, err := s.master.Exec(addSql, contract.ID, contract.Name, contract.Namespace, contract.Service,
		contract.Protocol, contract.Version, contract.Revision, contract.Content)
	return store.Error(err)
}

// UpdateServiceContract 更新服务契约信息
func (s *serviceContractStore) UpdateServiceContract(contract *model.ServiceContract) error {
	updateSql := "UPDATE service_contract SET content = $1, revision = $2, mtime = CURRENT_TIMESTAMP WHERE id = $3"
	_, err := s.master.Exec(updateSql, contract.Content, contract.Revision, contract.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteServiceContract 删除服务契约 删除该版本的全部数据
func (s *serviceContractStore) DeleteServiceContract(contract *model.ServiceContract) error {
	return s.master.processWithTransaction("DeleteServiceContract", func(tx *BaseTx) error {
		deleteSql := "UPDATE service_contract SET flag = 1, mtime = CURRENT_TIMESTAMP WHERE id = $1"
		if _, err := tx.Exec(deleteSql, contract.ID); err != nil {
			log.Errorf("[Store][database] all delete service contract err: %s", err.Error())
			return err
		}

		deleteDetailSql := "DELETE FROM service_contract_detail WHERE contract_id = $1"
		if _, err := tx.Exec(deleteDetailSql, contract.ID); err != nil {
			log.Errorf("[Store][database] all delete service contract detail err: %s", err.Error())
			return err
		}

		return tx.Commit()
	})
}

// GetServiceContract 通过ID查询服务契约数据
func (s *serviceContractStore) GetServiceContract(id string) (data *model.EnrichServiceContract, err error) {
	querySql := "SELECT id, name, namespace, service, protocol, version, revision, flag, content, " +
		"ctime, mtime FROM service_contract WHERE flag = 0 AND id = $1"

	args := []interface{}{id}
	rows, err := s.master.Query(querySql, args...)
	if err != nil {
		log.Error("[Store][Contract] list contract ", zap.String("query", querySql), zap.Any("args", args))
		return nil, store.Error(err)
	}
	defer func() {
		_ = rows.Close()
	}()

	list := make([]*model.ServiceContract, 0)
	for rows.Next() {
		var (
			flag               int64
			ctimeStr, mtimeStr string
		)
		contract := model.ServiceContract{}
		if scanErr := rows.Scan(&contract.ID, &contract.Name, &contract.Namespace, &contract.Service,
			&contract.Protocol, &contract.Version, &contract.Revision, &flag,
			&contract.Content, &ctimeStr, &mtimeStr); scanErr != nil {
			log.Errorf("[Store][Contract] fetch contract rows scan err: %s", scanErr.Error())
			return nil, store.Error(scanErr)
		}

		contract.Valid = flag == 0

		// 将字符串转换为int64
		ctimeFloat, err := strconv.ParseFloat(ctimeStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse create_time: %v", err)
		}
		mtimeFloat, err := strconv.ParseFloat(mtimeStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse modify_time: %v", err)
		}
		contract.CreateTime = time.Unix(int64(ctimeFloat), 0)
		contract.ModifyTime = time.Unix(int64(mtimeFloat), 0)

		list = append(list, &contract)
	}

	if len(list) == 0 {
		return nil, nil
	}
	return &model.EnrichServiceContract{
		ServiceContract: list[0],
	}, nil
}

// AddServiceContractInterfaces 创建服务契约API接口
func (s *serviceContractStore) AddServiceContractInterfaces(contract *model.EnrichServiceContract) error {
	return s.master.processWithTransaction("AddServiceContractDetail", func(tx *BaseTx) error {
		updateRevision := "UPDATE service_contract SET revision = $1, mtime = CURRENT_TIMESTAMP WHERE id = $2"
		if _, err := tx.Exec(updateRevision, contract.Revision, contract.ID); err != nil {
			log.Errorf("[Store][database] update service contract revision err: %s", err.Error())
			return err
		}

		// 新增批量数据
		for _, item := range contract.Interfaces {
			addSql := "INSERT INTO service_contract_detail(id, contract_id, name, method, path, content, revision, flag, ctime, mtime, source) " +
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $9) " +
				"ON CONFLICT (id) DO UPDATE SET contract_id = excluded.contract_id, name = excluded.name, method = excluded.method, " +
				"path = excluded.path, content = excluded.content, revision = excluded.revision, flag = excluded.flag, " +
				"ctime = excluded.ctime, mtime = excluded.mtime, source = excluded.source"

			if _, err := tx.Exec(addSql, item.ID, contract.ID, item.Name, item.Method, item.Path,
				item.Content, item.Revision, 0, int(item.Source)); err != nil {
				log.Errorf("[Store][database] add service contract detail err: %s", err.Error())
				return err
			}
		}
		return tx.Commit()
	})
}

// AppendServiceContractInterfaces 追加服务契约API接口
func (s *serviceContractStore) AppendServiceContractInterfaces(contract *model.EnrichServiceContract) error {
	return s.master.processWithTransaction("AppendServiceContractDetail", func(tx *BaseTx) error {
		updateRevision := "UPDATE service_contract SET revision = $1, mtime = CURRENT_TIMESTAMP WHERE id = $2"
		if _, err := tx.Exec(updateRevision, contract.Revision, contract.ID); err != nil {
			log.Errorf("[Store][database] update service contract revision err: %s", err.Error())
			return err
		}

		for _, item := range contract.Interfaces {
			addSql := "INSERT INTO service_contract_detail(id, contract_id, name, method, path, content, revision, flag, ctime, mtime, source) " +
				"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $9) " +
				"ON CONFLICT (id) DO NOTHING"

			if _, err := tx.Exec(addSql, item.ID, contract.ID, item.Name, item.Method, item.Path,
				item.Content, item.Revision, 0, int(item.Source)); err != nil {
				log.Errorf("[Store][database] append service contract detail err: %s", err.Error())
				return err
			}
		}
		return tx.Commit()
	})
}

// DeleteServiceContractInterfaces 删除服务契约API接口
func (s *serviceContractStore) DeleteServiceContractInterfaces(contract *model.EnrichServiceContract) error {
	return s.master.processWithTransaction("DeleteServiceContractDetail", func(tx *BaseTx) error {
		updateRevision := "UPDATE service_contract SET revision = $1, mtime = CURRENT_TIMESTAMP WHERE id = $2"
		if _, err := tx.Exec(updateRevision, contract.Revision, contract.ID); err != nil {
			log.Errorf("[Store][database] update service contract revision err: %s", err.Error())
			return err
		}

		for _, item := range contract.Interfaces {
			deleteSql := "DELETE FROM service_contract_detail WHERE contract_id = $1 AND method = $2 AND path = $3 AND name = $4"

			if _, err := tx.Exec(deleteSql, item.ContractID, item.Method, item.Path, item.Name); err != nil {
				log.Errorf("[Store][database] delete service contract detail err: %s", err.Error())
				return err
			}
		}
		return tx.Commit()
	})
}

// GetMoreServiceContracts 查询服务契约数据
func (s *serviceContractStore) GetMoreServiceContracts(firstUpdate bool, mtime time.Time) ([]*model.EnrichServiceContract, error) {
	querySql := "SELECT id, name, namespace, service, protocol, version, revision, flag, content, " +
		"ctime, mtime FROM service_contract WHERE mtime >= $1 "
	if firstUpdate {
		mtime = time.Unix(0, 1)
		querySql += " AND flag = 0 "
	}

	tx, err := s.slave.Begin()
	if err != nil {
		log.Error("[Store][Contract] list contract for cache when begin tx", zap.Error(err))
		return nil, store.Error(err)
	}
	defer func() {
		_ = tx.Commit()
	}()

	rows, err := tx.Query(querySql, mtime)
	if err != nil {
		log.Error("[Store][Contract] list contract for cache when query", zap.Error(err))
		return nil, store.Error(err)
	}
	defer func() {
		_ = rows.Close()
	}()

	list := make([]*model.EnrichServiceContract, 0)
	for rows.Next() {
		var (
			flag               float64
			ctimeStr, mtimeStr string
		)
		contract := &model.ServiceContract{}
		if scanErr := rows.Scan(&contract.ID, &contract.Name, &contract.Namespace, &contract.Service,
			&contract.Protocol, &contract.Version, &contract.Revision, &flag,
			&contract.Content, &ctimeStr, &mtimeStr); scanErr != nil {
			log.Error("[Store][Contract] fetch contract rows scan err: %s", zap.Error(scanErr))
			return nil, store.Error(scanErr)
		}

		contract.Valid = flag == 0
		// 将字符串转换为int64
		ctimeFloat, err := strconv.ParseFloat(ctimeStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse create_time: %v", err)
		}
		mtimeFloat, err := strconv.ParseFloat(mtimeStr, 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse modify_time: %v", err)
		}

		contract.CreateTime = time.Unix(int64(ctimeFloat), 0)
		contract.ModifyTime = time.Unix(int64(mtimeFloat), 0)

		list = append(list, &model.EnrichServiceContract{
			ServiceContract: contract,
		})
	}

	contractDetailMap := map[string][]*model.InterfaceDescriptor{}
	if len(list) > 0 {
		queryDetailSql := "SELECT sd.id, sd.contract_id, sd.name, sd.method, sd.path, sd.content, sd.revision, " +
			"sd.ctime, sd.mtime, COALESCE(sd.source, 1) " +
			"FROM service_contract_detail sd LEFT JOIN service_contract sc ON sd.contract_id = sc.id " +
			"WHERE sc.mtime >= $1"
		detailRows, err := tx.Query(queryDetailSql, mtime)
		if err != nil {
			log.Error("[Store][Contract] list contract detail", zap.String("query sql", queryDetailSql), zap.Error(err))
			return nil, store.Error(err)
		}
		defer func() {
			_ = detailRows.Close()
		}()
		for detailRows.Next() {
			var (
				flag, source       float64
				ctimeStr, mtimeStr string
			)
			detailItem := &model.InterfaceDescriptor{}
			if scanErr := detailRows.Scan(
				&detailItem.ID, &detailItem.ContractID, &detailItem.Name, &detailItem.Method,
				&detailItem.Path, &detailItem.Content, &detailItem.Revision,
				&ctimeStr, &mtimeStr, &source,
			); scanErr != nil {
				log.Error("[Store][Contract] fetch contract detail rows scan", zap.Error(scanErr))
				return nil, store.Error(scanErr)
			}

			detailItem.Valid = flag == 0
			// 将字符串转换为int64
			ctimeFloat, err := strconv.ParseFloat(ctimeStr, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse create_time: %v", err)
			}
			mtimeFloat, err := strconv.ParseFloat(mtimeStr, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse modify_time: %v", err)
			}
			detailItem.CreateTime = time.Unix(int64(ctimeFloat), 0)
			detailItem.ModifyTime = time.Unix(int64(mtimeFloat), 0)
			switch int(source) {
			case 2:
				detailItem.Source = service_manage.InterfaceDescriptor_Client
			default:
				detailItem.Source = service_manage.InterfaceDescriptor_Manual
			}

			if _, ok := contractDetailMap[detailItem.ContractID]; !ok {
				contractDetailMap[detailItem.ContractID] = make([]*model.InterfaceDescriptor, 0, 4)
			}
			contractDetailMap[detailItem.ContractID] = append(contractDetailMap[detailItem.ContractID], detailItem)
		}

		for _, item := range list {
			methods := contractDetailMap[item.ID]
			item.Interfaces = methods
			item.Format()
		}
	}
	return list, nil
}
