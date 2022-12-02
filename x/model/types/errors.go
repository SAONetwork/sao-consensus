package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/model module sentinel errors
var (
	ErrDataIdExists     = sdkerrors.Register(ModuleName, 4100, "dataId already exists")
	ErrInvalidDataId    = sdkerrors.Register(ModuleName, 4101, "invali dataId")
	ErrModelExists      = sdkerrors.Register(ModuleName, 4102, "model already exists")
	ErrOnlyOwner        = sdkerrors.Register(ModuleName, 4103, "owner only action")
	ErrInsufficientCoin = sdkerrors.Register(ModuleName, 4104, "insufficient coin")
	ErrInvalidOperation = sdkerrors.Register(ModuleName, 4105, "invalid operation")
)
