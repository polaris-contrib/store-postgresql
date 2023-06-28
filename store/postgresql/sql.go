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
	"strconv"
	"strings"

	"github.com/polarismesh/polaris/common/utils"
)

const (
	// OwnerAttribute
	OwnerAttribute string = "owner"

	// And
	And = " and"
)

// Order 排序结构体
type Order struct {
	Field    string
	Sequence string
}

// Page 分页结构体
type Page struct {
	Offset uint32
	Limit  uint32
}

func boolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

func genNamespaceWhereSQLAndArgs(str string, filter map[string][]string, order *Order,
	offset, limit int) (string, []interface{}) {
	num := 0
	var sqlIndex = 1

	for _, value := range filter {
		num += len(value)
	}
	args := make([]interface{}, 0, num+2)

	if num > 0 {
		str += "where"
		firstIndex := true

		for index, value := range filter {
			if !firstIndex {
				str += And
			}
			str += " ("

			firstItem := true
			for _, item := range value {
				if !firstItem {
					str += " or "
				}
				if index == OwnerAttribute {
					str += fmt.Sprintf("owner like $%d", sqlIndex)
					item = "%" + item + "%"
				} else {
					if index == NameAttribute && utils.IsWildName(item) {
						str += fmt.Sprintf("name like $%d", sqlIndex)
						item = utils.ParseWildNameForSql(item)
					} else {
						str += index + fmt.Sprintf("=$%d", sqlIndex)
					}
				}
				args = append(args, item)
				firstItem = false

				sqlIndex++
			}
			firstIndex = false
			str += ")"
		}
	}

	if order != nil {
		str += " order by " + order.Field + " " + order.Sequence
	}

	str += fmt.Sprintf(" limit $%d offset $%d", sqlIndex, sqlIndex+1)
	args = append(args, limit, offset)

	return str, args
}

// PlaceholdersNI PlaceholdersN 构造多个占位符
func PlaceholdersNI(size, indexSort int) (string, int) {
	if size <= 0 {
		return "", indexSort
	}
	var strs []string
	for i := 1; i <= size; i++ {
		strs = append(strs, fmt.Sprintf("$%d", indexSort))
		indexSort++
	}
	return strings.Join(strs, ","), indexSort
}

// genServiceFilterSQL 根据service filter生成where相关的语句
func genServiceFilterSQL(filter map[string]string, indexSort int) (string, []interface{}, int) {
	if len(filter) == 0 {
		return "", nil, indexSort
	}

	args := make([]interface{}, 0, len(filter))
	var str string
	firstIndex := true
	for key, value := range filter {
		if !firstIndex {
			str += And
		}
		firstIndex = false

		if key == OwnerAttribute {
			str += fmt.Sprintf(" (service.name, service.namespace) in (select service,"+
				"namespace from owner_service_map where owner=$%d)", indexSort)
		} else if key == "alias."+OwnerAttribute {
			str += fmt.Sprintf(" (alias.name, alias.namespace) in (select service,namespace "+
				"from owner_service_map where owner=$%d)", indexSort)
		} else if key == "business" {
			str += fmt.Sprintf(" %s like $%d", key, indexSort)
			value = "%" + value + "%"
		} else if key == "name" && utils.IsPrefixWildName(value) {
			str += fmt.Sprintf(" name like $%d", indexSort)
			value = "%" + value[0:len(value)-1] + "%"
		} else {
			str += " " + key + fmt.Sprintf("=$%d", indexSort)
		}

		indexSort++

		args = append(args, value)
	}

	return str, args, indexSort
}

// genOrderAndPage 生成order和page相关语句
func genOrderAndPage(order *Order, page *Page, indexSort int) (string, []interface{}, int) {
	var str string
	var args []interface{}
	if order != nil {
		str += " order by " + order.Field + " " + order.Sequence
	}
	if page != nil {
		str += fmt.Sprintf(" limit $%d offset $%d", indexSort, indexSort+1)
		args = append(args, page.Limit, page.Offset)
	}

	return str, args, indexSort + 2
}

// genServiceAliasWhereSQLAndArgs 生成service alias查询数据的where语句和对应参数
func genServiceAliasWhereSQLAndArgs(str string, filter map[string]string, order *Order,
	offset uint32, limit uint32, indexSort int) (
	string, []interface{}) {
	baseStr := str
	filterStr, filterArgs, indexSort1 := genServiceFilterSQL(filter, indexSort)
	indexSort = indexSort1
	if filterStr != "" {
		baseStr += " where "
	}
	page := &Page{offset, limit}
	opStr, opArgs, _ := genOrderAndPage(order, page, indexSort)

	return baseStr + filterStr + opStr, append(filterArgs, opArgs...)
}

