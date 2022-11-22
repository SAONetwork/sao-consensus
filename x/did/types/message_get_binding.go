package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGetBinding = "get_binding"

var _ sdk.Msg = &MsgGetBinding{}

func NewMsgGetBinding(creator string, accountId string) *MsgGetBinding {
	return &MsgGetBinding{
		Creator:   creator,
		AccountId: accountId,
	}
}

func (msg *MsgGetBinding) Route() string {
	return RouterKey
}

func (msg *MsgGetBinding) Type() string {
	return TypeMsgGetBinding
}

func (msg *MsgGetBinding) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGetBinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGetBinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
