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
)
