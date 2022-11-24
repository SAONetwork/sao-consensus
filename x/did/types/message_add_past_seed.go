package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddPastSeed = "add_past_seed"

var _ sdk.Msg = &MsgAddPastSeed{}

func NewMsgAddPastSeed(creator string, did string, pastSeed string) *MsgAddPastSeed {
	return &MsgAddPastSeed{
		Creator:  creator,
		Did:      did,
		PastSeed: pastSeed,
	}
}

func (msg *MsgAddPastSeed) Route() string {
	return RouterKey
}

func (msg *MsgAddPastSeed) Type() string {
	return TypeMsgAddPastSeed
}

func (msg *MsgAddPastSeed) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddPastSeed) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddPastSeed) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
