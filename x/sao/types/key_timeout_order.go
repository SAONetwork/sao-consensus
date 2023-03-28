package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TimeoutOrderKeyPrefix is the prefix to retrieve all TimeoutOrder
	TimeoutOrderKeyPrefix = "TimeoutOrder/value/"
)

// TimeoutOrderKey returns the store key to retrieve a TimeoutOrder from the index fields
func TimeoutOrderKey(
	height uint64,
) []byte {
	var key []byte

	heightBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBytes, height)
	key = append(key, heightBytes...)
	key = append(key, []byte("/")...)

	return key
}
