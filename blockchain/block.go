package blockchain

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/sha3"
)

type Header struct {
	ParentHash       common.Hash
	Miner            common.Address
	StateRoot        common.Hash
	TransactionsRoot common.Hash
	Number           uint64
	Timestamp        uint64
	ExtraData        []byte
}

type Block struct {
	Header       Header
	Transactions []SignedTx
}

func (b *Block) Hash() common.Hash {
	return rlpHash(b)
}

func SerializeBlock(b *Block) []byte {
	// Serialize the block
	serializedBlock, err := rlp.EncodeToBytes(b)
	if err != nil {
		panic(err)
	}
	return serializedBlock
}

func GetBlockWithHash(blockHash common.Hash) (Block, error) {
	// Get the block with the given hash
	block, err := database.ReadBlock(blockHash[:])
	if err != nil {
		return Block{}, err
	}
	return *DeserializeBlock(block), nil
}

func GetBlockWithNumber(blockNumber uint64) (common.Hash, *Block, error) {
	// Get the block with the given number
	blocksData, err := database.GetBlocksData()
	if err != nil {
		return common.Hash{}, nil, err
	}
	for _, blockData := range blocksData {
		block := DeserializeBlock(blockData)
		if block.Header.Number == blockNumber {
			return block.Hash(), block, nil
		}
	}
	return common.Hash{}, nil, nil
}

func DeserializeBlock(serializedBlock []byte) *Block {
	// Deserialize the block
	var block Block
	err := rlp.DecodeBytes(serializedBlock, &block)
	if err != nil {
		panic(err)
	}
	return &block
}

func SerializeHeader(h *Header) []byte {
	// Serialize the header
	serializedHeader, err := rlp.EncodeToBytes(h)
	if err != nil {
		panic(err)
	}
	return serializedHeader
}

func DeserializeHeader(serializedHeader []byte) *Header {
	// Deserialize the header
	var header Header
	err := rlp.DecodeBytes(serializedHeader, &header)
	if err != nil {
		panic(err)
	}
	return &header
}

func SignHeader(h Header) []byte {
	// Sign the header
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pvtKey := os.Getenv("MINER_PRIVATE_KEY")
	fmt.Println("Private Key: ", pvtKey)
	privateKey, err := crypto.HexToECDSA(pvtKey)
	fmt.Println("Private Key: ", privateKey)
	if err != nil {
		panic(err)
	}
	sig, err := crypto.Sign(SealHash(&h).Bytes(), privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Signature: ", sig, "Length: ", len(sig))

	signedBytes, err := rlp.EncodeToBytes(sig)
	if err != nil {
		panic(err)
	}
	fmt.Println("Signed Bytes: ", signedBytes, "Length: ", len(signedBytes))
	return signedBytes
}

func VerifyHeaderSig(h Header, sig []byte) bool {
	// Verify the signature of the header
	pubKey, err := crypto.Ecrecover(SealHash(&h).Bytes(), sig)
	if err != nil {
		panic(err)
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubKey[1:])[12:])
	return signer != h.Miner
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeHeader(hasher, header)
	hasher.Sum(hash[:0])
	fmt.Println("Seal Hash: ", hash, "Length: ", len(hash))
	return hash
}

func encodeHeader(w io.Writer, header *Header) {
	enc := []interface{}{
		header.Number,
		header.StateRoot,
		header.TransactionsRoot,
		header.ParentHash,
		header.Miner,
		header.Timestamp,
		header.ExtraData,
	}

	err := rlp.Encode(w, enc)
	if err != nil {
		panic("can't encode: " + err.Error())
	} else {
		fmt.Println("Encoded Header: ", enc)
	}
}
