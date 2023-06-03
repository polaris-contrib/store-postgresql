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
	"testing"

	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
)

func initConf() *PostgresqlStore {
	conf := &store.Config{
		Name: "Postgresql",
		Option: map[string]interface{}{
			"master": map[interface{}]interface{}{
				"dbType": "postgres",
				"dbUser": "postgres",
				"dbPwd":  "aaaaaa",
				"dbAddr": "192.168.31.19",
				"dbPort": "5432",
				"dbName": "polaris_server",

				"maxOpenConns":     10,
				"maxIdleConns":     10,
				"connMaxLifetime":  10,
				"txIsolationLevel": 2,
			},
			"slave": map[interface{}]interface{}{
				"dbType": "postgres",
				"dbUser": "postgres",
				"dbPwd":  "aaaaaa",
				"dbAddr": "192.168.31.19",
				"dbPort": "5432",
				"dbName": "polaris_server",

				"maxOpenConns":     10,
				"maxIdleConns":     10,
				"connMaxLifetime":  10,
				"txIsolationLevel": 2,
			},
		},
	}
	obj := &PostgresqlStore{}
	err := obj.Initialize(conf)
	fmt.Println(err)

	return obj
}

func TestCreateTransaction(t *testing.T) {
	obj := initConf()

	tran, err := obj.CreateTransaction()

	fmt.Println("tran: ", tran, err)
}

func TestAddNamespace(t *testing.T) {
	obj := initConf()

	modelNamespace := &model.Namespace{
		Name:    "Test",
		Comment: "Polaris-test",
		Token:   "2d1bfe5d12e04d54b8ee69e62494c7fe",
		Owner:   "polaris",
		Valid:   false,
		//CreateTime time.Time
		//ModifyTime time.Time
	}
	err := obj.namespaceStore.AddNamespace(modelNamespace)

	fmt.Printf("namespace: %+v\n", err)
}
