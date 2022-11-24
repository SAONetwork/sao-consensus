package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PastSeedsKeyPrefix is the prefix to retrieve all PastSeeds
	PastSeedsKeyPrefix = "PastSeeds/value/"
)

// PastSeedsKey returns the store key to retrieve a PastSeeds from the index fields
func PastSeedsKey(
	did string,
) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
