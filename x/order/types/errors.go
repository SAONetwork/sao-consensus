package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/order module sentinel errors
var (
	ErrInvalidReplica = sdkerrors.Register(ModuleName, 5100, "replica support 0 ~ 5")
)
