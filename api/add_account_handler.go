package api

import (
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/database"
)

func AddAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req struct {
		Name    string `json:"name"`
		Balance uint64 `json:"balance"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate account
	address := database.GenerateAddress(req.Name)
	account := database.Account{
		Name:    req.Name,
		Balance: req.Balance,
		Address: address,
		Nonce:   database.GetLatestNonce() + 1,
	}
	err := database.WriteAccount(account)
	if err != nil {
		http.Error(w, "Error adding account", http.StatusInternalServerError)
		return
	}

	// Respond with account address
	resp := struct {
		Address string `json:"address"`
	}{
		Address: address,
	}
	json.NewEncoder(w).Encode(resp)
}
