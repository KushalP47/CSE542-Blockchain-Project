package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/KushalP47/CSE542-Blockchain-Project/p2p"
	"github.com/KushalP47/CSE542-Blockchain-Project/pkg/utils"
)

// AddTxnHandler handles transaction requests
func AddTxnHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SignedTxn string `json:"signed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// fmt.Println(req.SignedTxn)
	signedTxn := blockchain.SignedTx{}
	signedTxnHash, err := hex.DecodeString(req.SignedTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = blockchain.RlpDecode(signedTxnHash, &signedTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// verify the transaction
	if !blockchain.VerifySignedTxn(signedTxn) {
		http.Error(w, "Invalid Transaction", http.StatusBadRequest)
		return
	}

	// add Txn to the transactionsData
	txnNumber, err := database.GetLastTxnKey()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	txnNumber++
	err = blockchain.AddTxn(txnNumber, signedTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	serializedTxn, err := blockchain.SerializeTxn(signedTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p2p.SendSignedTransactionToPeer(serializedTxn)
	totalTxns, err := database.GetTxnsData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Total Txns: ", len(totalTxns))
	if len(totalTxns) >= 5 {
		var txns []blockchain.SignedTx
		for _, txn := range totalTxns {
			deserializedTxn, err := blockchain.DeserializeTxn(txn)
			utils.HandleError(err)
			txns = append(txns, deserializedTxn)
		}
		fmt.Println("Creating Block")
		block := blockchain.CreateBlock(txns)
		fmt.Println("Block Created")
		err = blockchain.AddBlock(block)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	w.WriteHeader(http.StatusOK)

}
