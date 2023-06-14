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
	"errors"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRetry(t *testing.T) {
	Convey("重试可以成功", t, func() {
		var err error
		Retry("retry", func() error {
			err = errors.New("retry error")
			return err
		})
		So(err, ShouldNotBeNil)

		start := time.Now()
		count := 0
		Retry("retry", func() error {
			count++
			if count <= 10 {
				err = errors.New("invalid connection")
				return err
			}
			err = nil
			return nil
		})
		sub := time.Since(start)
		So(err, ShouldBeNil)
		So(sub, ShouldBeGreaterThan, time.Millisecond)
	})
	Convey("只捕获固定的错误", t, func() {
		for _, msg := range errMsg {
			var err error
			start := time.Now()
			Retry(fmt.Sprintf("retry-%s", msg), func() error {
				err = fmt.Errorf("my-error: %s", msg)
				return err
			})
			So(err, ShouldNotBeNil)
			So(time.Since(start), ShouldBeGreaterThan, time.Millisecond)
		}
	})
}
