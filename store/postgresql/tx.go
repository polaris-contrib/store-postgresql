package postgresql

import "github.com/polarismesh/polaris/store"

type Tx struct {
	delegateTx *BaseTx
}

func NewSqlDBTx(delegateTx *BaseTx) store.Tx {
	return &Tx{
		delegateTx: delegateTx,
	}
}

func (t *Tx) Commit() error {
	return t.delegateTx.Commit()
}

func (t *Tx) Rollback() error {
	return t.delegateTx.Rollback()
}

func (t *Tx) GetDelegateTx() interface{} {
	return t.delegateTx
}
