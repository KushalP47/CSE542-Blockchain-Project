package blockchain

type Account struct {
	Address [20]byte
	Nonce   uint64
	Balance uint64
}
