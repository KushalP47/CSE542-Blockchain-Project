package blockchain

import (
	"encoding/hex"
	"fmt"
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
	blockNumber := DeserializeBlock(lastBlock).Header.Number
	blockNumber++
	fmt.Println("Block Number: ", blockNumber)
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
		ExtraData:        []byte(""),
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
	err := database.WriteBlock(block.Hash(), serializedBlock)
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

	block, err := database.ReadBlock(hash[:])
	if err != nil {
		return nil, err
	}
	return DeserializeBlock(block), nil
}
