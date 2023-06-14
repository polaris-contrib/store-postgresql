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
)

func TestCreateLeaderElection(t *testing.T) {
	obj := initConf()

	for i := 0; i < 2; i++ {
		//go func() {
		key := fmt.Sprintf("test%d", i)
		err := obj.adminStore.StartLeaderElection(key)
		fmt.Printf("err: %+v\n", err)
		//}()
	}
}

func TestCheckMtimeExpired(t *testing.T) {
	obj := initConf()

	key := fmt.Sprintf("test%d", 1)
	err := obj.adminStore.StartLeaderElection(key)
	fmt.Printf("err: %+v\n", err)

	select {}
}

func TestBatchCleanDeletedInstances(t *testing.T) {
	obj := initConf()

	resp, err := obj.adminStore.BatchCleanDeletedInstances(10*time.Minute, 5)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)

	select {}
}

func TestGetUnHealthyInstances(t *testing.T) {
	obj := initConf()

	resp, err := obj.adminStore.GetUnHealthyInstances(10*time.Minute, 5)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)

	select {}
}

func TestBatchCleanDeletedClients(t *testing.T) {
	obj := initConf()

	resp, err := obj.adminStore.BatchCleanDeletedClients(10*time.Minute, 5)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}
