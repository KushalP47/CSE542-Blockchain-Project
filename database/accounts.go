package database

import (
	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger"
)

// key: account address
// value: account

// type Account struct {
// 	Address []byte
// 	Name    string
// 	Nonce   uint64
// 	Balance uint64
// }

func ReadAccount(address []byte) blockchain.Account {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	utils.HandleError(err)
	defer db.Close()

	var account blockchain.Account
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(address)
		utils.HandleError(err)
		err = item.Value(func(val []byte) error {
			account, err = blockchain.DeserializeAccount(val)
			return nil
		})
		return err
	})
	utils.HandleError(err)
	return account
}

func WriteAccount(account blockchain.Account) error {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	utils.HandleError(err)
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		serialized, _ := blockchain.SerializeAccount(account)
		err := txn.Set([]byte(account.Address), serialized)
		utils.HandleError(err)
		return err
	})
	utils.HandleError(err)
	return err
}

// getLatestNonce returns the nonce of the last account of database
func GetLatestNonce() uint64 {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	utils.HandleError(err)
	defer db.Close()

	var latestNonce uint64
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				account, _ := blockchain.DeserializeAccount(val)
				if account.Nonce > latestNonce {
					latestNonce = account.Nonce
				}
				return nil
			})
			utils.HandleError(err)
		}
		return nil
	})
	utils.HandleError(err)
	return latestNonce
}
