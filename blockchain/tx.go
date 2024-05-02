package blockchain

import "math/big"

type Txn struct {
	to      string
	value   uint64
	nonce   uint64
	v, r, s *big.Int
}
