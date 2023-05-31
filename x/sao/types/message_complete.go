package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgComplete = "complete"

var _ sdk.Msg = &MsgComplete{}

func NewMsgComplete(creator string, orderId uint64, cid string, size uint64, provider string) *MsgComplete {
	return &MsgComplete{
		Creator:  creator,
		OrderId:  orderId,
		Cid:      cid,
		Size_:    size,
		Provider: provider,
	}
}

func (msg *MsgComplete) Route() string {
	return RouterKey
}

func (msg *MsgComplete) Type() string {
	return TypeMsgComplete
}

func (msg *MsgComplete) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgComplete) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgComplete) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
