package blockchain

import (
	"github.com/KushalP47/CSE542-Blockchain-Project/database"
)

type Blockchain struct {
	// blocks is a list of blocks in the blockchain
	blocks []*Block
	// blockDB is the database for blocks
	blockDB *database.Database
	// accountState is the state of all accounts in the blockchain
	accounts []*Account
	// accountsDB is the database for accounts
	accountsDB *database.Database
	// txnPool is the pool of transactions
	txnPool []*Transaction
}
