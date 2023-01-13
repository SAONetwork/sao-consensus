package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/node module sentinel errors
var (
	ErrAlreadyRegistered = sdkerrors.Register(ModuleName, 1100, "already registered")
	ErrNodeNotFound      = sdkerrors.Register(ModuleName, 1101, "node not found")
	ErrOnlyOwner         = sdkerrors.Register(ModuleName, 1102, "only node owner can execute this action")
	ErrInvalidPeer       = sdkerrors.Register(ModuleName, 1103, "invalid peer")
	ErrNotValidator      = sdkerrors.Register(ModuleName, 1104, "node should staking as a validator first")
	ErrInvalidStatus     = sdkerrors.Register(ModuleName, 1105, "invalid status")
	ErrSignerAndCreator  = sdkerrors.Register(ModuleName, 2109, "signer shoud equal to creator")

	ErrPoolNotFound     = sdkerrors.Register(ModuleName, 3100, "pool not found")
	ErrPledgeNotFound   = sdkerrors.Register(ModuleName, 3101, "not pledged yet")
	ErrDenom            = sdkerrors.Register(ModuleName, 3102, "invalid denom")
	ErrInsufficientCoin = sdkerrors.Register(ModuleName, 3103, "insufficient coin")
	ErrInvalidCommitId  = sdkerrors.Register(ModuleName, 3104, "invalid commit")
	ErrInvalidLastOrder = sdkerrors.Register(ModuleName, 3105, "invalid last order id")
)
