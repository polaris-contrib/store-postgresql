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

func TestFindConfigFileReleaseByModifyTimeAfter(t *testing.T) {
	obj := initConf()

	ret, err := obj.FindConfigFileReleaseByModifyTimeAfter(time.Time{})
	fmt.Printf("ret: %+v, err: %+v\n", ret, err)
}

func TestQueryConfigFilesByGroup(t *testing.T) {
	obj := initConf()

	cnt, ret, err := obj.QueryConfigFilesByGroup("namespace", "group", 0, 10)
	fmt.Printf("cnt: %+v, ret: %+v, err: %+v\n", cnt, ret, err)
}

func TestQueryConfigFileReleaseHistories(t *testing.T) {
	obj := initConf()

	cnt, ret, err := obj.QueryConfigFileReleaseHistories("", "", "", 0, 10, 0)
	fmt.Printf("cnt: %+v, ret: %+v, err: %+v\n", cnt, ret, err)
}
