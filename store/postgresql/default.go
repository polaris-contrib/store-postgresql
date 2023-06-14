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
	"errors"
	"fmt"

	"github.com/polarismesh/polaris/common/log"
	"github.com/polarismesh/polaris/plugin"
	"github.com/polarismesh/polaris/store"
)

const (
	// SystemNamespace system namespace
	SystemNamespace = "Polaris"
	// STORENAME database storage name
	STORENAME = "postgresqlStore"
	// DefaultConnMaxLifetime default maximum connection lifetime
	DefaultConnMaxLifetime = 60 * 30 // 默认是30分钟
	// emptyEnableTime 规则禁用时启用时间的默认值
	emptyEnableTime = "STR_TO_DATE('1980-01-01 00:00:01', '%Y-%m-%d %H:%i:%s')"
)

// PostgresqlStore 实现了Store接口
type PostgresqlStore struct {
	*namespaceStore
	// client info stores
	*clientStore

	// 服务注册发现、治理
	*serviceStore
	*instanceStore
	*routingConfigStore
	*l5Store
	*rateLimitStore
	*circuitBreakerStore
	*faultDetectRuleStore

	// 工具
	*toolStore

	// 鉴权模块相关
	*userStore
	*groupStore
	*strategyStore

	// 配置中心store
	*configFileGroupStore
	*configFileStore
	*configFileReleaseStore
	*configFileReleaseHistoryStore
	*configFileTagStore
	*configFileTemplateStore

	// v2 存储
	*routingConfigStoreV2

	// maintain store
	*adminStore

	// 主数据库，可以进行读写
	master *BaseDB
	// 对主数据库的事务操作，可读写
	masterTx *BaseDB
	// 备数据库，提供只读
	slave *BaseDB
	start bool
}

// Name 实现Name函数
func (p *PostgresqlStore) Name() string {
	return STORENAME
}

// Initialize 初始化函数
func (p *PostgresqlStore) Initialize(conf *store.Config) error {
	if p.start {
		return nil
	}

	masterConfig, slaveConfig, err := parseDatabaseConf(conf.Option)
	if err != nil {
		return err
	}
	master, err := NewBaseDB(masterConfig, plugin.GetParsePassword())
	if err != nil {
		return err
	}
	p.master = master

	masterTx, err := NewBaseDB(masterConfig, plugin.GetParsePassword())
	if err != nil {
		return err
	}
	p.masterTx = masterTx

	if slaveConfig != nil {
		log.Infof("[Store][database] use slave database config: %+v", slaveConfig)
		slave, err := NewBaseDB(slaveConfig, plugin.GetParsePassword())
		if err != nil {
			return err
		}
		p.slave = slave
	}
	// 如果slave为空，意味着slaveConfig为空，用master数据库替代
	if p.slave == nil {
		p.slave = p.master
	}

	log.Infof("[Store][database] connect the database successfully")

	p.start = true

	p.newStore()

	return nil
}

// Destroy 退出函数
func (p *PostgresqlStore) Destroy() error {
	p.start = false

	if p.master != nil {
		_ = p.master.Close()
	}
	if p.masterTx != nil {
		_ = p.masterTx.Close()
	}
	if p.slave != nil {
		_ = p.slave.Close()
	}

	if p.adminStore != nil {
		p.adminStore.StopLeaderElections()
	}

	p.master = nil
	p.masterTx = nil
	p.slave = nil

	return nil
}

// CreateTransaction 创建一个事务
func (p *PostgresqlStore) CreateTransaction() (store.Transaction, error) {
	// 每次创建事务前，还是需要ping一下
	_ = p.masterTx.Ping()

	nt := &transaction{}
	tx, err := p.masterTx.Begin()
	if err != nil {
		log.Errorf("[Store][database] database begin err: %s", err.Error())
		return nil, err
	}

	nt.tx = tx

	return nt, nil
}

func (p *PostgresqlStore) StartTx() (store.Tx, error) {
	tx, err := p.masterTx.Begin()
	if err != nil {
		return nil, err
	}
	return NewSqlDBTx(tx), nil
}

func buildEtimeStr(enable bool) string {
	etimeStr := GetCurrentTimeFormat()
	if !enable {
		etimeStr = emptyEnableTime
	}
	return etimeStr
}

