package api

import (
	"encoding/json"
	"net/http"

	"github.com/KushalP47/CSE542-Blockchain-Project/blockchain"
	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/ethereum/go-ethereum/common"
)

// AddTxnHandler handles transaction requests
func AddTxnHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SignedTxn common.Hash `json:"signed"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// decode the signed transaction
	signedTxn := blockchain.SignedTx{}
	if err := blockchain.RlpDecode(req.SignedTxn[:], &signedTxn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	txn := blockchain.Txn{
		To:    signedTxn.To,
		Value: signedTxn.Value,
		Nonce: signedTxn.Nonce,
	}
	// get the address of the sender
	sender, err := blockchain.GetSenderAddress(&txn, signedTxn.R, signedTxn.S, signedTxn.V, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// verify if the sender exists
	senderAccount, err := blockchain.GetAccount(sender)
	if err == nil && senderAccount == (blockchain.Account{}) {
		http.Error(w, "sender does not exist", http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// verify is the sender has enough balance
	if !blockchain.VerifyTxn(sender, txn) {
		http.Error(w, "insufficient balance", http.StatusBadRequest)
		return
	}

	// change the balance of the sender and receiver
	// get the receiver's account
	receiverAccount, err := blockchain.GetAccount(signedTxn.To)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if receiverAccount == (blockchain.Account{}) {
		receiverAccount, err = blockchain.CreateAccount(signedTxn.To, signedTxn.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		receiverAccount.Balance += signedTxn.Value
		err = blockchain.SetAccount(receiverAccount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	_ = receiverAccount

	// get the sender's account
	senderAccount, err = blockchain.GetAccount(sender)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	senderAccount.Balance -= signedTxn.Value
	err = blockchain.SetAccount(senderAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// add Txn to the database
	err = blockchain.AddTxn(signedTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	totalTxns, err := database.GetTxnCount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if totalTxns == 5 {
		var block *blockchain.Block
		if blockchain.CheckIfGenesisBlockExists() {
			block = blockchain.CreateBlock()
		} else {
			block = blockchain.CreateGenesisBlock()
		}
		err = blockchain.AddBlock(block)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	w.WriteHeader(http.StatusOK)

}
