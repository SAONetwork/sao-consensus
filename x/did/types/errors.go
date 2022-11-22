package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/did module sentinel errors
var (
	ErrSample        = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrProofNotFound = sdkerrors.Register(ModuleName, 1101, "binding proof not found")
)
