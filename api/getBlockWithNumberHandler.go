package api

import (
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
	"github.com/ethereum/go-ethereum/common"
)

func GetBlockWithNumberHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		BlockNumber uint64 `json:"blockNumber"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	blockHash, block, err := blockchain.GetBlockWithNumber(req.BlockNumber)
	if blockHash == (common.Hash{}) && block == nil && err == nil {
		http.Error(w, "block not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := struct {
		BlockHash common.Hash `json:"blockHash"`
		Block     blockchain.Block
	}{
		BlockHash: blockHash,
		Block:     *block,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
