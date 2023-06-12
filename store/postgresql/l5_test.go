package postgresql

import (
	"fmt"
	"testing"
)

func TestGenNextL5Sid(t *testing.T) {
	obj := initConf()
	var id uint32 = 1

	resp, err := obj.l5Store.GenNextL5Sid(id)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}

func TestGetMoreL5Routes(t *testing.T) {
	obj := initConf()
	var id uint32 = 1

	resp, err := obj.l5Store.GetMoreL5Routes(id)
	fmt.Printf("resp: %+v, err: %+v\n", resp, err)
}
