package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DidKeyPrefix is the prefix to retrieve all Did
	DidKeyPrefix = "Did/value/"
)

// DidKey returns the store key to retrieve a Did from the index fields
func DidKey(
	accountId string,
) []byte {
	var key []byte

	accountIdBytes := []byte(accountId)
	key = append(key, accountIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
