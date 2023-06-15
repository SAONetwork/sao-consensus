package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddVstorage = "add_vstorage"

var _ sdk.Msg = &MsgAddVstorage{}

func NewMsgAddVstorage(creator string, size uint64) *MsgAddVstorage {
	return &MsgAddVstorage{
		Creator: creator,
		Size_:   size,
	}
}

func (msg *MsgAddVstorage) Route() string {
	return RouterKey
}

func (msg *MsgAddVstorage) Type() string {
	return TypeMsgAddVstorage
}

func (msg *MsgAddVstorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddVstorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddVstorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
