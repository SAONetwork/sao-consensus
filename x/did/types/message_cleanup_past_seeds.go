package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCleanupPastSeeds = "cleanup_past_seeds"

var _ sdk.Msg = &MsgCleanupPastSeeds{}

func NewMsgCleanupPastSeeds(creator string, did string) *MsgCleanupPastSeeds {
	return &MsgCleanupPastSeeds{
		Creator: creator,
		Did:     did,
	}
}

func (msg *MsgCleanupPastSeeds) Route() string {
	return RouterKey
}

func (msg *MsgCleanupPastSeeds) Type() string {
	return TypeMsgCleanupPastSeeds
}

func (msg *MsgCleanupPastSeeds) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCleanupPastSeeds) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCleanupPastSeeds) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
