package blockchain

import "math/big"

type Block struct {
	Header       Header
	Transactions []Transaction
}

type Header struct {
	parentHash      [32]byte
	miner           [20]byte
	stateRoot       [32]byte
	transactionRoot [32]byte
	difficulty      uint64
	totalDifficulty uint64
	number          uint64
	timestamp       uint64
	extraData       []byte
	nonce           uint64
}

type Transaction struct {
	to      [20]byte
	amount  uint64
	data    []byte
	v, r, s *big.Int
}
