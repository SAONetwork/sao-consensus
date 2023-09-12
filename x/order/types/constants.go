package types

const (
	OrderPending = iota
	OrderInProgress
	OrderUnexpected
	OrderCompleted
	OrderCanceled
	OrderExpired
	OrderDataReady
	OrderTerminated
)

const (
	TextOrderPending    = "Pending"
	TextOrderInProgress = "InProgress"
	TextOrderUnexpected = "Unexpected"
	TextOrderCompleted  = "Completed"
	TextOrderCanceled   = "Canceled"
	TextOrderExpired    = "Expired"
	TextOrderDataReady  = "DataReady"
	TextOrderTerminated = "Terminated"
)

const (
	ShardWaiting = iota
	ShardRejected
	ShardCompleted
	ShardTerminated
	ShardMigrating
	ShardTimeout
)

const (
	TextShardWaiting    = "Waiting"
	TextShardRejected   = "Rejeted"
	TextShardCompleted  = "Completed"
	TextShardTerminated = "Terminated"
	TextShardMigrating  = "Migrating"
	TextShardTimeout    = "Timeout"
)
