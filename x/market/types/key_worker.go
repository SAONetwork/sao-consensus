package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// WorkerKeyPrefix is the prefix to retrieve all Worker
	WorkerKeyPrefix = "Worker/value/"
)

// WorkerKey returns the store key to retrieve a Worker from the index fields
func WorkerKey(
	workername string,
) []byte {
	var key []byte

	workernameBytes := []byte(workername)
	key = append(key, workernameBytes...)
	key = append(key, []byte("/")...)

	return key
}
