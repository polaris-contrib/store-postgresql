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
	apimodel "github.com/polarismesh/specification/source/go/api/v1/model"
	"testing"
)

func TestLockConfigFile(t *testing.T) {
	obj := initConf()
	tx, err := obj.StartTx()
	if err != nil {
		return
	}

	key := &model.ConfigFileKey{
		Name:      "Name1",
		Namespace: "Namespace1",
		Group:     "Group1",
	}
	ret, err := obj.configFileStore.LockConfigFile(tx, key)
	fmt.Printf("ret: %+v, err: %+v", ret, err)
}

func TestCreateConfigFileTx(t *testing.T) {
	obj := initConf()
	tx, err := obj.StartTx()
	if err != nil {
		return
	}

	file := &model.ConfigFile{
		Id:            1,
		Name:          "Name1",
		Namespace:     "Namespace1",
		Group:         "Group1",
		OriginContent: "OriginContent1",
		Content:       "Content1",
		Comment:       "Comment1",
		Format:        "Format1",
		Flag:          1,
		Valid:         false,
		Metadata: map[string]string{
			"Metadata": "aaaa1",
		},
		Encrypt:     true,
		EncryptAlgo: "EncryptAlgo1",
		Status:      "Status1",
		CreateBy:    "2024-10-01 11:11:11",
		ModifyBy:    "2023-10-01 11:11:12",
		ReleaseBy:   "2024-10-10 11:11:13",
		CreateTime:  UnixSecondToTime(5),
		ModifyTime:  UnixSecondToTime(6),
		ReleaseTime: UnixSecondToTime(7),
	}

	err = obj.configFileStore.CreateConfigFileTx(tx, file)
	fmt.Printf("err: %+v\n", err)
}

func TestCountConfigFiles(t *testing.T) {
	obj := initConf()
	ret, err := obj.configFileStore.CountConfigFiles("Namespace1", "Group1")
	fmt.Printf("ret: %+v, err: %+v\n", ret, err)
}

func TestUpdateConfigFileTx(t *testing.T) {
	obj := initConf()
	tx, err := obj.StartTx()
	if err != nil {
		return
	}
	file := &model.ConfigFile{
		Id:        1,
		Name:      "Name1",
		Namespace: "Namespace1",
		Group:     "Group1",

		Content:    "Content2",
		Comment:    "Comment2",
		Format:     "Format2",
		ModifyBy:   "2023-10-02 11:11:12",
		ModifyTime: UnixSecondToTime(6),
	}
	err = obj.configFileStore.UpdateConfigFileTx(tx, file)
	fmt.Printf("err: %+v\n", err)
}

func TestCreateConfigFileGroup(t *testing.T) {
	obj := initConf()

	fileGroup := &model.ConfigFileGroup{
		Id:        1,
		Name:      "Name2",
		Namespace: "Namespace2",
		Comment:   "Comment2",
		Valid:     false,
		Metadata: map[string]string{
			"Metadata": "aaaa2",
		},
		CreateBy:   "2024-10-01 11:11:11",
		ModifyBy:   "2023-10-01 11:11:12",
		CreateTime: UnixSecondToTime(5),
		ModifyTime: UnixSecondToTime(6),
	}

	ret, err := obj.configFileGroupStore.CreateConfigFileGroup(fileGroup)
	fmt.Printf("ret: %+v, err: %+v\n", ret, err)
}

func TestUpdateConfigFileGroup(t *testing.T) {
	obj := initConf()

	fileGroup := &model.ConfigFileGroup{
		Id:        1,
		Name:      "Name2",
		Namespace: "Namespace2",
		Comment:   "Comment4",
		Valid:     false,
		Metadata: map[string]string{
			"Metadata": "aaaa3",
		},
		CreateBy:   "2024-10-01 11:11:11",
		ModifyBy:   "2023-10-01 11:11:12",
		CreateTime: UnixSecondToTime(5),
		ModifyTime: UnixSecondToTime(6),
	}

	err := obj.configFileGroupStore.UpdateConfigFileGroup(fileGroup)
	fmt.Printf("err: %+v\n", err)
}

func TestCreateConfigFileReleaseTx(t *testing.T) {
	obj := initConf()
	tx, err := obj.StartTx()
	if err != nil {
		return
	}

	simple := &model.SimpleConfigFileRelease{
		Version: 1,
		Comment: "Comment",
		Md5:     "Md5",
		Flag:    1,
		Active:  true,
		Valid:   true,
		Format:  "Format",
		Metadata: map[string]string{
			"aaa": "aaaa",
		},
		CreateTime:         UnixSecondToTime(5),
		CreateBy:           "2024-10-01 11:11:11",
		ModifyTime:         UnixSecondToTime(5),
		ModifyBy:           "2024-10-01 11:11:11",
		ReleaseDescription: "ReleaseDescription",
		BetaLabels:         nil,
	}
	key := &model.ConfigFileReleaseKey{
		Id:          1,
		Name:        "Name1",
		Namespace:   "Namespace1",
		Group:       "Group1",
		FileName:    "FileName1",
		ReleaseType: "ReleaseType1",
	}
	simple.ConfigFileReleaseKey = key
	clientLabels := make([]*apimodel.ClientLabel, 0)
	clientLabel := &apimodel.ClientLabel{
		Key: "key1",
		Value: &apimodel.MatchString{
			Type:      1,
			ValueType: 1,
		},
	}
	clientLabels = append(clientLabels, clientLabel)
	simple.BetaLabels = clientLabels

	file := &model.ConfigFileRelease{
		Content: "Content1",
	}
	file.SimpleConfigFileRelease = simple

	err = obj.configFileReleaseStore.CreateConfigFileReleaseTx(tx, file)
	fmt.Printf("1111, err: %+v\n", err)
}

func TestCreateConfigFileTemplate(t *testing.T) {
	obj := initConf()

	file := &model.ConfigFileTemplate{
		Id:         1,
		Name:       "Name2",
		Content:    "Content2",
		Comment:    "Comment2",
		Format:     "Format",
		CreateBy:   "2024-10-01 11:11:11",
		ModifyBy:   "2023-10-01 11:11:12",
		CreateTime: UnixSecondToTime(5),
		ModifyTime: UnixSecondToTime(6),
	}

	ret, err := obj.configFileTemplateStore.CreateConfigFileTemplate(file)
	fmt.Printf("ret: %+v, err: %+v\n", ret, err)
}
