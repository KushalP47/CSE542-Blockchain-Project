package blockchain

func CreateBlock(header *Header, transactions []SignedTx) *Block {
	// Create a new block
	return &Block{
		Header:       *header,
		Transactions: transactions,
	}
}
