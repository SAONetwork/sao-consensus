package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRenew = "renew"

var _ sdk.Msg = &MsgRenew{}

func NewMsgRenew(creator string, proposal *RenewProposal, signature *JwsSignature) *MsgRenew {
	return &MsgRenew{
		Creator:      creator,
		Proposal:     *proposal,
		JwsSignature: *signature,
	}
}

func (msg *MsgRenew) Route() string {
	return RouterKey
}

func (msg *MsgRenew) Type() string {
	return TypeMsgRenew
}

func (msg *MsgRenew) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRenew) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRenew) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
