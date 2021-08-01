package main

import (
	"KeyValueStore/constants"
	"KeyValueStore/transaction"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting key value store...")
	startConsoleApplication()
}

func startConsoleApplication() {
	reader := bufio.NewReader(os.Stdin)
	items := &transaction.TransactionStack{}

	for {
		fmt.Printf("> ")
		text, _ := reader.ReadString('\n')
		//	split the text into operation strings.
		operation := strings.Fields(text)
		switch operation[0] {
		case constants.BEGIN:
			items.PushTransaction()
			break
		case constants.ROLLBACK:
			items.RollBackTransaction()
		case constants.COMMIT:
			items.Commit()
			items.PopTransaction()
		case constants.END:
			items.PopTransaction()
		case constants.SET:
			transaction.Set(operation[1], operation[2], items)
		case constants.GET:
			transaction.Get(operation[1], items)
		case constants.EXIT:
			os.Exit(0)
		default:
			fmt.Printf("ERROR: Unrecognised Operation %s\n", operation[0])

		}
	}
}
