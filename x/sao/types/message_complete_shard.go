package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCompleteShard = "complete_shard"

var _ sdk.Msg = &MsgCompleteShard{}

func NewMsgCompleteShard(creator string, idx string, cid string, size int32) *MsgCompleteShard {
	return &MsgCompleteShard{
		Creator: creator,
		Idx:     idx,
		Cid:     cid,
		Size_:   size,
	}
}

func (msg *MsgCompleteShard) Route() string {
	return RouterKey
}

func (msg *MsgCompleteShard) Type() string {
	return TypeMsgCompleteShard
}

func (msg *MsgCompleteShard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCompleteShard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCompleteShard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
