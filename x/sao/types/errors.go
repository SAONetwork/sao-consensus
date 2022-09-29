package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sao module sentinel errors
var (
	ErrInvalidReplica        = sdkerrors.Register(ModuleName, 2100, "replica support 0 ~ 5")
	ErrOrderNotFound         = sdkerrors.Register(ModuleName, 2101, "order id does not exist")
	ErrOrderCompleted        = sdkerrors.Register(ModuleName, 2102, "order already completed")
	ErrOrderCanceled         = sdkerrors.Register(ModuleName, 2103, "order already canceled")
	ErrOrderComplete         = sdkerrors.Register(ModuleName, 2104, "order not waiting completed")
	ErrOrderShardProvider    = sdkerrors.Register(ModuleName, 2105, "not in shard map")
	ErrOrderUnexpectedStatus = sdkerrors.Register(ModuleName, 2106, "invalid order status")
	ErrNotCreator            = sdkerrors.Register(ModuleName, 2107, "invalid msg creator")
	ErrInvalidCid            = sdkerrors.Register(ModuleName, 2108, "invalid ipfs cid")

	ErrShardCompleted        = sdkerrors.Register(ModuleName, 2109, "shard already completed")
	ErrShardUnexpectedStatus = sdkerrors.Register(ModuleName, 2110, "invalid shard status")
)
