package api

import (
	"fmt"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
)

func CreateGenesisBlock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating Genesis Block")
	block := blockchain.CreateGenesisBlock()
	fmt.Println("Genesis Block Created")
	err := blockchain.AddBlock(block)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Genesis Block Added to the Blockchain")
}
