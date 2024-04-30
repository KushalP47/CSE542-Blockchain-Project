package database

import (
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger/v3"
)

func InitDB() (*badger.DB, error) {
	opts := badger.DefaultOptions("/tmp/badger")
	db, err := badger.Open(opts)
	utils.HandleError(err)
	return db, err
}

func CloseDB(db *badger.DB) {
	err := db.Close()
	utils.HandleError(err)
}
