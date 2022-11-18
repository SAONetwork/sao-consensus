package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DidBindingProofsKeyPrefix is the prefix to retrieve all DidBindingProofs
	DidBindingProofsKeyPrefix = "DidBindingProofs/value/"
)

// DidBindingProofsKey returns the store key to retrieve a DidBindingProofs from the index fields
func DidBindingProofsKey(
	accountId string,
) []byte {
	var key []byte

	accountIdBytes := []byte(accountId)
	key = append(key, accountIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
