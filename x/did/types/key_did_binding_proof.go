package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DidBingingProofKeyPrefix is the prefix to retrieve all DidBingingProof
	DidBingingProofKeyPrefix = "DidBingingProof/value/"
)

// DidBingingProofKey returns the store key to retrieve a DidBingingProof from the index fields
func DidBingingProofKey(
	accountId string,
) []byte {
	var key []byte

	accountIdBytes := []byte(accountId)
	key = append(key, accountIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
