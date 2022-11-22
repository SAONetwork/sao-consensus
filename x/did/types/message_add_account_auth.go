package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddAccountAuth = "add_account_auth"

var _ sdk.Msg = &MsgAddAccountAuth{}

func NewMsgAddAccountAuth(creator string, did string, accountAuth *AccountAuth) *MsgAddAccountAuth {
	return &MsgAddAccountAuth{
		Creator:     creator,
		Did:         did,
		AccountAuth: accountAuth,
	}
}

func (msg *MsgAddAccountAuth) Route() string {
	return RouterKey
}

func (msg *MsgAddAccountAuth) Type() string {
	return TypeMsgAddAccountAuth
}

func (msg *MsgAddAccountAuth) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddAccountAuth) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddAccountAuth) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
