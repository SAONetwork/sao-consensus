package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/loan module sentinel errors
var (
	ErrCreditNotFound    = sdkerrors.Register(ModuleName, 7100, "credit not found")
	ErrLoanPoolNotFound  = sdkerrors.Register(ModuleName, 7101, "loan pool not found")
	ErrInvalidAmount     = sdkerrors.Register(ModuleName, 7102, "invalid amount")
	ErrNoEnoughAvailable = sdkerrors.Register(ModuleName, 7103, "loan pool available coin is not available")
)
