package main

import (
	"log"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/api"
)

func main() {
	// Define HTTP handlers
	http.HandleFunc("/addAccount", api.AddAccountHandler)
	http.HandleFunc("/getAccount", api.GetAccountHandler)
	http.HandleFunc("/getNonce", api.GetNonceHandler)
	http.HandleFunc("/addTxn", api.AddTxnHandler)
	http.HandleFunc("/getBalance", api.GetBalanceHandler)
	// gives latest block number
	http.HandleFunc("/getBlockNumber", api.GetBlockNumberHandler)
	// gives block details when block number is given
	http.HandleFunc("/getBlockWithNumber", api.GetBlockWithNumberHandler)
	// gives block details when block hash is given
	http.HandleFunc("/getBlockWithHash", api.GetBlockWithHashHandler)
	http.HandleFunc("/createGenesisBlock", api.CreateGenesisBlock)
	// Start the HTTP server
	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
