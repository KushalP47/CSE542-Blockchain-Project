package blockchain

import "math/big"

type Txn struct {
	to      [20]bytes
	value   uint64
	from    [20]bytes
	// nonce   uint64
	// v, r, s *big.Int
}

// TransferFunds transfers funds from one account to another
func TransferFunds(txn Txn) error {
	// Read sender and receiver accounts
	sender := ReadAccount(txn.From)
	receiver := ReadAccount(txn.To)

	// Check if sender account exists
	if sender.Name == "" {
		return errors.New("sender account does not exist")
	}

	// Check if receiver account exists
	if receiver.Name == "" {
		return errors.New("receiver account does not exist")
	}

	// Check if sender has sufficient balance
	if sender.Balance < txn.Value {
		return errors.New("insufficient balance in sender account")
	}

	// Deduct amount from sender's balance
	sender.Balance -= txn.Value
	sender.Nonce++

	// Add amount to receiver's balance
	receiver.Balance += txn.Value

	// Update sender and receiver accounts
	err := UpdateAccount(txn.From, sender)
	if err != nil {
		return fmt.Errorf("error updating sender account: %v", err)
	}
	err = UpdateAccount(txn.To, receiver)
	if err != nil {
		return fmt.Errorf("error updating receiver account: %v", err)
	}

	return nil
}

