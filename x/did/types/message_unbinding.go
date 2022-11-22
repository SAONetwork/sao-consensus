package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUnbinding = "unbinding"

var _ sdk.Msg = &MsgUnbinding{}

func NewMsgUnbinding(creator string, accountId string) *MsgUnbinding {
	return &MsgUnbinding{
		Creator:   creator,
		AccountId: accountId,
	}
}

func (msg *MsgUnbinding) Route() string {
	return RouterKey
}

func (msg *MsgUnbinding) Type() string {
	return TypeMsgUnbinding
}

func (msg *MsgUnbinding) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUnbinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnbinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
