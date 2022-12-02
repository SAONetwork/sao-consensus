package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PaymentAddressKeyPrefix is the prefix to retrieve all PaymentAddress
	PaymentAddressKeyPrefix = "PaymentAddress/value/"
)

// PaymentAddressKey returns the store key to retrieve a PaymentAddress from the index fields
func PaymentAddressKey(
	did string,
) []byte {
	var key []byte

	didBytes := []byte(did)
	key = append(key, didBytes...)
	key = append(key, []byte("/")...)

	return key
}
