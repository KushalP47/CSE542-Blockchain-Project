package api

import (
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
	"github.com/ethereum/go-ethereum/common"
)

// AddTxnHandler handles transaction requests
func ReadTxnHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TxnHash common.Hash `json:"hash"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get the transaction
	txn, err := blockchain.GetTxn(req.TxnHash)
	if err == nil && txn == (blockchain.SignedTx{}) {
		http.Error(w, "transaction does not exist", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}
