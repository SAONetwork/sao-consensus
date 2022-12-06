package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/did module sentinel errors
var (
	ErrDocExists     = sdkerrors.Register(ModuleName, 6101, "sid document exists")
	ErrAuthExists    = sdkerrors.Register(ModuleName, 6102, "account auth exists")
	ErrBindingExists = sdkerrors.Register(ModuleName, 6103, "binding proof exists")
	ErrSeedExists    = sdkerrors.Register(ModuleName, 6104, "seed exists")

	ErrVersionsNotFound    = sdkerrors.Register(ModuleName, 6201, "sid document version list not found")
	ErrAccountListNotFound = sdkerrors.Register(ModuleName, 6202, "account list not found")
	ErrBindingNotFound     = sdkerrors.Register(ModuleName, 6203, "binding proof not found")
	ErrSeedsNotFound       = sdkerrors.Register(ModuleName, 6204, "past seeds not found")
	ErrPayAddrNotSet       = sdkerrors.Register(ModuleName, 6205, "payment address not set yet")

	ErrUnbindPayAddr        = sdkerrors.Register(ModuleName, 6301, "cannot unbind payment account")
	ErrInvalidAccountId     = sdkerrors.Register(ModuleName, 6302, "cannot set an account with invalid chainId as payment account")
	ErrDocInvalidKeys       = sdkerrors.Register(ModuleName, 6303, "invalid keys")
	ErrInconsistentDocId    = sdkerrors.Register(ModuleName, 6304, "inconsistent document id")
	ErrInvalidBindingProof  = sdkerrors.Register(ModuleName, 6305, "invalid binding proof")
	ErrUnsupportedAccountId = sdkerrors.Register(ModuleName, 6306, "unsupported account id")
)
