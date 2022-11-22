package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AccountListKeyPrefix is the prefix to retrieve all AccountList
	AccountListKeyPrefix = "AccountList/value/"
)

// AccountListKey returns the store key to retrieve a AccountList from the index fields
func AccountListKey(
	did string,
) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
