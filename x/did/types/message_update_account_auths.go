package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateAccountAuths = "update_account_auths"

var _ sdk.Msg = &MsgUpdateAccountAuths{}

func NewMsgUpdateAccountAuths(creator string, did string, update []*AccountAuth, remove []string) *MsgUpdateAccountAuths {
	return &MsgUpdateAccountAuths{
		Creator: creator,
		Did:     did,
		Update:  update,
		Remove:  remove,
	}
}

func (msg *MsgUpdateAccountAuths) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAccountAuths) Type() string {
	return TypeMsgUpdateAccountAuths
}

func (msg *MsgUpdateAccountAuths) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAccountAuths) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAccountAuths) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
