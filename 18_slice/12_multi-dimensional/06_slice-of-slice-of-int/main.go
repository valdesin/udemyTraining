package main

import "fmt"

func main() {
	transactions := make([][]int, 0)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 4)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}

	fmt.Println(transactions)
}
