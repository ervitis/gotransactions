package gotransactions

import (
	"context"
	"sync"
)

type (
	OnTransaction func() error

	OnRollback func() error

	TransactionIface interface {
		ExecuteTransaction() error
	}

	transaction struct {
		mtx sync.Mutex

		onTx OnTransaction
		onRb OnRollback
		Ctx  context.Context
	}
)

func New(onTransaction OnTransaction, onRollback OnRollback) TransactionIface {
	return &transaction{
		onRb: onRollback,
		onTx: onTransaction,
	}
}

func (tr *transaction) ExecuteTransaction() error {
	tr.mtx.Lock()
	defer tr.mtx.Unlock()

	if err := tr.onTx(); err != nil {
		return tr.onRb()
	}
	return nil
}
