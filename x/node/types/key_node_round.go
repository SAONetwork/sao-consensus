package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// NodeRoundKeyPrefix is the prefix to retrieve super node round.
	NodeRoundKeyPrefix = "NodeRound/value/"
)

// TODO: super node < 255?
// NodeKey returns the store key to retrieve a Node from the index fields
func NodeRoundKey() []byte {
	var key []byte

	round := []byte("round")
	key = append(key, round...)
	key = append(key, []byte("/")...)

	return key
}
