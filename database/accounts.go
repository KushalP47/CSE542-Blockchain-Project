package database

import (
	// "github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"

	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
	"github.com/dgraph-io/badger"
)

// key: account address
// value: account

type Account struct {
	Address [20]byte
	Name    string
	Nonce   uint64
	Balance uint64
}

func ReadAccount(address [20]byte) Account {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	utils.HandleError(err)
	defer db.Close()

	var account Account
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(address[:])
		utils.HandleError(err)
		err = item.Value(func(val []byte) error {
			account, err = DeserializeAccount(val)
			return nil
		})
		return err
	})
	utils.HandleError(err)
	return account
}

func WriteAccount(account Account) error {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	utils.HandleError(err)
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		serialized, _ := SerializeAccount(account)
		err := txn.Set(account.Address[:], serialized)
		utils.HandleError(err)
		return err
	})
	utils.HandleError(err)
	return err
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

func UpdateAccount(address [20]byte, account Account) error {
	db, err := badger.Open(badger.DefaultOptions("./database/tmp/accounts"))
	utils.HandleError(err)
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		serialized, _ := SerializeAccount(account)
		err := txn.Set(address[:], serialized)
		utils.HandleError(err)
		return err
	})
	utils.HandleError(err)
	return err
}

// SerializeAccount serializes the account struct to []byte
func SerializeAccount(account Account) ([]byte, error) {
	return json.Marshal(account)
}

// DeserializeAccount deserializes []byte to account struct
func DeserializeAccount(data []byte) (Account, error) {
	var account Account
	err := json.Unmarshal(data, &account)
	if err != nil {
		return Account{}, err
	}
	return account, nil
}

func GenerateAddress(name string) [20]byte {
	// Generate a random salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}

	// Append salt to the name
	data := []byte(name)
	data = append(data, salt...)

	// Calculate SHA-256 hash
	hash := sha256.Sum256(data)

	var address [20]byte
	copy(address[:], hash[:20]) // Copy first 20 bytes of the hash to the address

	return address
}
