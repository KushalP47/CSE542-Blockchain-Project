package database

import (
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger"
)

func ReadBlock(hash []byte) ([]byte, error) {
	// Read the block from the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks"))
	utils.HandleError(err)
	defer db.Close()

	var block []byte
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(hash)
		utils.HandleError(err)
		err = item.Value(func(val []byte) error {
			block = val
			return nil
		})
		return err
	})
	if err == badger.ErrKeyNotFound {
		return nil, nil
	}

	return block, err
}

func WriteBlock(key []byte, value []byte) error {
	// Write the block to the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks"))
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

func LastBlockHash() ([]byte, error) {
	// Get the last block from the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks"))
	utils.HandleError(err)
	defer db.Close()

	var lastBlockHash []byte
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Reverse = true
		it := txn.NewIterator(opts)
		defer it.Close()

		it.Rewind()
		if it.Valid() {
			item := it.Item()
			lastBlockHash = item.Key()
			utils.HandleError(err)
		}
		return nil
	})
	if err == badger.ErrKeyNotFound {
		return nil, nil
	}

	return lastBlockHash, err
}
