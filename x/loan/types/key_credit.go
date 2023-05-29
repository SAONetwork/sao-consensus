package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CreditKeyPrefix is the prefix to retrieve all Credit
	CreditKeyPrefix = "Credit/value/"
)

// CreditKey returns the store key to retrieve a Credit from the index fields
func CreditKey(
	account string,
) []byte {
	var key []byte

	accountBytes := []byte(account)
	key = append(key, accountBytes...)
	key = append(key, []byte("/")...)

	return key
}
