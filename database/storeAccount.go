package database

import (
	"github.com/KushalP47/CSE542-Blockchain-Project/Blockchain"
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger/v3"
)

func StoreAccount(db *badger.DB, address [20]byte, account blockchain.Account) error {

}
