package api

import (
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/ethereum/go-ethereum/common"
)

func GetBlockNumberHandler(w http.ResponseWriter, r *http.Request) {
	// Get the block number
	lastBlockNumber, _, err := database.LastBlock()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	blockNumber := lastBlockNumber
	resp := struct {
		BlockNumber uint64      `json:"blockNumber"`
		BlockHash   common.Hash `json:"blockHash"`
	}{
		BlockNumber: blockNumber,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
