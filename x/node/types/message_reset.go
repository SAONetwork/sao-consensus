package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReset = "reset"

var _ sdk.Msg = &MsgReset{}

func NewMsgReset(creator string, peer string) *MsgReset {
	return &MsgReset{
		Creator: creator,
		Peer:    peer,
	}
}

func (msg *MsgReset) Route() string {
	return RouterKey
}

func (msg *MsgReset) Type() string {
	return TypeMsgReset
}

func (msg *MsgReset) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
