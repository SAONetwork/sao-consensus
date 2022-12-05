package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdataPermission = "updata_permission"

var _ sdk.Msg = &MsgUpdataPermission{}

func NewMsgUpdataPermission(creator string, proposal PermissionProposal, signature JwsSignature) *MsgUpdataPermission {
	return &MsgUpdataPermission{
		Creator:      creator,
		Proposal:     proposal,
		JwsSignature: signature,
	}
}

func (msg *MsgUpdataPermission) Route() string {
	return RouterKey
}

func (msg *MsgUpdataPermission) Type() string {
	return TypeMsgUpdataPermission
}

func (msg *MsgUpdataPermission) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdataPermission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdataPermission) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
