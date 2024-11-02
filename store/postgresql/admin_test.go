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

const (
	TestElectKey = "test-key"
)

// mockgen -source=postgresql/admin.go -destination=mock/admin_mock.go -package=mock

func TestStartLeaderElection(t *testing.T) {
	obj := initConf()

	for i := 0; i < 2; i++ {
		key := fmt.Sprintf("test%d", i)
		err := obj.adminStore.StartLeaderElection(key)
		fmt.Printf("err: %+v\n", err)
	}
}

func TestStopLeaderElections(t *testing.T) {
	obj := initConf()
	obj.adminStore.StopLeaderElections()
}

func TestIsLeader(t *testing.T) {
	obj := initConf()
	ret := obj.adminStore.IsLeader("test1")
	fmt.Println("ret", ret)
}

func TestListLeaderElections(t *testing.T) {
	obj := initConf()
	list, err := obj.ListLeaderElections()
	fmt.Println("list", list, "err", err)
}

func TestReleaseLeaderElection(t *testing.T) {
	obj := initConf()
	err := obj.adminStore.ReleaseLeaderElection("test1")
	fmt.Printf("resp,err: %+v\n", err)
}

func TestBatchCleanDeletedInstances(t *testing.T) {
	obj := initConf()

	resp, err := obj.adminStore.BatchCleanDeletedInstances(10*time.Minute, 5)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestGetUnHealthyInstances(t *testing.T) {
	obj := initConf()

	resp, err := obj.adminStore.GetUnHealthyInstances(10*time.Minute, 5)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestBatchCleanDeletedClients(t *testing.T) {
	obj := initConf()

	resp, err := obj.adminStore.BatchCleanDeletedClients(10*time.Minute, 5)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}
