package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/did module sentinel errors
var (
	ErrBindingExists = sdkerrors.Register(ModuleName, 6101, "binding proof exists")
	ErrAuthExists    = sdkerrors.Register(ModuleName, 6102, "account auth exists")
	ErrDocExists     = sdkerrors.Register(ModuleName, 6103, "sid document exists")
	ErrSeedExists    = sdkerrors.Register(ModuleName, 6104, "seed exists")

	ErrAccountListNotFound = sdkerrors.Register(ModuleName, 6201, "account list not found")
	ErrVersionsNotFound    = sdkerrors.Register(ModuleName, 6202, "sid document version list not found")
	ErrSeedsNotFound       = sdkerrors.Register(ModuleName, 6202, "past seeds not found")
)
