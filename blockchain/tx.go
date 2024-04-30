package blockchain

import "math/big"

type Transaction struct {
	to      [20]byte
	amount  uint64
	data    []byte
	v, r, s *big.Int
}
