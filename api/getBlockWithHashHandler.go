package api

import (
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/ethereum/go-ethereum/common"
)

func GetBlockWithHashHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		BlockHash common.Hash `json:"blockHash"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blockExists := database.BlockExists(req.BlockHash)
	if !blockExists {
		http.Error(w, "block not found", http.StatusNotFound)
		return
	}

	block, err := database.ReadBlockHash(req.BlockHash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Deserialize the block
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
