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
	"github.com/polarismesh/polaris/common/model"
	"testing"
	"time"
)

func TestGetCircuitBreakerRules(t *testing.T) {
	obj := initConf()

	filter := map[string]string{
		"brief": "true",
		"level": "2,1",
	}
	cnt, ret, err := obj.circuitBreakerStore.GetCircuitBreakerRules(filter, 0, 10)
	fmt.Printf("cnt: %+v, ret: %+v, err: %+v\n", cnt, ret, err)
}

func TestCreateCircuitBreakerRule(t *testing.T) {
	obj := initConf()

	reqModel := &model.CircuitBreakerRule{
		ID:           "ID",
		Name:         "Name",
		Namespace:    "Namespace",
		Description:  "Description",
		Level:        1,
		SrcService:   "SrcService",
		SrcNamespace: "SrcNamespace",
		DstService:   "DstService",
		DstNamespace: "DstNamespace",
		DstMethod:    "DstMethod",
		Rule:         "Rule",
		Revision:     "Revision",
		Enable:       true,
		Valid:        true,
		CreateTime:   time.Now(),
		ModifyTime:   time.Now(),
		EnableTime:   time.Now(),
	}

	err := obj.circuitBreakerStore.CreateCircuitBreakerRule(reqModel)

	fmt.Printf("createCircuitBreakerRule: %+v\n", err)
}

func TestUpdateCircuitBreakerRule(t *testing.T) {
	obj := initConf()

	reqModel := &model.CircuitBreakerRule{
		ID:   "ID",
		Name: "Name1",
		/*Namespace:    "Namespace",
		Description:  "Description",
		Level:        1,
		SrcService:   "SrcService",
		SrcNamespace: "SrcNamespace",
		DstService:   "DstService",
		DstNamespace: "DstNamespace",
		DstMethod:    "DstMethod",
		Rule:         "Rule",
		Revision:     "Revision",
		Enable:       true,
		Valid:        true,
		CreateTime:   time.Now(),
		ModifyTime:   time.Now(),
		EnableTime:   time.Now(),*/
	}

	err := obj.circuitBreakerStore.UpdateCircuitBreakerRule(reqModel)

	fmt.Printf("UpdateCircuitBreakerRule: %+v\n", err)
}

func TestEnableCircuitBreakerRule(t *testing.T) {
	obj := initConf()

	reqModel := &model.CircuitBreakerRule{
		ID:       "ID",
		Name:     "Name1",
		Enable:   false,
		Revision: "Revision1",
		/*Namespace:    "Namespace",
		Description:  "Description",
		Level:        1,
		SrcService:   "SrcService",
		SrcNamespace: "SrcNamespace",
		DstService:   "DstService",
		DstNamespace: "DstNamespace",
		DstMethod:    "DstMethod",
		Rule:         "Rule",
		Valid:        true,
		CreateTime:   time.Now(),
		ModifyTime:   time.Now(),
		EnableTime:   time.Now(),*/
	}

	err := obj.circuitBreakerStore.EnableCircuitBreakerRule(reqModel)

	fmt.Printf("EnableCircuitBreakerRule: %+v\n", err)
}
