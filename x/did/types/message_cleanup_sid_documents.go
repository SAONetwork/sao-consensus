package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCleanupSidDocuments = "cleanup_sid_documents"

var _ sdk.Msg = &MsgCleanupSidDocuments{}

func NewMsgCleanupSidDocuments(creator string, rootDocId string) *MsgCleanupSidDocuments {
	return &MsgCleanupSidDocuments{
		Creator:   creator,
		RootDocId: rootDocId,
	}
}

func (msg *MsgCleanupSidDocuments) Route() string {
	return RouterKey
}

func (msg *MsgCleanupSidDocuments) Type() string {
	return TypeMsgCleanupSidDocuments
}

func (msg *MsgCleanupSidDocuments) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCleanupSidDocuments) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCleanupSidDocuments) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
