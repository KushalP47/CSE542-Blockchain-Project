package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/KushalP47/CSE542-Blockchain-Project/api"
	"github.com/KushalP47/CSE542-Blockchain-Project/p2p"
)

func main() {
	go func() {
		// Define HTTP handlers
		http.HandleFunc("/addAccount", api.AddAccountHandler)
		http.HandleFunc("/getAccount", api.GetAccountHandler)
		http.HandleFunc("/getNonce", api.GetNonceHandler)
		http.HandleFunc("/addTxn", api.AddTxnHandler)
		http.HandleFunc("/getBalance", api.GetBalanceHandler)
		http.HandleFunc("/getBlockNumber", api.GetBlockNumberHandler)
		http.HandleFunc("/getBlockWithNumber", api.GetBlockWithNumberHandler)
		http.HandleFunc("/getBlockWithHash", api.GetBlockWithHashHandler)
		http.HandleFunc("/createGenesisBlock", api.CreateGenesisBlock)
		http.HandleFunc("/getKnownHosts", api.GetKnownHostHandler)

		// Start the HTTP server in a new goroutine

		log.Println("Server started on port 8000")
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

	// node, err := p2p.StartNewNode()
	// if err != nil {
	// 	panic(err)
	// }

	// // Print the node's listening addresses
	// for _, addr := range node.Addrs() {
	// 	println("Listening on", addr)
	// }

	// Create a new context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the libp2p node in a new goroutine

	// Start the libp2p node in a new goroutine
	go func() {
		if len(os.Args) > 1 {
			// If a peer address is provided, start a peer node
			p2p.StartPeerNode(ctx, os.Args[1])
		} else {
			// If no peer address is provided, start a root node
			p2p.StartRootNode(ctx)
		}
	}()

	// Handle termination signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	// Signal the network to stop
	// cancel()
}
