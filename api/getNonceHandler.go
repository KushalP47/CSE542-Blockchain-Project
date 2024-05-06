package api

import (
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
	"github.com/ethereum/go-ethereum/common"
)

func GetNonceHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Address common.Address `json:"address"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := blockchain.GetAccount(req.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := struct {
		Nonce uint64 `json:"nonce"`
	}{
		Nonce: account.Nonce,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
