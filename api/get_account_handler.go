package api

import (
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/database"
)

func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request query parameter
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing address query parameter", http.StatusBadRequest)
		return
	}

	// Get account from database
	account := database.ReadAccount([]byte(address))

	// Respond with account details
	json.NewEncoder(w).Encode(account)
}
