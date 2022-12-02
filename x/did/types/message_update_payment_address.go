package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdatePaymentAddress = "update_payment_address"

var _ sdk.Msg = &MsgUpdatePaymentAddress{}

func NewMsgUpdatePaymentAddress(creator string, accountId string) *MsgUpdatePaymentAddress {
	return &MsgUpdatePaymentAddress{
		Creator:   creator,
		AccountId: accountId,
	}
}

func (msg *MsgUpdatePaymentAddress) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePaymentAddress) Type() string {
	return TypeMsgUpdatePaymentAddress
}

func (msg *MsgUpdatePaymentAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePaymentAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePaymentAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
