package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveVstorage = "remove_vstorage"

var _ sdk.Msg = &MsgRemoveVstorage{}

func NewMsgRemoveVstorage(creator string, size uint64) *MsgRemoveVstorage {
	return &MsgRemoveVstorage{
		Creator: creator,
		Size_:   size,
	}
}

func (msg *MsgRemoveVstorage) Route() string {
	return RouterKey
}

func (msg *MsgRemoveVstorage) Type() string {
	return TypeMsgRemoveVstorage
}

func (msg *MsgRemoveVstorage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveVstorage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveVstorage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
