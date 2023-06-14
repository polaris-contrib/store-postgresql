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
	"bytes"
	"database/sql"
	"strings"
	"time"
	"unicode"

	"github.com/polarismesh/polaris/common/log"
	"github.com/polarismesh/polaris/store"
)

func toUnderscoreName(name string) string {
	var buf bytes.Buffer
	for i, token := range name {
		if unicode.IsUpper(token) && i > 0 {
			buf.WriteString("_")
		}
		buf.WriteString(strings.ToLower(string(token)))
	}
	return buf.String()
}

// QueryHandler is the interface that wraps the basic Query method.
type QueryHandler func(query string, args ...interface{}) (*sql.Rows, error)

// BatchHandler 批量查询数据的回调函数
type BatchHandler func(objects []interface{}) error

// BatchQuery 批量查询数据的对外接口
// 每次最多查询200个
func BatchQuery(label string, data []interface{}, handler BatchHandler) error {
	// start := time.Now()
	maxCount := 200
	beg := 0
	remain := len(data)
	if remain == 0 {
		return nil
	}

	progress := 0
	for {
		if remain > maxCount {
			if err := handler(data[beg : beg+maxCount]); err != nil {
				return err
			}

			beg += maxCount
			remain -= maxCount
			progress += maxCount
			if progress%20000 == 0 {
				log.Infof("[Store][database][Batch] query (%s) progress(%d / %d)", label, progress, len(data))
			}
		} else {
			if err := handler(data[beg : beg+remain]); err != nil {
				return err
			}
			break
		}
	}

	return nil
}

// queryEntryCount 单独查询count个数的执行函数
func queryEntryCount(conn *BaseDB, str string, args []interface{}) (uint32, error) {
	var count uint32
	var err error
	Retry("queryRow", func() error {
		err = conn.QueryRow(str, args...).Scan(&count)
		return err
	})
	switch {
	case err == sql.ErrNoRows:
		log.Errorf("[Store][database] not found any entry(%s)", str)
		return 0, err
	case err != nil:
		log.Errorf("[Store][database] query entry count(%s) err: %s", str, err.Error())
		return 0, err
	default:
		return count, nil
	}
}

// aliasFilter2Where 别名查询转换
var aliasFilter2Where = map[string]string{
	"service":         "source.name",
	"namespace":       "source.namespace",
	"alias":           "alias.name",
	"alias_namespace": "alias.namespace",
	"owner":           "alias.owner",
}

// serviceAliasFilter2Where 别名查询字段转换函数
func serviceAliasFilter2Where(filter map[string]string) map[string]string {
	out := make(map[string]string)
	for k, v := range filter {
		if d, ok := aliasFilter2Where[k]; ok {
			out[d] = v
		} else {
			out[k] = v
		}
	}

	return out
}

// BatchOperation 批量操作
// @note 每次最多操作100个
func BatchOperation(label string, data []interface{}, handler BatchHandler) error {
	if data == nil {
		return nil
	}
	maxCount := 100
	progress := 0
	for begin := 0; begin < len(data); begin += maxCount {
		end := begin + maxCount
		if end > len(data) {
			end = len(data)
		}
		if err := handler(data[begin:end]); err != nil {
			return err
		}
		progress += end - begin
		if progress%maxCount == 0 {
			log.Infof("[Store][database][Batch] operation (%s) progress(%d/%d)", label, progress, len(data))
		}
	}
	return nil
}

// checkDataBaseAffectedRows 检查数据库处理返回的行数
func checkDataBaseAffectedRows(result sql.Result, counts ...int64) error {
	n, err := result.RowsAffected()
	if err != nil {
		log.Errorf("[Store][Database] get rows affected err: %s", err.Error())
		return err
	}

	for _, c := range counts {
		if n == c {
			return nil
		}
	}

	log.Errorf("[Store][Database] get rows affected result(%d) is not match expect(%+v)", n, counts)
	return store.NewStatusError(store.AffectedRowsNotMatch, "affected rows not matched")
}

func GetLocation() *time.Location {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loc = time.FixedZone("CTS", 8*3600)
	}
	return loc
}

// GetCurrentTimeFormat 获取格式化时间
// @param millTime：时间毫秒
// @return 2006-01-02 15:04:05
func GetCurrentTimeFormat() string {
	loc := GetLocation()
	currentTime := time.Now().In(loc)
	format := currentTime.Format("2006-01-02 15:04:05")
	return format
}

func CurrDiffTimeSecond(beforeTime time.Time) float64 {
	loc := GetLocation()
	currentTime := time.Now().In(loc)
	diff := beforeTime.Sub(currentTime)
	return diff.Seconds()
}

func GetCurrentSsTimestamp() int64 {
	loc := GetLocation()
	now := time.Now().In(loc)
	curTime := now.Unix()
	return curTime
}

// UnixSecondToTime 秒级时间戳转time
func UnixSecondToTime(second int64) time.Time {
	loc := GetLocation()
	return time.Unix(second, 0).In(loc)
}

// UnixMilliToTime 毫秒级时间戳转time
func UnixMilliToTime(milli int64) time.Time {
	loc := GetLocation()
	return time.Unix(milli/1000, (milli%1000)*(1000*1000)).In(loc)
}

// UnixNanoToTime 纳秒级时间戳转time
func UnixNanoToTime(nano int64) time.Time {
	loc := GetLocation()
	return time.Unix(nano/(1000*1000*1000), nano%(1000*1000*1000)).In(loc)
}

// ConvertSecond 转换指定时间的毫秒
func ConvertSecond(ctime time.Time) int64 {
	return ctime.In(GetLocation()).Unix()
}
