package types

const (
	// ModuleName defines the module name
	ModuleName = "sao"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_sao"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	OrderKey      = "Order/value/"
	OrderCountKey = "Order/count/"
)

const (
	ShardKey      = "Shard/value/"
	ShardCountKey = "Shard/count/"
)

const (
	EventCid          = "cid"
	OrderEventCreator = "creator"
)

const (
	EventOrderId          = "order-id"
	EventTimeoutOrderList = "order-list"
	EventErrorCode        = "error-code"
	EventErrorInfo        = "error-info"
	OrderEventProvider    = "peer"
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
	OrderTimeoutEventType   = "order-timeout"

	FaultsReportedEventType        = "faults-reported"
	FaultsConfirmedEventType       = "faults-confirmed"
	FaultsRecoverDeclaredEventType = "faults-recover-declared"
	FaultsRecoveredEventType       = "faults-recovered"
)
