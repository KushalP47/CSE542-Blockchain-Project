package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Header struct {
	ParentHash       common.Hash
	Miner            common.Address
	StateRoot        common.Hash
	TransactionsRoot common.Hash
	Difficulty       uint64
	TotalDifficulty  uint64
	Number           uint64
	Timestamp        uint64
	ExtraData        []byte
	Nonce            uint64
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
