package api

// AddTxnHandler handles transaction requests
// func AddTxnHandler(w http.ResponseWriter, r *http.Request) {
// 	decoder := json.NewDecoder(r.Body)
// 	var txnData struct {
// 		From  [20]byte `json:"from"`
// 		To    [20]byte `json:"to"`
// 		Value uint64   `json:"value"`
// 	}

// 	err := decoder.Decode(&txnData)
// 	if err != nil {
// 		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
// 		return
// 	}

// 	// Check if sender account exists
// 	fromAccount := database.ReadAccount(txnData.From)
// 	if fromAccount.Name == "" {
// 		utils.RespondWithError(w, http.StatusBadRequest, "Sender account does not exist")
// 		return
// 	}

// 	// Check if receiver account exists
// 	toAccount := database.ReadAccount(txnData.To)
// 	if toAccount.Name == "" {
// 		// Generate new address for receiver
// 		txnData.To = database.GenerateAddress("NewAccount")

// 		// Check if sender is equal to receiver
// 		if bytes.Equal(txnData.From[:], txnData.To[:]) {
// 			utils.RespondWithError(w, http.StatusBadRequest, "Sender and receiver addresses cannot be the same")
// 			return
// 		}
// 	}

// 	// Perform the transaction
// 	txn := database.Txn{
// 		To:    txnData.To,
// 		Value: txnData.Value,
// 		From:  txnData.From,
// 	}
// 	err = database.TransferFunds(txn)
// 	if err != nil {
// 		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Transaction successful"})
// }
