package transaction

import (
	"fmt"
)

var GlobalStore = make(map[string]string)

type Transaction struct {
	store map[string]string // local context of the transaction
	next  *Transaction
}

type TransactionStack struct {
	top  *Transaction
	size int
}

/* Creates and adds a new trxn to stack */
func (ts *TransactionStack) PushTransaction() {
	temp := Transaction{store: make(map[string]string)}
	temp.next = ts.top
	ts.top = &temp
	ts.size++
}

func (ts *TransactionStack) PopTransaction() {
	if ts.top == nil {
		fmt.Println("ERROR: no active transactions")
	} else {
		node := &Transaction{}
		ts.top = ts.top.next
		node.next = nil
		ts.size--
	}
}

func (ts *TransactionStack) Peek() *Transaction {
	return ts.top
}

/**
the global context variables are updated with the values of the
current trxn local context values.
*/
func (ts *TransactionStack) Commit() {
	currentTrxn := ts.Peek()
	if currentTrxn != nil {
		for key, value := range currentTrxn.store {
			GlobalStore[key] = value
			if currentTrxn.next != nil {
				//  update parent transaction
				currentTrxn.next.store[key] = value
			}
		}
	} else {
		fmt.Println("INFO: nothig to commit")
	}
}

func (ts *TransactionStack) RollBackTransaction() {
	if ts.top == nil {
		fmt.Println("ERROR: No active transactions")
	} else {
		for key := range ts.top.store {
			delete(ts.top.store, key)
		}
	}
}

//  Getter and setter
func Set(key string, value string, trxnStack *TransactionStack) {
	//  key:value store from active transaction
	currentTrxn := trxnStack.Peek()
	if currentTrxn == nil {
		GlobalStore[key] = value
	} else {
		currentTrxn.store[key] = value
	}
}

/*Get value of key from Store*/
func Get(key string, T *TransactionStack) {
	ActiveTransaction := T.Peek()
	if ActiveTransaction == nil {
		if val, ok := GlobalStore[key]; ok {
			fmt.Printf("%s\n", val)
		} else {
			fmt.Printf("%s not set\n", key)
		}
	} else {
		if val, ok := ActiveTransaction.store[key]; ok {
			fmt.Printf("%s\n", val)
		} else {
			fmt.Printf("%s not set\n", key)
		}
	}
}
