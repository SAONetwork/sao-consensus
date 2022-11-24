package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/market module sentinel errors
var (
	ErrInvalidAmount = sdkerrors.Register(ModuleName, 6100, "invalid amount")
)