// genWhereSQLAndArgs 生成service和instance查询数据的where语句和对应参数
func genWhereSQLAndArgs(str string, filter, metaFilter map[string]string, order *Order,
	offset uint32, limit uint32) (string, []interface{}) {
	baseStr := str
	var (
		args  []interface{}
		index = 1
	)

	filterStr, filterArgs, index1 := genFilterSQL(filter, index)
	index = index1
	var conjunction = " where "
	if filterStr != "" {
		baseStr += " where " + filterStr
		conjunction = " and "
	}
	args = append(args, filterArgs...)
	var metaStr string
	var metaArgs []interface{}
	if len(metaFilter) > 0 {
		metaStr, metaArgs, index = genInstanceMetadataArgs(metaFilter, index)
		args = append(args, metaArgs...)
		baseStr += conjunction + metaStr
	}
	page := &Page{offset, limit}
	opStr, opArgs, index2 := genOrderAndPage(order, page, index)
	index = index2

	return baseStr + opStr, append(args, opArgs...)
}

// genFilterSQL 根据filter生成where相关的语句
func genFilterSQL(filter map[string]string, index int) (string, []interface{}, int) {
	if len(filter) == 0 {
		return "", nil, index
	}

	args := make([]interface{}, 0, len(filter))
	var str string
	firstIndex := true
	for key, value := range filter {
		if !firstIndex {
			str += And
		}
		firstIndex = false
		// 这个查询组装，先这样完成，后续优化filter TODO
		if key == OwnerAttribute || key == "alias."+OwnerAttribute || key == "business" {
			str += fmt.Sprintf(" %s like $%d", key, index)
			value = "%" + value + "%"
		} else if key == "name" && utils.IsWildName(value) {
			str += fmt.Sprintf(" name like $%d", index)
			value = utils.ParseWildNameForSql(value)
		} else if key == "id" {
			if utils.IsWildName(value) {
				str += fmt.Sprintf(" instance.id like $%d", index)
				value = utils.ParseWildNameForSql(value)
			} else {
				str += fmt.Sprintf(" instance.id = $%d", index)
			}
		} else if key == "host" {
			hosts := strings.Split(value, ",")
			placeholder, index1 := PlaceholdersNI(len(hosts), index)
			index = index1
			str += " host in (" + placeholder + ")"
			for _, host := range hosts {
				args = append(args, host)
			}
		} else if key == "managed" {
			str += fmt.Sprintf(" managed = $%d", index)
			managed, _ := strconv.ParseBool(value)
			args = append(args, boolToInt(managed))
			index++
			continue
		} else if key == "namespace" && utils.IsWildName(value) {
			str += fmt.Sprintf(" namespace like $%d", index)
			value = utils.ParseWildNameForSql(value)
		} else {
			str += " " + key + fmt.Sprintf("=$%d", index)
		}
		if key != "host" {
			args = append(args, value)
		}
		index++
	}

	return str, args, index
}

func genInstanceMetadataArgs(metaFilter map[string]string, index int) (string, []interface{}, int) {
	str := fmt.Sprintf(`instance.id in (select id from instance_metadata where $%d = $%d and $%d = $%d)`,
		index, index+1, index+2, index+3)
	args := make([]interface{}, 0, 2)
	for k, v := range metaFilter {
		args = append(args, k)
		args = append(args, v)
	}
	return str, args, index + 4
}

// genRuleFilterSQL 根据规则的filter生成where相关的语句
func genRuleFilterSQL(tableName string, filter map[string]string,
	index int) (string, []interface{}, int) {
	if len(filter) == 0 {
		return "", nil, index
	}

	args := make([]interface{}, 0, len(filter))
	var str string
	firstIndex := true
	for key, value := range filter {
		if tableName != "" {
			key = tableName + "." + key
		}
		if !firstIndex {
			str += And
		}
		if key == OwnerAttribute || key == (tableName+"."+OwnerAttribute) {
			str += fmt.Sprintf(" %s like $%d ", key, index)
			value = "%" + value + "%"
		} else {
			str += " " + key + fmt.Sprintf(" = $%d ", index)
		}
		args = append(args, value)
		firstIndex = false
	}
	return str, args, index
}
