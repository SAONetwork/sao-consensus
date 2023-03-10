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
	OrderMigrating
)

const (
	ShardWaiting = iota
	ShardRejected
	ShardCompleted
	ShardTerminated
)
