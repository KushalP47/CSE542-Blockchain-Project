package blockchain

import (
	"github.com/ethereum/go-ethereum/common"
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
