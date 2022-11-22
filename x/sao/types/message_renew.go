package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRenew = "renew"

var _ sdk.Msg = &MsgRenew{}

func NewMsgRenew(creator string, data []string, duration int32, timeout int32) *MsgRenew {
	return &MsgRenew{
		Creator:  creator,
		Data:     data,
		Duration: duration,
		Timeout:  timeout,
	}
}

func (msg *MsgRenew) Route() string {
	return RouterKey
}

func (msg *MsgRenew) Type() string {
	return TypeMsgRenew
}

func (msg *MsgRenew) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRenew) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRenew) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
