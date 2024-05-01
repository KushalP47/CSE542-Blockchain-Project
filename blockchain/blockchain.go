package blockchain

import (
	"github.com/KushalP47/CSE542-Blockchain-Project/database"
)

type Blockchain struct {
	lastBlockNumber uint64
	// blockDB is the database for blocks
	// key-> block number, value-> block
	blockDB *database.Database
	// accountsDB is the database for accounts
	// key-> account address, value-> account
	accountsDB *database.Database
	// txnPool is the pool of transactions
	txnPool []*Transaction
}
