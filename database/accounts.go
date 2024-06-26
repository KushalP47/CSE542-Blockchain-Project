package database

import (
	// "github.com/KushalP47/CSE542-Blockchain-Project/blockchain"

	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
)

// key: account address
// value: account

func ReadAccount(address common.Address) ([]byte, error) {
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
	if err == badger.ErrKeyNotFound {
		return nil, nil
	}
	return account, err
}

func WriteAccount(key common.Address, value []byte) error {
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

func AccountExists(key common.Address) bool {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	utils.HandleError(err)
	defer db.Close()

	exists := false
	err = db.View(func(txn *badger.Txn) error {
		_, err := txn.Get(key[:])
		if err == nil {
			exists = true
		}
		return nil
	})
	if err != nil {
		utils.HandleError(err)
	}
	return exists
}

// function to hash the hash of all account values
func GetStateRoot() common.Hash {
	// Get the state root
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	utils.HandleError(err)
	defer db.Close()

	var stateRoot []common.Hash
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		it := txn.NewIterator(opts)
		defer it.Close()

		it.Rewind()
		for it.Valid() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				stateRoot = append(stateRoot, common.BytesToHash(val))
				return nil
			})
			utils.HandleError(err)
			it.Next()
		}
		return nil
	})
	if err != nil {
		utils.HandleError(err)
	}

	// hash the hash of all account values
	// return common.BytesToHash(rlpHash(stateRoot))
	rootHash := CalculateRootHash(stateRoot)
	return rootHash
}
