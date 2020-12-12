package main

import (
	"fmt"
	gotr "github.com/ervitis/gotransactions"
)

func main() {
	count := 1

	onTransaction := gotr.OnTransaction(func() error {
		count++
		return nil
	})

	onRollback := gotr.OnRollback(func() error {
		count--
		return nil
	})

	transaction := gotr.New(onTransaction, onRollback)

	if err := transaction.ExecuteTransaction(); err != nil {
		panic(err)
	}

	fmt.Println(count)
}
