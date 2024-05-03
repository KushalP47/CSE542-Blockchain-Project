package database

import (
	// "github.com/KushalP47/CSE542-Blockchain-Project/blockchain"

	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger"
)

// key: account address
// value: account

func ReadAccount(address [20]byte) ([]byte, error) {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	utils.HandleError(err)
	defer db.Close()

	var account []byte
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(address[:])
		utils.HandleError(err)
		err = item.Value(func(val []byte) error {
			account, err = val, nil
			return nil
		})
		return err
	})
	utils.HandleError(err)
	return account, err
}

func WriteAccount(key [20]byte, value []byte) error {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	utils.HandleError(err)
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key[:], value)
		utils.HandleError(err)
		return err
	})
	utils.HandleError(err)
	return err
}
