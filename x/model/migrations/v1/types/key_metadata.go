package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// MetadataKeyPrefix is the prefix to retrieve all Metadata
	MetadataKeyPrefix = "Metadata/value/"
)

// MetadataKey returns the store key to retrieve a Metadata from the index fields
func MetadataKey(
	dataId string,
) []byte {
	var key []byte

	dataIdBytes := []byte(dataId)
	key = append(key, dataIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
