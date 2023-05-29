package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgReportFaults = "report_faults"

var _ sdk.Msg = &MsgReportFaults{}

func NewMsgReportFaults(creator string) *MsgReportFaults {
	return &MsgReportFaults{
		Creator: creator,
	}
}

func (msg *MsgReportFaults) Route() string {
	return RouterKey
}

func (msg *MsgReportFaults) Type() string {
	return TypeMsgReportFaults
}

func (msg *MsgReportFaults) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgReportFaults) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgReportFaults) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
