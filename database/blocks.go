package database

import (
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
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

func WriteBlock(key common.Hash, value []byte) error {
	// Write the block to the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks"))
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

func LastBlock() (common.Hash, []byte, error) {
	// Get the last block from the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks"))
	utils.HandleError(err)
	defer db.Close()

	var lastBlockHash []byte
	var lastBlock []byte
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Reverse = true
		it := txn.NewIterator(opts)
		defer it.Close()
		c := 0
		// Position the iterator at the end of the database
		for it.Seek([]byte{}); it.Valid(); it.Next() {
			item := it.Item()
			lastBlockHash = item.Key()
			err := item.Value(func(val []byte) error {
				lastBlock = val
				return nil
			})
			if err != nil {
				return err
			}
			c++
			if c == 1 {
				break
			}
		}
		return nil
	})
	if err == badger.ErrKeyNotFound {
		return common.Hash(lastBlockHash), lastBlock, nil
	}

	return common.Hash(lastBlockHash), lastBlock, err
}
func GetBlocksData() ([][]byte, error) {
	// Get all the blocks from the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks"))
	utils.HandleError(err)
	defer db.Close()

	var blocks [][]byte
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Reverse = true
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				// Append the entire val slice as a single element in blocks
				blocks = append(blocks, val)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return blocks, nil
}

func CountBlocks() (uint64, error) {
	// Count the number of blocks in the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks"))
	utils.HandleError(err)
	defer db.Close()

	count := uint64(0)
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Reverse = true
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			count++
			if count == 1 {
				break
			}
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return count, nil
}
