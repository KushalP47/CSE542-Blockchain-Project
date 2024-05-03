package database

import (
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
)

// key: Hash of the transaction
// value: transaction

func WriteTxn(key []byte, value []byte) error {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions"))
	utils.HandleError(err)
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		utils.HandleError(err)
		return err
	})
	utils.HandleError(err)
	return err
}

func ReadTxn(key []byte) ([]byte, error) {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions"))
	utils.HandleError(err)
	defer db.Close()

	var Txn []byte
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		utils.HandleError(err)
		err = item.Value(func(val []byte) error {
			Txn, err = val, nil
			return nil
		})
		return err
	})
	if err == badger.ErrKeyNotFound {
		return nil, nil
	} else if err != nil {
		utils.HandleError(err)
	}
	return Txn, err
}

func GetTxns() (map[common.Hash][]byte, error) {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions"))
	utils.HandleError(err)
	defer db.Close()

	txns := make(map[common.Hash][]byte)
	err = db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		count := 0
		for it.Rewind(); it.Valid(); it.Next() {
			if count == 5 {
				break
			}
			item := it.Item()
			var txnData []byte
			err := item.Value(func(val []byte) error {
				txnData = val
				return nil
			})
			if err != nil {
				return err
			}
			txns[common.Hash(item.Key())] = txnData

			// Delete the transaction
			err = txn.Delete(item.Key())
			if err != nil {
				return err
			}

			count++
		}
		return nil
	})
	if err != nil {
		utils.HandleError(err)
	}
	return txns, err
}
