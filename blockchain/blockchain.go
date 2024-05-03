package blockchain

import (
	"os"
	"time"

	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/ethereum/go-ethereum/common"
)

func CreateBlock() *Block {
	// Create a new block

	txns, err := database.GetTxns()
	if err != nil {
		panic(err)
	}
	var newTxns []SignedTx
	var txnsHashes []common.Hash
	for txnHash, txn := range txns {
		txnsHashes = append(txnsHashes, txnHash)
		deserializedTxn, err := DeserializeTxn(txn)
		if err != nil {
			panic(err)
		}
		newTxns = append(newTxns, deserializedTxn)
	}
	lastBlockHash, lastBlock, err := database.LastBlock()
	parentHash := lastBlockHash
	if err != nil {
		panic(err)
	}
	minerAddress := os.Getenv("MINER_ADDRESS")
	stateRoot := database.GetStateRoot()
	transactionsRoot := database.CalculateRootHash(txnsHashes)
	blockNumber := DeserializeBlock(lastBlock).Header.Number + 1
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
		Transactions: newTxns,
	}

	return &block

}
