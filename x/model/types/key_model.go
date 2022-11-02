package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ModelKeyPrefix is the prefix to retrieve all Model
	ModelKeyPrefix = "Model/value/"
)

// ModelKey returns the store key to retrieve a Model from the index fields
func ModelKey(
	_key string,
) []byte {
	var key []byte

	keyBytes := []byte(_key)
	key = append(key, keyBytes...)
	key = append(key, []byte("/")...)

	return key
}
