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
)

func TestCreateGrayResourceTx(t *testing.T) {
	obj := initConf()
	tx, err := obj.StartTx()
	if err != nil {
		return
	}

	grayResource := &model.GrayResource{
		Name:       "Name1",
		MatchRule:  "MatchRule1",
		Valid:      false,
		CreateBy:   "2024-10-01 11:11:11",
		ModifyBy:   "2023-10-01 11:11:12",
		CreateTime: UnixSecondToTime(5),
		ModifyTime: UnixSecondToTime(6),
	}

	err = obj.grayStore.CreateGrayResourceTx(tx, grayResource)
	fmt.Printf("err: %+v\n", err)
}
