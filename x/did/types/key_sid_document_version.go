package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// SidDocumentVersionKeyPrefix is the prefix to retrieve all SidDocumentVersion
	SidDocumentVersionKeyPrefix = "SidDocumentVersion/value/"
)

// SidDocumentVersionKey returns the store key to retrieve a SidDocumentVersion from the index fields
func SidDocumentVersionKey(
	docId string,
) []byte {
	var key []byte

	docIdBytes := []byte(docId)
	key = append(key, docIdBytes...)
	key = append(key, []byte("/")...)

	return key
}
