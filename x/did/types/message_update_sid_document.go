package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateSidDocument = "update_sid_document"

var _ sdk.Msg = &MsgUpdateSidDocument{}

func NewMsgUpdateSidDocument(creator string, signingKey string, encryptKey string, rootDocId string) *MsgUpdateSidDocument {
	return &MsgUpdateSidDocument{
		Creator:    creator,
		SigningKey: signingKey,
		EncryptKey: encryptKey,
		RootDocId:  rootDocId,
	}
}

func (msg *MsgUpdateSidDocument) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSidDocument) Type() string {
	return TypeMsgUpdateSidDocument
}

func (msg *MsgUpdateSidDocument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSidDocument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSidDocument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
