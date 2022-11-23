package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/did module sentinel errors
var (
	ErrBindingExists       = sdkerrors.Register(ModuleName, 6101, "binding proof exists")
	ErrAuthExists          = sdkerrors.Register(ModuleName, 6102, "account auth exists")
	ErrAccountListNotFound = sdkerrors.Register(ModuleName, 6103, "account list not found")
	ErrDocExists           = sdkerrors.Register(ModuleName, 6104, "sid document exists")
	ErrVersionsNotFound    = sdkerrors.Register(ModuleName, 6105, "sid document version list not found")
)
