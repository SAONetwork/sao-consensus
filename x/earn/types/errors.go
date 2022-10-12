package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/earn module sentinel errors
var (
	ErrPoolNotFound     = sdkerrors.Register(ModuleName, 3100, "pool not found")
	ErrPledgeNotFound   = sdkerrors.Register(ModuleName, 3101, "not pledged yet")
	ErrDenom            = sdkerrors.Register(ModuleName, 3102, "invalid denom")
	ErrInsufficientCoin = sdkerrors.Register(ModuleName, 3103, "insufficient coin")
)
