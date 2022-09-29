package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTerminate = "terminate"

var _ sdk.Msg = &MsgTerminate{}

func NewMsgTerminate(creator string, orderId uint64) *MsgTerminate {
	return &MsgTerminate{
		Creator: creator,
		OrderId: orderId,
	}
}

func (msg *MsgTerminate) Route() string {
	return RouterKey
}

func (msg *MsgTerminate) Type() string {
	return TypeMsgTerminate
}

func (msg *MsgTerminate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTerminate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTerminate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
