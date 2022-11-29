package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStore = "store"

var _ sdk.Msg = &MsgStore{}

func NewMsgStore(creator string, proposal *Proposal, signature *JwsSignature) *MsgStore {
	return &MsgStore{
		Creator:      creator,
		Proposal:     *proposal,
		JwsSignature: *signature,
	}
}

func (msg *MsgStore) Route() string {
	return RouterKey
}

func (msg *MsgStore) Type() string {
	return TypeMsgStore
}

func (msg *MsgStore) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStore) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStore) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
