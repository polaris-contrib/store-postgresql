package postgresql

import (
	"fmt"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
	apiservice "github.com/polarismesh/specification/source/go/api/v1/service_manage"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
	"time"
)

func TestAddInstance(t *testing.T) {
	obj := initConf()
	modelService := &model.Instance{
		Proto: &apiservice.Instance{
			Id: &wrapperspb.StringValue{
				Value: "1111t",
			},
			Service: &wrapperspb.StringValue{
				Value: "2222t",
			},
			Namespace: &wrapperspb.StringValue{
				Value: "3333t",
			},
			VpcId: &wrapperspb.StringValue{
				Value: "4444t",
			},
			Host: &wrapperspb.StringValue{Value: "5555t"},
			Port: &wrapperspb.UInt32Value{
				Value: 1001,
			},
			Healthy: &wrapperspb.BoolValue{
				Value: true,
			},
			Isolate: &wrapperspb.BoolValue{
				Value: false,
			},
			HealthCheck: &apiservice.HealthCheck{
				Type: apiservice.HealthCheck_HEARTBEAT,
				Heartbeat: &apiservice.HeartbeatHealthCheck{
					Ttl: &wrapperspb.UInt32Value{
						Value: 1,
					},
				},
			},
			Metadata: map[string]string{
				"mkey":   "6666t",
				"mvalue": "7777t",
			},
		},
		ServiceID:         "111t",
		ServicePlatformID: "222t",
		// Valid Whether it is deleted by logic
		Valid: false,
		// ModifyTime Update time of instance
		ModifyTime: time.Now(),
	}
	err := obj.instanceStore.AddInstance(modelService)
	fmt.Println("err: ", err)
}

func TestBatchAddInstances(t *testing.T) {
	obj := initConf()
	modelService := &model.Instance{
		Proto: &apiservice.Instance{
			Id: &wrapperspb.StringValue{
				Value: "1111i",
			},
			Service: &wrapperspb.StringValue{
				Value: "2222i",
			},
			Namespace: &wrapperspb.StringValue{
				Value: "3333i",
			},
			VpcId: &wrapperspb.StringValue{
				Value: "4444i",
			},
			Host: &wrapperspb.StringValue{Value: "5555i"},
			Port: &wrapperspb.UInt32Value{
				Value: 1001,
			},
			Healthy: &wrapperspb.BoolValue{
				Value: true,
			},
			Isolate: &wrapperspb.BoolValue{
				Value: false,
			},
			HealthCheck: &apiservice.HealthCheck{
				Type: apiservice.HealthCheck_HEARTBEAT,
				Heartbeat: &apiservice.HeartbeatHealthCheck{
					Ttl: &wrapperspb.UInt32Value{
						Value: 1,
					},
				},
			},
			Metadata: map[string]string{
				"mkey":   "6666i",
				"mvalue": "7777i",
			},
		},
		ServiceID:         "111i",
		ServicePlatformID: "222i",
		// Valid Whether it is deleted by logic
		Valid: false,
		// ModifyTime Update time of instance
		ModifyTime: time.Now(),
	}
	ms := make([]*model.Instance, 0)
	ms = append(ms, modelService)
	err := obj.instanceStore.BatchAddInstances(ms)
	fmt.Println("err: ", err)
}

func TestUpdateInstance(t *testing.T) {
	obj := initConf()
	modelService := &model.Instance{
		Proto: &apiservice.Instance{
			Id: &wrapperspb.StringValue{
				Value: "1111t",
			},
			Service: &wrapperspb.StringValue{
				Value: "2222t",
			},
			Namespace: &wrapperspb.StringValue{
				Value: "3333t",
			},
			VpcId: &wrapperspb.StringValue{
				Value: "4444t",
			},
			Host: &wrapperspb.StringValue{Value: "5555t"},
			Port: &wrapperspb.UInt32Value{
				Value: 1001,
			},
			Healthy: &wrapperspb.BoolValue{
				Value: true,
			},
			Isolate: &wrapperspb.BoolValue{
				Value: false,
			},
			/*HealthCheck: &apiservice.HealthCheck{
				Type: apiservice.HealthCheck_HEARTBEAT,
				Heartbeat: &apiservice.HeartbeatHealthCheck{
					Ttl: &wrapperspb.UInt32Value{
						Value: 1,
					},
				},
			},*/
			Metadata: map[string]string{
				"mkey":   "6666t",
				"mvalue": "7777t",
			},
		},
		ServiceID:         "111t",
		ServicePlatformID: "222t",
		// Valid Whether it is deleted by logic
		Valid: false,
		// ModifyTime Update time of instance
		ModifyTime: time.Now(),
	}
	err := obj.instanceStore.UpdateInstance(modelService)
	fmt.Println("err: ", err)
}

