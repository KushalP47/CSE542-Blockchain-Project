package blockchain

import (
	"bytes"
	"encoding/gob"

	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
)

type Account struct {
	Address []byte
	Name    string
	Nonce   uint64
	Balance uint64
}

// generateAddress generates a new address for an account which generates a random 20-byte address for an account
func generateAddress(name string) []byte {
	address := make([]byte, 20)
	copy(address[:], name)
	return address
}

// SerializeAccount serializes the account struct to []byte
func SerializeAccount(account Account) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(account)
	if err != nil {
		utils.HandleError(err)
	}
	return buf.Bytes()
}

// DeserializeAccount deserializes []byte to account struct
func DeserializeAccount(data []byte) Account {
	var account Account
	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&account)
	if err != nil {
		utils.HandleError(err)
	}
	return account
}
