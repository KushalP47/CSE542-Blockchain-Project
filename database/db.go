package database

import (
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger/v3"
)

type Database struct {
	DB *badger.DB
}

func (db *Database) InitDB(dir string) (*badger.DB, error) {
	loc := "./tmp/" + dir
	opts := badger.DefaultOptions(loc)
	var err error
	db.DB, err = badger.Open(opts)
	utils.HandleError(err)
	return db.DB, err
}

func (db *Database) Read(key string) ([]byte, error) {
	var valCopy []byte
	err := db.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		utils.HandleError(err)
		err = item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)
			return nil
		})
		utils.HandleError(err)
		return nil
	})
	return valCopy, err
}

func (db *Database) Write(key string, val []byte) error {
	err := db.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), val)
		utils.HandleError(err)
		return nil
	})
	return err
}

func (db *Database) CloseDB() {
	err := db.DB.Close()
	utils.HandleError(err)
}
