package types

import "encoding/binary"

const (
	FaultKeyPrefix = "Fault/value/"
)

func FaultKey(
	provider string,
	shardId uint64,
) []byte {
	var key []byte

	providerBytes := []byte(provider)
	shardIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(shardIdBytes, shardId)
	key = append(key, providerBytes...)
	key = append(key, shardIdBytes...)
	key = append(key, []byte("/")...)

	return key
}

const (
	FaultStatusConfirming = 1
	FaultStatusConfirmed  = 2
	FaultStatusRecovering = 3
)
