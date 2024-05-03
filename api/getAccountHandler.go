package api

import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
)

func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request query parameter
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing address query parameter", http.StatusBadRequest)
		return
	}
	account, err := blockchain.GetAccount(address)
	if err != nil {
		http.Error(w, "Error getting account", http.StatusInternalServerError)
		return
	}

	// Respond with account address
	resp := struct {
		Address string `json:"address"`
		Nonce   uint64 `json:"nonce"`
		Balance uint64 `json:"balance"`
	}{
		Address: hex.EncodeToString(account.Address[:]),
		Nonce:   account.Nonce,
		Balance: account.Balance,
	}
	// Respond with account details
	json.NewEncoder(w).Encode(resp)
}
