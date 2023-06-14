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

	"github.com/polarismesh/polaris/common/model"
	apimodel "github.com/polarismesh/specification/source/go/api/v1/model"
	apiservice "github.com/polarismesh/specification/source/go/api/v1/service_manage"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func createMockClients(total int) []*model.Client {
	ret := make([]*model.Client, 0, total)

	for i := 0; i < total; i++ {
		ret = append(ret, model.NewClient(&apiservice.Client{
			Host:    &wrapperspb.StringValue{Value: fmt.Sprintf("client-%d", i)},
			Type:    0,
			Version: &wrapperspb.StringValue{Value: fmt.Sprintf("client-%d", i)},
			Location: &apimodel.Location{
				Region: &wrapperspb.StringValue{Value: fmt.Sprintf("client-%d-region", i)},
				Zone:   &wrapperspb.StringValue{Value: fmt.Sprintf("client-%d-zone", i)},
				Campus: &wrapperspb.StringValue{Value: fmt.Sprintf("client-%d-campus", i)},
			},
			Id: &wrapperspb.StringValue{Value: fmt.Sprintf("client-%d", i)},
			Stat: []*apiservice.StatInfo{
				{
					Target:   &wrapperspb.StringValue{Value: "prometheus"},
					Port:     &wrapperspb.UInt32Value{Value: 8080},
					Path:     &wrapperspb.StringValue{Value: "/metrics"},
					Protocol: &wrapperspb.StringValue{Value: "http"},
				},
			},
		}))
	}

	return ret
}

func TestBatchAddClients(t *testing.T) {
	obj := initConf()
	requests := createMockClients(5)
	err := obj.clientStore.BatchAddClients(requests)
	fmt.Printf("err: %+v", err)
}
