package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// KidKeyPrefix is the prefix to retrieve all Kid
	KidKeyPrefix = "Kid/value/"
)

// KidKey returns the store key to retrieve a Kid from the index fields
func KidKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
