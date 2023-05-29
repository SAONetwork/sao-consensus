package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReportFailures = "report_failures"

var _ sdk.Msg = &MsgReportFailures{}

func NewMsgReportFailures(creator string) *MsgReportFailures {
	return &MsgReportFailures{
		Creator: creator,
	}
}

func (msg *MsgReportFailures) Route() string {
	return RouterKey
}

func (msg *MsgReportFailures) Type() string {
	return TypeMsgReportFailures
}

func (msg *MsgReportFailures) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReportFailures) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReportFailures) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
