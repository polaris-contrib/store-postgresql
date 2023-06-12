package postgresql

import (
	"fmt"
	"github.com/polarismesh/polaris/common/model"
	apimodel "github.com/polarismesh/specification/source/go/api/v1/model"
	apiservice "github.com/polarismesh/specification/source/go/api/v1/service_manage"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
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
