package blockchain

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"

	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger/v3"
)

type Account struct {
	Address [20]byte
	Name    string
	Nonce   uint64
	Balance uint64
}

type Accounts struct {
	Accounts   []Account
	Nonce      uint64
	AccountsDb *database.Database
}

func InitAccounts() *Accounts {
	db := database.Database{}
	db.InitDB("accounts")
	accountsList := make([]Account, 0)
	return &Accounts{
		Accounts:   accountsList,
		Nonce:      0,
		AccountsDb: &db,
	}
}

func (a *Accounts) AddAccount(name string, balance uint64) {
	address, err := GenerateAddress(name)
	utils.HandleError(err)
	account := Account{
		Name:    name,
		Nonce:   a.Nonce,
		Balance: balance,
		Address: address,
	}
	a.Accounts = append(a.Accounts, account)
	a.Nonce++

	// Save account to database
	key := address[:]
	err = a.WriteAccount(key, account)
	utils.HandleError(err)
}

func GenerateAddress(name string) ([20]byte, error) {
	// Generate a random 32-byte salt
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		return [20]byte{}, err
	}

	// Combine name and salt
	data := []byte(name)
	data = append(data, salt...)

	// Calculate SHA-256 hash
	hash := sha256.Sum256(data)

	// Truncate hash to 20 bytes
	var address [20]byte
	copy(address[:], hash[:20])

	return address, nil
}

// Read retrieves the value associated with the given key from the database.
func (a *Accounts) ReadAccount(key []byte) (*Account, error) {
	var account Account
	db := a.AccountsDb
	err := db.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			return json.Unmarshal(val, &account)
		})
	})
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// Write stores the given value in the database with the specified key.
func (a *Accounts) WriteAccount(key []byte, account Account) error {
	val, err := json.Marshal(account)
	db := a.AccountsDb
	if err != nil {
		return err
	}
	return db.DB.Update(func(txn *badger.Txn) error {
		return txn.Set(key, val)
	})
}
