package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/did module sentinel errors
var (
	ErrDocExists     = sdkerrors.Register(ModuleName, 6101, "sid document exists")
	ErrAuthExists    = sdkerrors.Register(ModuleName, 6102, "account auth exists")
	ErrBindingExists = sdkerrors.Register(ModuleName, 6103, "account has already been bound")
	ErrSeedExists    = sdkerrors.Register(ModuleName, 6104, "seed exists")

	ErrVersionsNotFound    = sdkerrors.Register(ModuleName, 6201, "sid document version list not found")
	ErrAccountListNotFound = sdkerrors.Register(ModuleName, 6202, "account list not found")
	ErrBindingNotFound     = sdkerrors.Register(ModuleName, 6203, "binding proof not found")
	ErrSeedsNotFound       = sdkerrors.Register(ModuleName, 6204, "past seeds not found")
	ErrPayAddrNotSet       = sdkerrors.Register(ModuleName, 6205, "payment address not set yet")
	ErrAccountIdNotFound   = sdkerrors.Register(ModuleName, 6206, "account id not found")

	ErrUnbindPayAddr        = sdkerrors.Register(ModuleName, 6301, "cannot unbind payment account")
	ErrInvalidAccountId     = sdkerrors.Register(ModuleName, 6302, "cannot set an account with invalid chainId as payment account")
	ErrInvalidKeys          = sdkerrors.Register(ModuleName, 6303, "invalid keys")
	ErrInconsistentDocId    = sdkerrors.Register(ModuleName, 6304, "inconsistent document id")
	ErrInvalidBindingProof  = sdkerrors.Register(ModuleName, 6305, "invalid binding proof")
	ErrUnsupportedAccountId = sdkerrors.Register(ModuleName, 6306, "unsupported account id")
	ErrInconsistentDid      = sdkerrors.Register(ModuleName, 6307, "inconsistent did")
	ErrInvalidDid           = sdkerrors.Register(ModuleName, 6308, "invalid did")
	ErrUnsupportedDid       = sdkerrors.Register(ModuleName, 6309, "unsupported did")
	ErrUnhandledAccountDid  = sdkerrors.Register(ModuleName, 6310, "unhandled accountDid while update accountAuth")
	ErrInvalidCreator       = sdkerrors.Register(ModuleName, 6311, "invalid creator, creator should binding to current did")
	ErrNoNeedToUpdate       = sdkerrors.Register(ModuleName, 6312, "only need to update when remove bindings")
	ErrInvalidAuthCount     = sdkerrors.Register(ModuleName, 6313, "invalid account auth count")
	ErrUpdateAccAuthEmpty   = sdkerrors.Register(ModuleName, 6314, "update account auth is empty")
	ErrSamePayAddr          = sdkerrors.Register(ModuleName, 6315, "update the same address as the old one")
	ErrOutOfDate            = sdkerrors.Register(ModuleName, 6316, "out of date")
	ErrChangePayAddr        = sdkerrors.Register(ModuleName, 6317, "change payment address is not supported yet")
	ErrKidExist             = sdkerrors.Register(ModuleName, 6318, "creator has been bound to a kid")
)
