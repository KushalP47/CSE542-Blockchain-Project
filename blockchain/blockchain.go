package blockchain

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/ethereum/go-ethereum/common"
)

func CreateBlock(txns []SignedTx) *Block {

	// check if genesis block exists
	if !CheckIfGenesisBlockExists() {
		return CreateGenesisBlock()
	}
	// Create a new block
	var txnsHashes []common.Hash
	for _, txn := range txns {
		txnsHashes = append(txnsHashes, hashSigned(&txn))
	}
	lastBlockNum, lastBlock, err := database.LastBlock()
	parentHash := DeserializeBlock(lastBlock).Hash()
	if err != nil {
		panic(err)
	}
	minerAddress := os.Getenv("MINER_ADDRESS")
	stateRoot := database.GetStateRoot()
	transactionsRoot := database.CalculateRootHash(txnsHashes)
	blockNumber := lastBlockNum
	blockNumber++
	time := time.Now().Unix()
	header := Header{
		ParentHash:       parentHash,
		Miner:            common.HexToAddress(minerAddress),
		StateRoot:        stateRoot,
		TransactionsRoot: transactionsRoot,
		Number:           blockNumber,
		Timestamp:        uint64(time),
	}
	header.ExtraData = SignHeader(header)

	block := Block{
		Header:       header,
		Transactions: txns,
	}

	return &block
}

func ValidateBlock(block *Block) (bool, error) {
	// Validate the block header
	// Validate the transactions
	valid, err := ValidateTxns(block.Transactions)
	if !valid {
		return valid, err
	}
	return true, nil
}

func CreateGenesisBlock() *Block {
	// Create a new block

	minerAddress := os.Getenv("MINER_ADDRESS")
	stateRoot := database.GetStateRoot()
	header := Header{
		ParentHash:       common.Hash{},
		Miner:            common.HexToAddress(minerAddress),
		StateRoot:        stateRoot,
		TransactionsRoot: common.Hash{},
		Number:           0,
		Timestamp:        uint64(time.Now().Unix()),
	}
	header.ExtraData = SignHeader(header)

	block := Block{
		Header:       header,
		Transactions: []SignedTx{},
	}

	return &block
}

func CheckIfGenesisBlockExists() bool {
	// Check if the genesis block exists

	_, _, err := database.LastBlock()
	fmt.Println(err)
	return err == nil
}

func AddBlock(block *Block) error {

	// Add a block to the database
	serializedBlock := SerializeBlock(block)
	err := database.WriteBlock(block.Header.Number, serializedBlock)
	if err != nil {
		return err
	}

	// write the block in blocksHash
	err = database.WriteBlockHash(block.Hash(), serializedBlock)
	if err != nil {
		return err
	}

	// write the block in blocks.txt
	f, err := os.OpenFile("./database/tmp/blocks/blocks.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "BlockHash: %s\n BlockParentHash: %s\n BlockMiner: %s\n BlockStateRoot: %s\n BlockTxnRoot: %s\n BlockNumber: %d\n BlockTimeStamp %d\n BlockExtraData: %s\n \n", block.Hash().String(), block.Header.ParentHash.String(), block.Header.Miner.String(), block.Header.StateRoot.String(), block.Header.TransactionsRoot.String(), block.Header.Number, block.Header.Timestamp, hex.EncodeToString(block.Header.ExtraData))
	if err != nil {
		return err
	}

	return nil
}

func GetBlock(hash common.Hash) (*Block, error) {
	// Get a block from the database

	block, err := database.ReadBlockHash(hash)
	if err != nil {
		return nil, err
	}
	return DeserializeBlock(block), nil
}

func ValidateTxns(txns []SignedTx) (bool, error) {
	// Validate the transactions

	for _, txn := range txns {

		// verify if transaction already exists in the database
		txnHash := hashSigned(&txn)
		if database.TxnNumberExists(txnHash) {
			txnNumber, err := database.ReadTxnNumber(txnHash)
			if err != nil {
				return false, err
			}
			if !database.TxnDataExists(txnNumber) {
				return false, fmt.Errorf("transaction is already committed to block")
			}
			continue
		}

		if !VerifySignedTxn(txn) {
			return false, fmt.Errorf("transaction is not valid")
		}

		// add the transaction in the transactionsHash
		// add Txn to the transactionsData
		txnNumber, err := database.GetLastTxnKey()
		if err != nil {
			return false, err
		}
		txnNumber++
		err = AddTxn(txnNumber, txn)
		if err != nil {
			return false, err
		}

	}
	return true, nil
}

func DeleteTransactions(txns []SignedTx) error {
	// Delete the transactions

	for _, txn := range txns {
		txnHash := hashSigned(&txn)
		txnNumber, err := database.ReadTxnNumber(txnHash)
		if err != nil {
			return err
		}
		err = database.DeleteTxnData(txnNumber)
		if err != nil {
			return err
		}
	}
	return nil
}

func VerifySignedTxn(signedTxn SignedTx) bool {
	// Verify the transaction
	txn := Txn{
		To:    signedTxn.To,
		Value: signedTxn.Value,
		Nonce: signedTxn.Nonce,
	}
	// get the address of the sender
	sender, err := GetSenderAddress(&txn, signedTxn.R, signedTxn.S, signedTxn.V, true)
	if err != nil {
		return false
	}

	// verify if the sender exists
	if !database.AccountExists(sender) {
		return false
	}

	// verify is the sender has enough balance
	if !VerifyTxn(sender, txn) {
		return false
	}

	// change the balance of the sender and receiver
	// get the receiver's account
	var receiverAccount Account
	if !database.AccountExists(signedTxn.To) {
		receiverAccount, err = CreateAccount(signedTxn.To, signedTxn.Value)
		if err != nil {
			return false
		}
	} else {
		receiverAccount, err = GetAccount(signedTxn.To)
		if err != nil {
			return false
		}
		receiverAccount.Balance += signedTxn.Value
		err = SetAccount(receiverAccount)
		if err != nil {
			return false
		}
	}

	// get the sender's account
	senderAccount, err := GetAccount(sender)
	if err != nil {
		return false
	}
	senderAccount.Balance -= signedTxn.Value
	err = SetAccount(senderAccount)
	return err == nil
}
