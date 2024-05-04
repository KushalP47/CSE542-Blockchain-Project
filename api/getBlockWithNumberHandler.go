package api

import (
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
	"github.com/KushalP47/CSE542-Blockchain-Project/database"
)

func GetBlockWithNumberHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		BlockNumber uint64 `json:"blockNumber"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blockExists := database.CheckIfBlockExists(req.BlockNumber)
	if !blockExists {
		http.Error(w, "block not found", http.StatusNotFound)
		return
	}

	block, err := database.ReadBlock(req.BlockNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	deserializedBlock := blockchain.DeserializeBlock(block)
	resp := struct {
		Block blockchain.Block
	}{
		Block: *deserializedBlock,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
