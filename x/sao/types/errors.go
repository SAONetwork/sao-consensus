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
	ErrSignerAndCreator      = sdkerrors.Register(ModuleName, 2109, "signer shoud equal to creator")

	ErrShardCompleted        = sdkerrors.Register(ModuleName, 2110, "shard already completed")
	ErrShardUnexpectedStatus = sdkerrors.Register(ModuleName, 2111, "invalid shard status")
	ErrInsufficientCoin      = sdkerrors.Register(ModuleName, 2112, "insufficient coin")
	ErrorInvalidProvider     = sdkerrors.Register(ModuleName, 2113, "invalid order provider")
	ErrorInvalidMetadata     = sdkerrors.Register(ModuleName, 2114, "invalid metadata")
	ErrorInvalidDid          = sdkerrors.Register(ModuleName, 2115, "invalid did")
	ErrorInvalidSignature    = sdkerrors.Register(ModuleName, 2116, "invalid signature")
	ErrorInvalidAddress      = sdkerrors.Register(ModuleName, 2117, "invalid address")
	ErrorInvalidDataId       = sdkerrors.Register(ModuleName, 2118, "invalid dataId")
	ErrorInvalidProposal     = sdkerrors.Register(ModuleName, 2119, "invalid proposal")
	ErrorInvalidOperation    = sdkerrors.Register(ModuleName, 2120, "invalid operation")
	ErrorNoPermission        = sdkerrors.Register(ModuleName, 2121, "invalid permission")
	ErrorInvalidShardSize    = sdkerrors.Register(ModuleName, 2122, "invalid shard size")
	ErrorOrderPledgeFailed   = sdkerrors.Register(ModuleName, 2123, "order pledge failed")
	ErrorInvalidOwner        = sdkerrors.Register(ModuleName, 2124, "invalid owner")
	ErrorInvalidDuration     = sdkerrors.Register(ModuleName, 2125, "invalid duration")
)
