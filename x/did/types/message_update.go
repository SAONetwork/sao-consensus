package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdate = "Update"

var _ sdk.Msg = &MsgUpdate{}

func NewMsgUpdate(creator string, did string, newDocId string, keys []*PubKey, timestamp uint64, updateAccountAuth []*AccountAuth, removeAccountDid []string, pastSeed string) *MsgUpdate {
	return &MsgUpdate{
		Creator:           creator,
		Did:               did,
		NewDocId:          newDocId,
		Keys:              keys,
		Timestamp:         timestamp,
		UpdateAccountAuth: updateAccountAuth,
		RemoveAccountDid:  removeAccountDid,
		PastSeed:          pastSeed,
	}
}

func (msg *MsgUpdate) Route() string {
	return RouterKey
}

func (msg *MsgUpdate) Type() string {
	return TypeMsgUpdate
}

func (msg *MsgUpdate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
