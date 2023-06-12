package postgresql

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateLeaderElection(t *testing.T) {
	obj := initConf()

	for i := 0; i < 2; i++ {
		//go func() {
		key := fmt.Sprintf("test%d", i)
		err := obj.adminStore.StartLeaderElection(key)
		fmt.Printf("err: %+v\n", err)
		//}()
	}
}

func TestCheckMtimeExpired(t *testing.T) {
	obj := initConf()

	key := fmt.Sprintf("test%d", 1)
	err := obj.adminStore.StartLeaderElection(key)
	fmt.Printf("err: %+v\n", err)

	select {}
}

func TestBatchCleanDeletedInstances(t *testing.T) {
	obj := initConf()

	resp, err := obj.adminStore.BatchCleanDeletedInstances(10*time.Minute, 5)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)

	select {}
}

func TestGetUnHealthyInstances(t *testing.T) {
	obj := initConf()

	resp, err := obj.adminStore.GetUnHealthyInstances(10*time.Minute, 5)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)

	select {}
}

func TestBatchCleanDeletedClients(t *testing.T) {
	obj := initConf()

	resp, err := obj.adminStore.BatchCleanDeletedClients(10*time.Minute, 5)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}
