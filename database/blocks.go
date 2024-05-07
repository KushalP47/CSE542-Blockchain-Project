package database

import (
	"fmt"

	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

// key: block number
// value: block data
func ReadBlock(key uint64) ([]byte, error) {
	// Read the block from the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks/blocksData"))
	utils.HandleError(err)
	defer db.Close()

	var block []byte
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(uint64ToBytes(key))
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
	fmt.Printf("Block: %v\n", block)
	return block, err
}

func WriteBlock(key uint64, value []byte) error {
	// Write the block to the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks/blocksData"))
	utils.HandleError(err)
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set(uint64ToBytes(key), value)
		utils.HandleError(err)
		return err
	})
	utils.HandleError(err)
	return err
}

func LastBlock() (uint64, []byte, error) {
	// Get the last block from the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks/blocksData"))
	utils.HandleError(err)
	defer db.Close()

	var lastBlockNumber uint64
	var lastBlock []byte
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Reverse = true
		it := txn.NewIterator(opts)
		defer it.Close()

		// Position the iterator at the end of the database
		it.Seek([]byte{})
		if it.Valid() {
			item := it.Item()
			lastBlockNumber = bytesToUint64(item.Key())
			err := item.Value(func(val []byte) error {
				lastBlock = val
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	return lastBlockNumber, lastBlock, err
}

func CheckIfBlockExists(key uint64) bool {
	// Check if the block exists in the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks/blocksData"))
	utils.HandleError(err)
	defer db.Close()

	exists := false
	err = db.View(func(txn *badger.Txn) error {
		_, err := txn.Get(uint64ToBytes(key))
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

// key: block Hash
// value: block Data
func WriteBlockHash(key common.Hash, value []byte) error {
	// Write the block hash to the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks/blocksHash"))
	utils.HandleError(err)
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		bytesKey, err := rlp.EncodeToBytes(key)
		utils.HandleError(err)
		err = txn.Set(bytesKey, value)
		utils.HandleError(err)
		return err
	})
	utils.HandleError(err)
	return err
}

func ReadBlockHash(key common.Hash) ([]byte, error) {
	// Read the block hash from the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks/blocksHash"))
	utils.HandleError(err)
	defer db.Close()

	var block []byte
	err = db.View(func(txn *badger.Txn) error {
		bytesKey, err := rlp.EncodeToBytes(key)
		utils.HandleError(err)
		item, err := txn.Get(bytesKey)
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

func BlockExists(key common.Hash) bool {
	// Check if the block exists in the database
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/blocks/blocksHash"))
	utils.HandleError(err)
	defer db.Close()

	exists := false
	err = db.View(func(txn *badger.Txn) error {
		bytesKey, err := rlp.EncodeToBytes(key)
		utils.HandleError(err)
		_, err = txn.Get(bytesKey)
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
