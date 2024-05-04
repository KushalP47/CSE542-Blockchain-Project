package blockchain

import (
	"fmt"
	"os"

	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Account struct {
	Address common.Address
	Nonce   uint64
	Balance uint64
}

// SerializeAccount serializes an account into a byte slice
func SerializeAccount(account Account) ([]byte, error) {
	encodedAccount, err := rlp.EncodeToBytes(account)
	if err != nil {
		return nil, err
	}
	return encodedAccount, err
}

// DeserializeAccount deserializes a byte slice into an account
func DeserializeAccount(data []byte) (Account, error) {
	account := Account{}
	err := rlp.DecodeBytes(data, &account)
	if err != nil {
		return Account{}, err
	}
	return account, err
}

// getLatestNonce returns the nonce of the last account of database
func GetLatestNonce() uint64 {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	utils.HandleError(err)
	defer db.Close()

	var latestNonce uint64
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			err := item.Value(func(val []byte) error {
				account, _ := DeserializeAccount(val)
				if account.Nonce > latestNonce {
					latestNonce = account.Nonce
				}
				return nil
			})
			utils.HandleError(err)
		}
		return nil
	})
	utils.HandleError(err)
	return latestNonce
}

// ReadAccount reads an account from the database
func GetAccount(address common.Address) (Account, error) {
	bytesAccout, err := database.ReadAccount(address)
	if err != nil {
		return Account{}, err
	}
	account, err := DeserializeAccount(bytesAccout)
	return account, err
}

// WriteAccount writes an account to the database
func SetAccount(account Account) error {

	bytesAccount, err := SerializeAccount(account)
	utils.HandleError(err)
	err = database.WriteAccount(account.Address, bytesAccount)
	utils.HandleError(err)

	f, err := os.OpenFile("./database/tmp/accounts/accounts.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "Address: %s\n Balance: %d\n Nonce: %d\n \n \n", account.Address.String(), account.Balance, account.Nonce)
	if err != nil {
		return err
	}
	return err
}

// UpdateAccount updates an account in the database
func UpdateAccount(address common.Address, value uint64) error {
	account, err := GetAccount(address)
	utils.HandleError(err)
	account.Balance = value
	err = SetAccount(account)
	utils.HandleError(err)
	return err
}

// CreateAccount creates a new account
func CreateAccount(address common.Address, balance uint64) (Account, error) {
	account := Account{Address: address, Balance: balance, Nonce: GetLatestNonce() + 1}
	err := SetAccount(account)
	utils.HandleError(err)
	return account, err
}
