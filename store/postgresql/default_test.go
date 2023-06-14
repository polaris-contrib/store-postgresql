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
	"time"

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
			/*"slave": map[interface{}]interface{}{
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
			},*/
		},
	}
	obj := &PostgresqlStore{}
	err := obj.Initialize(conf)
	fmt.Println(err)

	return obj
}

func TestNewBaseDB(t *testing.T) {
	obj := initConf()
	fmt.Println("obj: ", obj)
}

func TestCreateTransaction(t *testing.T) {
	obj := initConf()

	tran, err := obj.CreateTransaction()

	fmt.Println("tran: ", tran, err)
}

func TestAddNamespace(t *testing.T) {
	obj := initConf()

	modelNamespace := &model.Namespace{
		Name:    "Test1",
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

func TestUpdateNamespace(t *testing.T) {
	obj := initConf()

	modelNamespace := &model.Namespace{
		Name:    "Test",
		Comment: "Polaris-test1",
		Token:   "2d1bfe5d12e04d54b8ee69e62494c7fe",
		Owner:   "polaris",
		Valid:   false,
		//CreateTime time.Time
		//ModifyTime time.Time
	}
	err := obj.namespaceStore.UpdateNamespace(modelNamespace)

	fmt.Printf("namespace: %+v\n", err)
}

func TestUpdateNamespaceToken(t *testing.T) {
	obj := initConf()

	err := obj.UpdateNamespaceToken("Test", "2d1bfe5d12e04d54b8ee69e62494c7fr")

	fmt.Printf("response: %+v\n", err)
}

func TestGetNamespace(t *testing.T) {
	obj := initConf()

	response, err := obj.GetNamespace("Test")

	fmt.Printf("res: %+v, err: %+v\n", response.Name, err)
}

func TestGetNamespaces(t *testing.T) {
	obj := initConf()

	filter := map[string][]string{
		"name": {"Test", "default"},
	}
	response, cnt, err := obj.GetNamespaces(filter, 0, 10)

	fmt.Printf("res: %+v, cnt: %+v, err: %+v\n", response, cnt, err)
}

func TestGetMoreNamespaces(t *testing.T) {
	obj := initConf()

	response, err := obj.GetMoreNamespaces(time.Time{})

	for _, resp := range response {
		fmt.Printf("res: %+v, err: %+v\n", resp.Name, err)
	}

}
