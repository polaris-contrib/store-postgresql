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
)

func TestGenNextL5Sid(t *testing.T) {
	obj := initConf()
	var id uint32 = 1

	resp, err := obj.l5Store.GenNextL5Sid(id)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestGetMoreL5Routes(t *testing.T) {
	obj := initConf()
	var id uint32 = 1

	resp, err := obj.l5Store.GetMoreL5Routes(id)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}
