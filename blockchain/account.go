package blockchain

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type Account struct {
	Address string
	Name    string
	Nonce   uint64
	Balance uint64
}

// generateAddress generates a new address for an account which generates a random 20-byte address for an account
func GenerateAddress(name string) string {
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

	return hex.EncodeToString(hash[:])
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
