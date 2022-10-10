package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PledgeKeyPrefix is the prefix to retrieve all Pledge
	PledgeKeyPrefix = "Pledge/value/"
)

// PledgeKey returns the store key to retrieve a Pledge from the index fields
func PledgeKey(
	creator string,
) []byte {
	var key []byte

	creatorBytes := []byte(creator)
	key = append(key, creatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
