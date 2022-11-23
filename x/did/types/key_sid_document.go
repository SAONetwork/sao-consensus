package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SidDocumentKeyPrefix is the prefix to retrieve all SidDocument
	SidDocumentKeyPrefix = "SidDocument/value/"
)

// SidDocumentKey returns the store key to retrieve a SidDocument from the index fields
func SidDocumentKey(
	versionId string,
) []byte {
	var key []byte

	versionIdBytes := []byte(versionId)
	key = append(key, versionIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
