package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ExpiredOrderKeyPrefix is the prefix to retrieve all ExpiredOrder
	ExpiredOrderKeyPrefix = "ExpiredOrder/value/"
)

// ExpiredOrderKey returns the store key to retrieve a ExpiredOrder from the index fields
func ExpiredOrderKey(
	height uint64,
) []byte {
	var key []byte

	heightBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBytes, height)
	key = append(key, heightBytes...)
	key = append(key, []byte("/")...)

	return key
}
