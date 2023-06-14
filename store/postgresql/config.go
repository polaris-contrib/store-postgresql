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
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Pgsql struct {
	// Pgsql账号
	Username string `json:"username"`
	// Pgsql密码
	Password string `json:"password"`
	// Pgsql地址
	Address string `json:"address"`
	// Pgsql端口
	Port int `json:"port"`
	// 数据库名称
	Database string `json:"database"`
	// 表名称
	Table string `json:"table"`
	// 账号字段名称
	AccountField string `json:"accountField"`
	// 密码字段名称
	PwdField string `json:"pwdField"`
}

// 配置文件路径
var ConfigPath string = "config/config.yaml"

// postgresql配置信息缓存
var PgsqlData *Pgsql

func Init() {
	//初始化配置对象
	PgsqlData = new(Pgsql)
	//读取配置文件
	file, err := os.Open(ConfigPath)
	if err != nil {
		fmt.Println("config path:", err)
		os.Exit(1)
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("config file:", err)
		os.Exit(1)
	}
	//使用json转换至config对象中
	err = yaml.Unmarshal(bytes, PgsqlData)
	if err != nil {
		fmt.Println("json unmarshal:", err)
		os.Exit(1)
	}
}
