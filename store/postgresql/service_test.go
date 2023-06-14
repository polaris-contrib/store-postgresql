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

func TestAddService(t *testing.T) {
	obj := initConf()

	modelService := &model.Service{
		ID:          "1111",
		Name:        "Test1",
		Namespace:   "2222",
		Business:    "33333",
		Ports:       "4444",
		Meta:        map[string]string{"a": "b"},
		Comment:     "Polaris-test",
		Department:  "555",
		CmdbMod1:    "666",
		CmdbMod2:    "777",
		CmdbMod3:    "888",
		Token:       "2d1bfe5d12e04d54b8ee69e62494c7fe",
		Owner:       "polaris",
		Revision:    "999",
		Reference:   "101010",
		ReferFilter: "111111",
		PlatformID:  "121212",
		Valid:       false,
		CreateTime:  time.Now(),
		ModifyTime:  time.Now(),
		Mtime:       111,
		Ctime:       222,
	}
	err := obj.serviceStore.AddService(modelService)
	fmt.Println("err: ", err)
}

func TestDeleteService(t *testing.T) {
	obj := initConf()

	modelService := &model.Service{
		ID:        "1111",
		Name:      "Test1",
		Namespace: "2222",
	}
	err := obj.serviceStore.DeleteService(modelService.ID, modelService.Name, modelService.Namespace)
	fmt.Println("err: ", err)
}

func TestDeleteServiceAlias(t *testing.T) {
	obj := initConf()

	modelService := &model.Service{
		ID:        "1111",
		Name:      "Test1",
		Namespace: "2222",
	}
	err := obj.serviceStore.DeleteServiceAlias(modelService.Name, modelService.Namespace)
	fmt.Println("err: ", err)
}

func TestUpdateServiceAlias(t *testing.T) {
	obj := initConf()

	modelService := &model.Service{
		ID:        "1111",
		Name:      "Test1",
		Namespace: "2222",
		Revision:  "3333",
		Reference: "1111",
		Owner:     "5555",
	}
	err := obj.serviceStore.UpdateServiceAlias(modelService, true)
	fmt.Println("err: ", err)
}

func TestUpdateService(t *testing.T) {
	obj := initConf()

	modelService := &model.Service{
		ID:          "1111",
		Name:        "Test1",
		Namespace:   "2222",
		Business:    "33333",
		Ports:       "4444",
		Meta:        map[string]string{"a": "b"},
		Comment:     "Polaris-test",
		Department:  "555",
		CmdbMod1:    "666",
		CmdbMod2:    "777",
		CmdbMod3:    "888",
		Token:       "2d1bfe5d12e04d54b8ee69e62494c7fe",
		Owner:       "polaris",
		Revision:    "999",
		Reference:   "101010",
		ReferFilter: "111111",
		PlatformID:  "121212",
		Valid:       false,
		CreateTime:  time.Now(),
		ModifyTime:  time.Now(),
		Mtime:       111,
		Ctime:       222,
	}
	err := obj.serviceStore.UpdateService(modelService, true)
	fmt.Println("err: ", err)
}

func TestUpdateServiceToken(t *testing.T) {
	obj := initConf()

	modelService := &model.Service{
		ID:          "1111",
		Name:        "Test1",
		Namespace:   "2222",
		Business:    "33333",
		Ports:       "4444",
		Meta:        map[string]string{"a": "b"},
		Comment:     "Polaris-test",
		Department:  "555",
		CmdbMod1:    "666",
		CmdbMod2:    "777",
		CmdbMod3:    "888",
		Token:       "2d1bfe5d12e04d54b8ee69e62494c7fe",
		Owner:       "polaris",
		Revision:    "999",
		Reference:   "101010",
		ReferFilter: "111111",
		PlatformID:  "121212",
		Valid:       false,
		CreateTime:  time.Now(),
		ModifyTime:  time.Now(),
		Mtime:       111,
		Ctime:       222,
	}
	err := obj.serviceStore.UpdateServiceToken(modelService.ID, modelService.Token, modelService.Revision)
	fmt.Println("err: ", err)
}

func TestGetService(t *testing.T) {
	obj := initConf()
	resp, err := obj.serviceStore.GetService("Test1", "2222")
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestGetSourceServiceToken(t *testing.T) {
	obj := initConf()
	resp, err := obj.serviceStore.GetSourceServiceToken("Test1", "2222")
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestGetServiceByID(t *testing.T) {
	obj := initConf()
	resp, err := obj.serviceStore.GetServiceByID("1111")
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestGetServices(t *testing.T) {
	obj := initConf()
	serviceFilters := map[string]string{}
	serviceMetas := map[string]string{
		"mkey":   "111",
		"mvalue": "222",
	}
	instanceFilters := &store.InstanceArgs{
		Hosts: []string{"127.0.0.1", "127.0.0.2"},
		Ports: []uint32{80, 1230},
	}
	var (
		offset uint32 = 0
		limit  uint32 = 10
	)

	cnt, resp, err := obj.serviceStore.GetServices(serviceFilters, serviceMetas, instanceFilters, offset, limit)
	fmt.Printf("cnt: %+v, resp: %+v, err: %+v\n", cnt, resp, err)
}

func TestGetMoreServices(t *testing.T) {
	obj := initConf()
	curTime := UnixSecondToTime(1685779323)
	resp, err := obj.serviceStore.GetMoreServices(curTime, false, true, false)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestGetSystemServices(t *testing.T) {
	obj := initConf()
	resp, err := obj.serviceStore.GetSystemServices()
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestGetServiceAliases(t *testing.T) {
	obj := initConf()
	serviceFilters := map[string]string{}
	cnt, resp, err := obj.serviceStore.GetServiceAliases(serviceFilters, 0, 10)
	fmt.Printf("cnt: %+v, resp: %+v, err: %+v\n", cnt, resp, err)
}
