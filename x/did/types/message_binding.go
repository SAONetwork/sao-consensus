package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBinding = "binding"

var _ sdk.Msg = &MsgBinding{}

func NewMsgBinding(creator string, accountId string, rootDocId string, keys []*PubKey, accountAuth *AccountAuth, proof *BindingProof) *MsgBinding {
	return &MsgBinding{
		Creator:     creator,
		AccountId:   accountId,
		RootDocId:   rootDocId,
		Keys:        keys,
		AccountAuth: accountAuth,
		Proof:       proof,
	}
}

func (msg *MsgBinding) Route() string {
	return RouterKey
}

func (msg *MsgBinding) Type() string {
	return TypeMsgBinding
}

func (msg *MsgBinding) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
