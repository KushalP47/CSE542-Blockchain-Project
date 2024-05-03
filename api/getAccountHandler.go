package api

import (
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
	"github.com/ethereum/go-ethereum/common"
)

func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request query parameter
	var req struct {
		Address common.Address `json:"address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	account, err := blockchain.GetAccount(req.Address)
	if err != nil {
		http.Error(w, "Error getting account", http.StatusInternalServerError)
		return
	}

	// Respond with account address
	resp := struct {
		Address common.Address `json:"address"`
		Nonce   uint64         `json:"nonce"`
		Balance uint64         `json:"balance"`
	}{
		Address: account.Address,
		Nonce:   account.Nonce,
		Balance: account.Balance,
	}
	// Respond with account details
	json.NewEncoder(w).Encode(resp)
}
