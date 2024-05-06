package api

import (
	"bytes"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/p2p"
)

func GetKnownHostHandler(w http.ResponseWriter, r *http.Request) {

	// Parse query parameter "address" to get the account address
	addrList := p2p.GetPeerAddrs()
	var buf bytes.Buffer
	for _, addr := range addrList {
		buf.WriteString(addr)
		buf.WriteByte('\n') // Add a newline separator
	}

	// Convert the buffer to a byte slice
	addrBytes := buf.Bytes()

	// message := "Transaction added scessfully"
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(addrBytes))
	if err != nil {
		// Handle error if unable to write to response
		panic(err)
	}
}
