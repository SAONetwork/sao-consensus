package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AccountIdKeyPrefix is the prefix to retrieve all AccountId
	AccountIdKeyPrefix = "AccountId/value/"
)

// AccountIdKey returns the store key to retrieve a AccountId from the index fields
func AccountIdKey(
	accountDid string,
) []byte {
	var key []byte

	accountDidBytes := []byte(accountDid)
	key = append(key, accountDidBytes...)
	key = append(key, []byte("/")...)

	return key
}
