package main

import (
	"fmt"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
)

func main() {
	accountState := blockchain.InitAccounts()
	accountState.AddAccount("Alice", 100)
	accountState.AddAccount("Bob", 50)
	accountState.AddAccount("Charlie", 75)
	accountState.AddAccount("David", 25)
	fmt.Println(accountState)

	for _, account := range accountState.Accounts {
		fmt.Printf("Account: %s\n", account.Name)
		fmt.Printf("Address: %x\n", account.Address)
		fmt.Printf("Nonce: %d\n", account.Nonce)
		fmt.Printf("Balance: %d\n", account.Balance)
		fmt.Println()
	}
}