// parseDatabaseConf 解析数据库配置
func parseDatabaseConf(opt map[string]interface{}) (*dbConfig, *dbConfig, error) {
	// 必填
	masterEnter, ok := opt["master"]
	if !ok || masterEnter == nil {
		return nil, nil, errors.New("database master db config is missing")
	}
	masterConfig, err := parseStoreConfig(masterEnter)
	if err != nil {
		return nil, nil, err
	}

	// 只读数据库可选
	slaveEntry, ok := opt["slave"]
	if !ok || slaveEntry == nil {
		return masterConfig, nil, nil
	}
	slaveConfig, err := parseStoreConfig(slaveEntry)
	if err != nil {
		return nil, nil, err
	}

	return masterConfig, slaveConfig, nil
}

// parseStoreConfig 解析store的配置
func parseStoreConfig(opts interface{}) (*dbConfig, error) {
	obj, _ := opts.(map[interface{}]interface{})

	needCheckFields := map[string]string{"dbType": "", "dbUser": "", "dbPwd": "", "dbAddr": "", "dbPort": "", "dbName": ""}

	for key := range needCheckFields {
		val, ok := obj[key]
		if !ok {
			return nil, fmt.Errorf("config Plugin %s:%s type must be string", STORENAME, key)
		}

		needCheckFields[key] = fmt.Sprintf("%v", val)
	}

	c := &dbConfig{
		dbType: needCheckFields["dbType"],
		dbUser: needCheckFields["dbUser"],
		dbPwd:  needCheckFields["dbPwd"],
		dbAddr: needCheckFields["dbAddr"],
		dbPort: needCheckFields["dbPort"],
		dbName: needCheckFields["dbName"],
	}
	if maxOpenConns, _ := obj["maxOpenConns"].(int); maxOpenConns > 0 {
		c.maxOpenConns = maxOpenConns
	}
	if maxIdleConns, _ := obj["maxIdleConns"].(int); maxIdleConns > 0 {
		c.maxIdleConns = maxIdleConns
	}
	c.connMaxLifetime = DefaultConnMaxLifetime
	if connMaxLifetime, _ := obj["connMaxLifetime"].(int); connMaxLifetime > 0 {
		c.connMaxLifetime = connMaxLifetime
	}

	return c, nil
}

// newStore 初始化子类
func (p *PostgresqlStore) newStore() {
	p.namespaceStore = &namespaceStore{master: p.master, slave: p.slave}

	p.serviceStore = &serviceStore{master: p.master, slave: p.slave}

	p.instanceStore = &instanceStore{master: p.master, slave: p.slave}

	p.routingConfigStore = &routingConfigStore{master: p.master, slave: p.slave}

	p.l5Store = &l5Store{master: p.master, slave: p.slave}

	p.rateLimitStore = &rateLimitStore{master: p.master, slave: p.slave}

	p.circuitBreakerStore = &circuitBreakerStore{master: p.master, slave: p.slave}

	p.toolStore = &toolStore{db: p.master}

	p.userStore = &userStore{master: p.master, slave: p.slave}

	p.groupStore = &groupStore{master: p.master, slave: p.slave}

	p.strategyStore = &strategyStore{master: p.master, slave: p.slave}

	p.faultDetectRuleStore = &faultDetectRuleStore{master: p.master, slave: p.slave}

	p.configFileGroupStore = &configFileGroupStore{master: p.master, slave: p.slave}

	p.configFileStore = &configFileStore{master: p.master, slave: p.slave}

	p.configFileReleaseStore = &configFileReleaseStore{db: p.master, slave: p.slave}

	p.configFileReleaseHistoryStore = &configFileReleaseHistoryStore{db: p.master}

	p.configFileTagStore = &configFileTagStore{db: p.master}

	p.configFileTemplateStore = &configFileTemplateStore{db: p.master}

	p.clientStore = &clientStore{master: p.master, slave: p.slave}

	p.routingConfigStoreV2 = &routingConfigStoreV2{master: p.master, slave: p.slave}

	p.adminStore = newAdminStore(p.master)
}

func init() {
	s := &PostgresqlStore{}
	err := store.RegisterStore(s)
	if err != nil {
		log.Errorf("[Store][database] RegisterStore err: %+v", err)
	}
}
