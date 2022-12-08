package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ExpiredDataKeyPrefix is the prefix to retrieve all ExpiredData
	ExpiredDataKeyPrefix = "ExpiredData/value/"
)

// ExpiredDataKey returns the store key to retrieve a ExpiredData from the index fields
func ExpiredDataKey(
	height uint64,
) []byte {
	var key []byte

	heightBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBytes, height)
	key = append(key, heightBytes...)
	key = append(key, []byte("/")...)

	return key
}
