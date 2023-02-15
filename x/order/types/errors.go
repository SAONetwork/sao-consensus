package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/order module sentinel errors
var (
	ErrInvalidReplica        = sdkerrors.Register(ModuleName, 5100, "replica support 0 ~ 5")
	ErrOrderUnexpectedStatus = sdkerrors.Register(ModuleName, 5101, "invalid order status")
	ErrInsufficientCoin      = sdkerrors.Register(ModuleName, 5102, "insufficient coin")
	ErrorOrderPayment        = sdkerrors.Register(ModuleName, 5103, "failed to pay order")
	ErrorRefundOrder         = sdkerrors.Register(ModuleName, 5104, "failed to refund order")
)
