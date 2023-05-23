package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PledgeDebtKeyPrefix is the prefix to retrieve all PledgeDebt
	PledgeDebtKeyPrefix = "PledgeDebt/value/"
)

// PledgeDebtKey returns the store key to retrieve a PledgeDebt from the index fields
func PledgeDebtKey(
	sp string,
) []byte {
	var key []byte

	spBytes := []byte(sp)
	key = append(key, spBytes...)
	key = append(key, []byte("/")...)

	return key
}
