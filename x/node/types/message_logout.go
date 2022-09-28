package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLogout = "logout"

var _ sdk.Msg = &MsgLogout{}

func NewMsgLogout(creator string) *MsgLogout {
	return &MsgLogout{
		Creator: creator,
	}
}

func (msg *MsgLogout) Route() string {
	return RouterKey
}

func (msg *MsgLogout) Type() string {
	return TypeMsgLogout
}

func (msg *MsgLogout) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLogout) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLogout) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
