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
)

func TestCreateRoutingConfig(t *testing.T) {
	obj := initConf()

	conf := &model.RoutingConfig{
		ID:        "1111",
		InBounds:  "2222",
		OutBounds: "3333",
		Revision:  "4444",
		Valid:     false,
		//CreateTime time.Time
		//ModifyTime time.Time
	}
	err := obj.routingConfigStore.CreateRoutingConfig(conf)
	fmt.Printf("err: %+v\n", err)
}

func TestUpdateRoutingConfig(t *testing.T) {
	obj := initConf()

	conf := &model.RoutingConfig{
		ID:        "1111",
		InBounds:  "2223",
		OutBounds: "3333",
		Revision:  "4444",
		Valid:     false,
		//CreateTime time.Time
		//ModifyTime time.Time
	}
	err := obj.routingConfigStore.UpdateRoutingConfig(conf)
	fmt.Printf("err: %+v\n", err)
}

func TestDeleteRoutingConfig(t *testing.T) {
	obj := initConf()
	err := obj.routingConfigStore.DeleteRoutingConfig("1111")
	fmt.Printf("err: %+v\n", err)
}

func TestGetRoutingConfigsForCache(t *testing.T) {
	obj := initConf()
	resp, err := obj.routingConfigStore.GetRoutingConfigsForCache(UnixSecondToTime(1686293155), true)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestGetRoutingConfigWithService(t *testing.T) {
	obj := initConf()
	resp, err := obj.routingConfigStore.GetRoutingConfigWithService("Test1", "2222")
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestGetRoutingConfigs(t *testing.T) {
	obj := initConf()
	filter := map[string]string{
		"routing_config.id": "1111",
		"in_bounds":         "2223",
	}
	cnt, resp, err := obj.routingConfigStore.GetRoutingConfigs(filter, 0, 10)
	fmt.Printf("cnt: %+v, resp: %+v, err: %+v\n", cnt, resp, err)
}

func TestCreateRoutingConfigV2(t *testing.T) {
	obj := initConf()

	conf := &model.RouterConfig{
		ID: "1111",
		// namespace router config owner namespace
		Namespace: "2222",
		// name router config name
		Name: "3333",
		// policy Rules
		Policy: "4444",
		// config Specific routing rules content
		Config: "5555",
		// enable Whether the routing rules are enabled
		Enable: true,
		// priority Rules priority
		Priority: 1,
		// revision Edition information of routing rules
		Revision: "6666",
		// Description Simple description of rules
		Description: "7777",
		// valid Whether the routing rules are valid and have not been deleted by logic
		Valid: true,
		//CreateTime time.Time `json:"ctime"`
		//ModifyTime time.Time `json:"mtime"`
		//EnableTime time.Time `json:"etime"`
	}
	err := obj.routingConfigStoreV2.CreateRoutingConfigV2(conf)
	fmt.Printf("err: %+v\n", err)
}

func TestUpdateRoutingConfigV2(t *testing.T) {
	obj := initConf()

	conf := &model.RouterConfig{
		ID: "1111",
		// namespace router config owner namespace
		Namespace: "2223",
		// name router config name
		Name: "3334",
		// policy Rules
		Policy: "4444",
		// config Specific routing rules content
		Config: "5555",
		// enable Whether the routing rules are enabled
		Enable: true,
		// priority Rules priority
		Priority: 1,
		// revision Edition information of routing rules
		Revision: "6666",
		// Description Simple description of rules
		Description: "7777",
		// valid Whether the routing rules are valid and have not been deleted by logic
		Valid: true,
		//CreateTime time.Time `json:"ctime"`
		//ModifyTime time.Time `json:"mtime"`
		//EnableTime time.Time `json:"etime"`
	}
	err := obj.routingConfigStoreV2.UpdateRoutingConfigV2(conf)
	fmt.Printf("err: %+v\n", err)
}

func TestEnableRouting(t *testing.T) {
	obj := initConf()

	conf := &model.RouterConfig{
		ID: "1111",
		// namespace router config owner namespace
		Namespace: "2223",
		// name router config name
		Name: "3334",
		// policy Rules
		Policy: "4444",
		// config Specific routing rules content
		Config: "5555",
		// enable Whether the routing rules are enabled
		Enable: true,
		// priority Rules priority
		Priority: 1,
		// revision Edition information of routing rules
		Revision: "6666",
		// Description Simple description of rules
		Description: "7777",
		// valid Whether the routing rules are valid and have not been deleted by logic
		Valid: true,
		//CreateTime time.Time `json:"ctime"`
		//ModifyTime time.Time `json:"mtime"`
		//EnableTime time.Time `json:"etime"`
	}
	err := obj.routingConfigStoreV2.EnableRouting(conf)
	fmt.Printf("err: %+v\n", err)
}

func TestGetRoutingConfigsV2ForCache(t *testing.T) {
	obj := initConf()
	resp, err := obj.routingConfigStoreV2.GetRoutingConfigsV2ForCache(UnixSecondToTime(1686293155), true)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestGetRoutingConfigV2WithID(t *testing.T) {
	obj := initConf()
	resp, err := obj.routingConfigStoreV2.GetRoutingConfigV2WithID("1111")
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}
