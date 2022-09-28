package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLogin = "login"

var _ sdk.Msg = &MsgLogin{}

func NewMsgLogin(creator string, peer string) *MsgLogin {
	return &MsgLogin{
		Creator: creator,
		Peer:    peer,
	}
}

func (msg *MsgLogin) Route() string {
	return RouterKey
}

func (msg *MsgLogin) Type() string {
	return TypeMsgLogin
}

func (msg *MsgLogin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLogin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLogin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
