# GoTransactions

Transactions in golang

## API

We have the following functions:

- `OnTransaction`
- `OnRollback`

Where we define the transactions and its rollback in case of failure and go back to the state we desire.

There are some examples in the `examples` folder.

## How to use it

Let's write a simple example.

```go
func main() {
	onTransaction := gotransactions.OnTransaction(func() error {
		return nil
    })
	
	onRollback := gotransactions.OnRollback(func() error {
		return nil
    })
	
	transaction := gotransactions.New(onTransaction, onRollback)
	if err := transaction.ExecuteTransaction(); err != nil {
		fmt.Println(err)
    }
}
```

## TODO

- Context support