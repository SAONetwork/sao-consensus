package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// OrderFinishKeyPrefix is the prefix to retrieve all OrderFinish
	OrderFinishKeyPrefix = "OrderFinish/value/"
)

// OrderFinishKey returns the store key to retrieve a OrderFinish from the index fields
func OrderFinishKey(
	timestamp uint64,
) []byte {
	var key []byte

	timestampBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBytes, timestamp)
	key = append(key, timestampBytes...)
	key = append(key, []byte("/")...)

	return key
}
