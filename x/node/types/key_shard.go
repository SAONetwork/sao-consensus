package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ShardKeyPrefix is the prefix to retrieve all Shard
	ShardKeyPrefix = "Shard/value/"
)

// ShardKey returns the store key to retrieve a Shard from the index fields
func ShardKey(
	idx string,
) []byte {
	var key []byte

	idxBytes := []byte(idx)
	key = append(key, idxBytes...)
	key = append(key, []byte("/")...)

	return key
}
