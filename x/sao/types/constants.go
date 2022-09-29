package types

const (
	OrderPending = iota
	OrderInProgress
	OrderUnexpected
	OrderCompleted
	OrderCanceled
	OrderExpired
	OrderTerminated
)

const (
	ShardWaiting = iota
	ShardRejected
	ShardCompleted
)
