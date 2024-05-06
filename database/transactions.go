package database

import (
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

// key: Number of transactions
// value: transaction data
func WriteTxnData(key uint64, value []byte) error {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions/transactionsData"))
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

func ReadTxnData(key uint64) ([]byte, error) {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions/transactionsData"))
	utils.HandleError(err)
	defer db.Close()

	var Txn []byte
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(uint64ToBytes(key))
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

func TxnDataExists(key uint64) bool {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions/transactionsData"))
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

func GetTxnsData() (map[uint64][]byte, error) {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions/transactionsData"))
	utils.HandleError(err)
	defer db.Close()

	txns := make(map[uint64][]byte)
	err = db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Reverse = false
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
			txns[bytesToUint64(item.Key())] = txnData

			count++
		}
		return nil
	})
	if err != nil {
		utils.HandleError(err)
	}
	return txns, err
}

func DeleteTxnData(key uint64) error {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions/transactionsData"))
	utils.HandleError(err)
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Delete(uint64ToBytes(key))
		utils.HandleError(err)
		return err
	})
	utils.HandleError(err)
	return err
}

func GetLastTxnKey() (uint64, error) {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions/transactionsData"))
	utils.HandleError(err)
	defer db.Close()

	count := uint64(0)
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.Reverse = true
		it := txn.NewIterator(opts)
		defer it.Close()

		it.Seek([]byte{})
		if it.Valid() {
			item := it.Item()
			count = bytesToUint64(item.Key())
		}
		return nil
	})
	if err != nil {
		utils.HandleError(err)
	}
	return count, err
}

// key: Transaction Hash
// value: Transaction Data
func WriteTxnHash(key common.Hash, value []byte) error {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions/transactionsHash"))
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

func ReadTxnHash(key common.Hash) ([]byte, error) {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions/transactionsHash"))
	utils.HandleError(err)
	defer db.Close()

	var Txn []byte
	err = db.View(func(txn *badger.Txn) error {
		bytesKey, err := rlp.EncodeToBytes(key)
		utils.HandleError(err)
		item, err := txn.Get(bytesKey)
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

// key: Transaction Hash
// value: Transaction Number
func WriteTxnNumber(key common.Hash, value uint64) error {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions/transactionsNumber"))
	utils.HandleError(err)
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		bytesKey, err := rlp.EncodeToBytes(key)
		utils.HandleError(err)
		err = txn.Set(bytesKey, uint64ToBytes(value))
		utils.HandleError(err)
		return err
	})
	utils.HandleError(err)
	return err
}

func TxnNumberExists(key common.Hash) bool {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions/transactionsNumber"))
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

func ReadTxnNumber(key common.Hash) (uint64, error) {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/transactions/transactionsNumber"))
	utils.HandleError(err)
	defer db.Close()

	var Txn uint64
	err = db.View(func(txn *badger.Txn) error {
		bytesKey, err := rlp.EncodeToBytes(key)
		utils.HandleError(err)
		item, err := txn.Get(bytesKey)
		utils.HandleError(err)
		err = item.Value(func(val []byte) error {
			Txn = bytesToUint64(val)
			return nil
		})
		return err
	})
	if err == badger.ErrKeyNotFound {
		return 0, nil
	} else if err != nil {
		utils.HandleError(err)
	}
	return Txn, err
}
