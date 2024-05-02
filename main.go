package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgraph-io/badger/v4"
)

func main() {



	
	// Initialize Badger DB
	db, err := initializeDatabase()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()

	// Start server
	startServer(db)
}

// initializeDatabase initializes Badger DB and returns a DB object.
func initializeDatabase() (*badger.DB, error) {
	// Example: Open Badger DB
	opts := badger.DefaultOptions("").WithInMemory(true) // In-memory for testing
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// startServer starts the HTTP server.
func startServer(db *badger.DB) {
	// Define HTTP routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the blockchain server!")
	})

	// Start HTTP server
	server := &http.Server{
		Addr:         ":8000",
		Handler:      nil, // Use default handler (http.DefaultServeMux)
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Starting server on port 8080...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
