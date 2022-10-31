package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReady = "ready"

var _ sdk.Msg = &MsgReady{}

func NewMsgReady(creator string, orderId uint64) *MsgReady {
	return &MsgReady{
		Creator: creator,
		OrderId: orderId,
	}
}

func (msg *MsgReady) Route() string {
	return RouterKey
}

func (msg *MsgReady) Type() string {
	return TypeMsgReady
}

func (msg *MsgReady) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReady) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReady) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
