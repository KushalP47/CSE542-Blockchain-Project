package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
)

func AddAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var req struct {
		Address string `json:"address"`
		Balance uint64 `json:"balance"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	account, err := blockchain.CreateAccount(req.Address, req.Balance)
	if err != nil {
		http.Error(w, "Error adding account", http.StatusInternalServerError)
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
	// Append account details to file
	f, err := os.OpenFile("database/tmp/accounts.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "ByteAddress: %s\n Address: %s\n Nonce: %d\n Balance: %d\n\n", account.Address, resp.Address, resp.Nonce, resp.Balance)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(resp)
}
