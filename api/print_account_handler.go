package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/dgraph-io/badger"
)

func PrintAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Open the accounts database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	if err != nil {
		fmt.Println("Error opening accounts database:", err)
		return
	}
	defer db.Close()

	// Print all accounts
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Prefix = []byte{} // Iterate over all keys
		iter := txn.NewIterator(opts)
		defer iter.Close()

		for iter.Seek([]byte("")); iter.Valid(); iter.Next() {
			item := iter.Item()

			// Get the value (account) from the item
			var account database.Account
			err := item.Value(func(val []byte) error {
				return json.Unmarshal(val, &account)
			})
			if err != nil {
				fmt.Println("Error decoding account:", err)
				continue
			}

			// Print the account
			fmt.Printf("Address: %s\n", account.Address)
			fmt.Println("Name:", account.Name)
			fmt.Println("Nonce:", account.Nonce)
			fmt.Println("Balance:", account.Balance)
			fmt.Println()
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error reading accounts database:", err)
	}
}
