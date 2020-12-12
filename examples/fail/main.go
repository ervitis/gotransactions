package main

import (
	"errors"
	"fmt"
	gotr "github.com/ervitis/gotransactions"
)

func main() {
	count := 1

	onTransaction := gotr.OnTransaction(func() error {
		return errors.New("ups")
	})

	onRollback := gotr.OnRollback(func() error {
		if count > 1 {
			count--
		}
		return nil
	})

	transaction := gotr.New(onTransaction, onRollback)

	if err := transaction.ExecuteTransaction(); err != nil {
		panic(err)
	}

	fmt.Println(count)
}
