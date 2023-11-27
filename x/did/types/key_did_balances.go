package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// DidBalancesKeyPrefix is the prefix to retrieve all DidBalances
	DidBalancesKeyPrefix = "DidBalances/value/"
)

// DidBalancesKey returns the store key to retrieve a DidBalances from the index fields
func DidBalancesKey(
	did string,
) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
