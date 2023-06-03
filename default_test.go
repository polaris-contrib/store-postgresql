package postgresql

import (
	"fmt"
	"github.com/polarismesh/polaris/common/model"
	"github.com/polarismesh/polaris/store"
	"testing"
)

func initConf() *postgresqlStore {
	conf := &store.Config{
		Name: "Postgresql",
		Option: map[string]interface{}{
			"master": map[interface{}]interface{}{
				"dbType": "postgres",
				"dbUser": "postgres",
				"dbPwd":  "aaaaaa",
				"dbAddr": "192.168.31.19",
				"dbPort": "5432",
				"dbName": "polaris_server",

				"maxOpenConns":     10,
				"maxIdleConns":     10,
				"connMaxLifetime":  10,
				"txIsolationLevel": 2,
			},
			"slave": map[interface{}]interface{}{
				"dbType": "postgres",
				"dbUser": "postgres",
				"dbPwd":  "aaaaaa",
				"dbAddr": "192.168.31.19",
				"dbPort": "5432",
				"dbName": "polaris_server",

				"maxOpenConns":     10,
				"maxIdleConns":     10,
				"connMaxLifetime":  10,
				"txIsolationLevel": 2,
			},
		},
	}
	obj := &postgresqlStore{}
	err := obj.Initialize(conf)
	fmt.Println(err)

	return obj
}

func TestCreateTransaction(t *testing.T) {
	obj := initConf()

	tran, err := obj.CreateTransaction()

	fmt.Println("tran: ", tran, err)
}

func TestAddNamespace(t *testing.T) {
	obj := initConf()

	modelNamespace := &model.Namespace{
		Name:    "Test",
		Comment: "Polaris-test",
		Token:   "2d1bfe5d12e04d54b8ee69e62494c7fe",
		Owner:   "polaris",
		Valid:   false,
		//CreateTime time.Time
		//ModifyTime time.Time
	}
	err := obj.namespaceStore.AddNamespace(modelNamespace)

	fmt.Printf("namespace: %+v\n", err)
}
