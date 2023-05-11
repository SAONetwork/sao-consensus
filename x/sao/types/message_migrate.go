package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMigrate = "migrate"

var _ sdk.Msg = &MsgMigrate{}

func NewMsgMigrate(creator string, data []string, provider string) *MsgMigrate {
	return &MsgMigrate{
		Creator:  creator,
		Data:     data,
		Provider: provider,
	}
}

func (msg *MsgMigrate) Route() string {
	return RouterKey
}

func (msg *MsgMigrate) Type() string {
	return TypeMsgMigrate
}

func (msg *MsgMigrate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMigrate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMigrate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
