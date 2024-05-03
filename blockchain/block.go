package blockchain

import (
	"io"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

type Header struct {
	ParentHash       common.Hash
	Miner            common.Address
	StateRoot        common.Hash
	TransactionsRoot common.Hash
	Number           uint64
	Timestamp        uint64
	ExtraData        []byte
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

func SignHeader(h Header) []byte {
	// Sign the header
	privateKey, err := crypto.HexToECDSA(os.Getenv("MINER_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}
	sig, err := crypto.Sign(SealHash(&h).Bytes(), privateKey)
	if err != nil {
		panic(err)
	}
	return sig
}

func VerifyHeaderSig(h Header, sig []byte) bool {
	// Verify the signature of the header
	pubKey, err := crypto.Ecrecover(SealHash(&h).Bytes(), sig)
	if err != nil {
		panic(err)
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubKey[1:])[12:])
	return signer != h.Miner
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeHeader(hasher, header)
	hasher.Sum(hash[:0])

	return hash
}

func encodeHeader(w io.Writer, header *Header) {
	enc := []interface{}{
		header.Number,
		header.StateRoot,
		header.TransactionsRoot,
		header.ParentHash,
		header.Miner,
		header.Timestamp,
	}

	if err := rlp.Encode(w, enc); err != nil {
		panic("can't encode: " + err.Error())
	}
}
