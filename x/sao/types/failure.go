package types

const (
	FailureKeyPrefix = "Failure/value/"
)

func FailureKey(
	provider string,
) []byte {
	var key []byte

	providerBytes := []byte(provider)
	key = append(key, providerBytes...)
	key = append(key, []byte("/")...)

	return key
}
