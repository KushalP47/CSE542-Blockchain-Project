package blockchain

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"os"
	"sync"

	"github.com/KushalP47/CSE542-Blockchain-Project/database"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

type Txn struct {
	To    common.Address // Address of the receiver
	Value uint64         // Amount of tokens to send
	Nonce uint64         // Nonce of the sender
}

type SignedTx struct {
	To      common.Address
	Value   uint64
	Nonce   uint64
	V, R, S *big.Int // signature values
}

func SerializeTxn(txn SignedTx) ([]byte, error) {
	encodedTxn, err := rlp.EncodeToBytes(txn)
	if err != nil {
		return nil, err
	}
	return encodedTxn, err
}

func DeserializeTxn(data []byte) (SignedTx, error) {
	txn := SignedTx{}
	err := rlp.DecodeBytes(data, &txn)
	if err != nil {
		return SignedTx{}, err
	}
	return txn, err
}

// VerifyTxn verifies if the sender has enough balance to send the given amount of tokens
func VerifyTxn(sender common.Address, txn Txn) bool {
	// get the balance of the sender
	account, err := GetAccount(sender)
	if err != nil {
		panic(err)
	}
	if account.Balance < txn.Value {
		return false
	}
	return true
}

func AddTxn(signedTxn SignedTx) error {

	txnHash := hashSigned(&signedTxn)
	serializedTxn, err := SerializeTxn(signedTxn)
	if err != nil {
		panic(err)
	}
	err = database.WriteTxn(txnHash[:], serializedTxn)
	if err != nil {
		panic(err)
	}
	//  add the Txn in transactions.txt
	f, err := os.OpenFile("./database/tmp/transactions/transactions.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write the signed transaction to the file
	_, err = fmt.Fprintf(f, "TxnHash: %s\n To: %s\n Value: %d\n Nonce: %d\n V: %d\n R: %d\n S: %d\n \n \n", txnHash.String(), signedTxn.To.String(), signedTxn.Value, signedTxn.Nonce, signedTxn.V, signedTxn.R, signedTxn.S)

	return err
}

func GetTxn(txnHash common.Hash) (SignedTx, error) {
	serializedTxn, err := database.ReadTxn(txnHash[:])
	if err != nil {
		return SignedTx{}, err
	}
	txn, err := DeserializeTxn(serializedTxn)
	return txn, err
}

// rlpDecode decodes the RLP-encoded bytes and stores the result in the given interface.
func RlpDecode(data []byte, val interface{}) error {
	reader := bytes.NewReader(data)
	if err := rlp.Decode(reader, val); err != nil {
		return err
	}
	return nil
}

// recoverPlain recovers the address which has signed the given data using the v, r, and s values
func GetSenderAddress(tx *Txn, R, S, Vb *big.Int, homestead bool) (common.Address, error) {
	sighash := hash(tx)
	if Vb.BitLen() > 8 {
		// return common.Address{}, ErrInvalidSig
		panic("invalid signature")
	}
	V := byte(Vb.Uint64() - 27)
	if !crypto.ValidateSignatureValues(V, R, S, homestead) {
		// return common.Address{}, ErrInvalidSig
		panic("invalid signature")
	}
	// encode the signature in uncompressed format
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, crypto.SignatureLength)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = V
	// recover the public key from the signature
	pub, err := crypto.Ecrecover(sighash[:], sig)
	if err != nil {
		return common.Address{}, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return common.Address{}, errors.New("invalid public key")
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(pub[1:])[12:])
	return addr, nil
}

func hash(tx *Txn) common.Hash {
	return rlpHash([]interface{}{
		tx.To,
		tx.Value,
		tx.Nonce,
	})
}

// HashSigned returns the tx hash
func hashSigned(tx *SignedTx) common.Hash {
	return rlpHash(tx)
}

func rlpHash(x interface{}) (h common.Hash) {
	sha := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	rlp.Encode(sha, x)
	sha.Read(h[:])

	return h
}

// hasherPool holds LegacyKeccak256 hashers for rlpHash.
var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}
