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

		onTr OnTransaction
		onRb OnRollback
		Ctx  context.Context
	}

	Option func()
)

func New(onTransaction OnTransaction, onRollback OnRollback) TransactionIface {
	return &transaction{
		onRb: onRollback,
		onTr: onTransaction,
	}
}

func (tr *transaction) ExecuteTransaction() error {
	tr.mtx.Lock()
	defer tr.mtx.Unlock()

	if err := tr.onTr(); err != nil {
		return tr.onRb()
	}
	return nil
}
