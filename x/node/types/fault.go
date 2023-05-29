package types

const (
	FaultKeyPrefix = "Fault/value/"
)

func FaultKey(
	provider string,
	shardId string,
) []byte {
	var key []byte

	providerBytes := []byte(provider)
	shardIdBytes := []byte(shardId)
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
