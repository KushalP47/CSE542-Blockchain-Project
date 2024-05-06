package database

import (
	"encoding/binary"

	"github.com/ethereum/go-ethereum/common"
)

func uint64ToBytes(i uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, i)
	return buf
}

func bytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

// Function to calculate the root hash using Hash List structure
func CalculateRootHash(hashes []common.Hash) common.Hash {
	// Iterate through the hashes and re-hash them to get the root hash
	for len(hashes) > 1 {
		var newHashes []common.Hash
		for i := 0; i < len(hashes); i += 2 {
			// Concatenate and hash pairs of hashes
			if i+1 < len(hashes) {
				combinedHash := append(hashes[i][:], hashes[i+1][:]...)
				newHash := common.BytesToHash(combinedHash)
				newHashes = append(newHashes, newHash)
			} else {
				// If there's an odd number of hashes, hash the last hash with itself
				newHashes = append(newHashes, hashes[i])
			}
		}
		hashes = newHashes
	}
	// Return the root hash
	return hashes[0]
}
