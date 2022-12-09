package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DidBindingProofKeyPrefix is the prefix to retrieve all DidBindingProof
	DidBindingProofKeyPrefix = "DidBindingProof/value/"
)

// DidBindingProofKey returns the store key to retrieve a DidBindingProof from the index fields
func DidBindingProofKey(
	accountId string,
) []byte {
	var key []byte

	accountIdBytes := []byte(accountId)
	key = append(key, accountIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
