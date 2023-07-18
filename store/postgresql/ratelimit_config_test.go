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

func TestCreateRateLimit(t *testing.T) {
	obj := initConf()

	// Method: Labels: Priority:0 Rule:{"service":{"value":"polaris.limiter"},"namespace":{"value":"Polaris"},"type":1,"amounts":[{"maxAmount":{"value":1},"validDuration":{"seconds":1}}],"action":{"value":"REJECT"},"disable":{},"regex_combine":{"value":true},"method":{"value":{}},"arguments":[{"key":"aa","value":{"value":{"value":"2"}}}],"name":{"value":"aa"},"max_queue_delay":{"value":1}} Revision:52df77fff7c14ec9a5a8306d381091fc Disable:false Valid:false CreateTime:0001-01-01 00:00:00 +0000 UTC ModifyTime:0001-01-01 00:00:00 +0000 UTC EnableTime:0001-01-01 00:00:00 +0000 UTC}),
	conf := &model.RateLimit{
		ID:        "d7af189c1986413ab928d501db200600",
		Name:      "aa",
		Disable:   true,
		ServiceID: "",
		Method:    "",
		Labels:    "",
		Priority:  1,
		Revision:  "4444",
	}
	err := obj.rateLimitStore.CreateRateLimit(conf)
	fmt.Printf("err: %+v\n", err)
}
