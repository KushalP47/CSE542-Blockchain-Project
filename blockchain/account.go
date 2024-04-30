package blockchain

type Account struct {
	Address [20]byte
	Nonce   uint64
	Balance uint64
}

type AccountState struct {
	Accounts map[[20]byte]Account // map of address to account
}
