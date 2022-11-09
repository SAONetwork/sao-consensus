package types

const (
	// ModuleName defines the module name
	ModuleName = "node"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_node"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PoolKey = "Pool/value/"
)

const (
	NodeEventCreator = "creator"
	NodeEventPeer    = "peer"
)

const (
	LoginEventType  = "node-login"
	LogoutEventType = "node-logout"
	ResetEventType  = "node-reset"
)