func TestDeleteInstance(t *testing.T) {
	obj := initConf()
	err := obj.instanceStore.DeleteInstance("1111t")
	fmt.Println("err: ", err)
}

func TestGetInstance(t *testing.T) {
	obj := initConf()
	resp, err := obj.instanceStore.GetInstance("1111t")
	fmt.Printf("resp: %+v, err: %+v", resp, err)
}

func TestBatchGetInstanceIsolate(t *testing.T) {
	obj := initConf()
	ids := map[string]bool{
		"1111":  true,
		"1111a": true,
		"1111b": true,
		"1111c": true,
	}
	resp, err := obj.instanceStore.BatchGetInstanceIsolate(ids)
	fmt.Printf("resp: %+v, err: %+v", resp, err)
}

func TestGetInstancesBrief(t *testing.T) {
	obj := initConf()
	ids := map[string]bool{
		"1111":  true,
		"1111a": true,
		"1111b": true,
		"1111c": true,
	}
	resp, err := obj.instanceStore.GetInstancesBrief(ids)
	fmt.Printf("resp: %+v, err: %+v", resp, err)
}

func TestGetInstancesCount(t *testing.T) {
	obj := initConf()
	resp, err := obj.instanceStore.GetInstancesCount()
	fmt.Printf("resp: %+v, err: %+v", resp, err)
}

func TestGetInstancesMainByService(t *testing.T) {
	obj := initConf()
	resp, err := obj.instanceStore.GetInstancesMainByService("111a", "55555b")
	fmt.Printf("resp: %+v, err: %+v", resp, err)
}

func TestGetExpandInstances(t *testing.T) {
	obj := initConf()
	filter := map[string]string{
		"service_id": "111a",
		"id":         "1111b",
	}
	metaFilter := map[string]string{
		"mkey":   "mvalue",
		"mvalue": "7777t",
	}
	cnt, resp, err := obj.instanceStore.GetExpandInstances(filter, metaFilter, 0, 10)
	fmt.Printf("cnt: %+v, resp: %+v, err: %+v", cnt, resp, err)
}

func TestGetMoreInstances(t *testing.T) {
	obj := initConf()
	resp, err := obj.instanceStore.GetMoreInstances(UnixSecondToTime(1685551812), false, true, []string{"111", "111a"})
	fmt.Printf("resp: %+v, err: %+v", resp, err)
}

func TestSetInstanceHealthStatus(t *testing.T) {
	obj := initConf()
	err := obj.instanceStore.SetInstanceHealthStatus("1111", 1, "reversion")
	fmt.Printf("err: %+v", err)
}

func TestBatchSetInstanceHealthStatus(t *testing.T) {
	obj := initConf()
	ids := []interface{}{"1111a", "1111b"}
	err := obj.instanceStore.BatchSetInstanceHealthStatus(ids, 1, "reversion")
	fmt.Printf("err: %+v", err)
}

func TestBatchSetInstanceIsolate(t *testing.T) {
	obj := initConf()
	ids := []interface{}{"1111a", "1111b"}
	err := obj.instanceStore.BatchSetInstanceIsolate(ids, 1, "reversion")
	fmt.Printf("err: %+v", err)
}

func TestBatchAppendInstanceMetadata(t *testing.T) {
	obj := initConf()
	request := &store.InstanceMetadataRequest{
		InstanceID: "111",
		Revision:   "revision1",
		Keys:       []string{"111", "222"},
		Metadata: map[string]string{
			"aaa": "1111",
			"bbb": "2222",
		},
	}
	requests := make([]*store.InstanceMetadataRequest, 0)
	requests = append(requests, request)
	err := obj.instanceStore.BatchAppendInstanceMetadata(requests)
	fmt.Printf("err: %+v", err)
}

func TestBatchRemoveInstanceMetadata(t *testing.T) {
	obj := initConf()
	request := &store.InstanceMetadataRequest{
		InstanceID: "111",
		Revision:   "revision1",
		Keys:       []string{"111", "222"},
		Metadata: map[string]string{
			"aaa": "1111",
			"bbb": "2222",
		},
	}
	requests := make([]*store.InstanceMetadataRequest, 0)
	requests = append(requests, request)
	err := obj.instanceStore.BatchRemoveInstanceMetadata(requests)
	fmt.Printf("err: %+v", err)
}
