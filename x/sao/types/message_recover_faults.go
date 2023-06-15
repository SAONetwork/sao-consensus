package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRecoverFaults = "recover_faults"

var _ sdk.Msg = &MsgRecoverFaults{}

func NewMsgRecoverFaults(creator string) *MsgRecoverFaults {
	return &MsgRecoverFaults{
		Creator: creator,
	}
}

func (msg *MsgRecoverFaults) Route() string {
	return RouterKey
}

func (msg *MsgRecoverFaults) Type() string {
	return TypeMsgRecoverFaults
}

func (msg *MsgRecoverFaults) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRecoverFaults) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRecoverFaults) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
