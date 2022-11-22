package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AccountAuthKeyPrefix is the prefix to retrieve all AccountAuth
	AccountAuthKeyPrefix = "AccountAuth/value/"
)

// AccountAuthKey returns the store key to retrieve a AccountAuth from the index fields
func AccountAuthKey(
	accountDid string,
) []byte {
	var key []byte

	accountDidBytes := []byte(accountDid)
	key = append(key, accountDidBytes...)
	key = append(key, []byte("/")...)

	return key
}
