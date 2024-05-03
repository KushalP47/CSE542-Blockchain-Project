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
	// http.HandleFunc("/printAccount", api.PrintAccountHandler)

	// Start the HTTP server
	log.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
