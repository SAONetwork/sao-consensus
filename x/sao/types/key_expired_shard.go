package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ExpiredShardKeyPrefix is the prefix to retrieve all ExpiredShard
	ExpiredShardKeyPrefix = "ExpiredShard/value/"
)

// ExpiredShardKey returns the store key to retrieve a ExpiredShard from the index fields
func ExpiredShardKey(
	height uint64,
) []byte {
	var key []byte

	heightBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBytes, height)
	key = append(key, heightBytes...)
	key = append(key, []byte("/")...)

	return key
}
