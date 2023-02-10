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

const (
	EventOrderId       = "order-id"
	EventErrorCode     = "error-code"
	EventErrorInfo     = "error-info"
	OrderEventProvider = "peer"
)

const (
	ShardEventProvider = "provider"
)

const (
	NewOrderEventType        = "new-order"
	CancelOrderEventType     = "cancel-order"
	OrderCompletedEventType  = "order-completed"
	OrderUnexpectedEventType = "order-unexpected"
	OrderDataReadyEventType  = "order-data-ready"
	TerminateOrderEventType  = "terminate-order"

	NewShardEventType       = "new-shard"
	ShardCompletedEventType = "shard-completed"
	RejectShardEventType    = "reject-shard"
	TerminateShardEventType = "terminate-shard"
)
