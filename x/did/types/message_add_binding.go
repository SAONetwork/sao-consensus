package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddBinding = "add_binding"

var _ sdk.Msg = &MsgAddBinding{}

func NewMsgAddBinding(creator string, accountId string, proof *BindingProof) *MsgAddBinding {
	return &MsgAddBinding{
		Creator:   creator,
		AccountId: accountId,
		Proof:     proof,
	}
}

func (msg *MsgAddBinding) Route() string {
	return RouterKey
}

func (msg *MsgAddBinding) Type() string {
	return TypeMsgAddBinding
}

func (msg *MsgAddBinding) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddBinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddBinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
