package api

import (
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
)

func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request query parameter
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "Missing address query parameter", http.StatusBadRequest)
		return
	}
	bytesAddress, err := hex.DecodeString(address)
	utils.HandleError(err)

	var arrayAddress [20]byte
	copy(arrayAddress[:], bytesAddress)
	// Get account from database
	account := database.ReadAccount(arrayAddress)

	// Respond with account details
	json.NewEncoder(w).Encode(account)
}
