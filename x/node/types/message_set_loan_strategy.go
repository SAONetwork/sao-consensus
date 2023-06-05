package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetLoanStrategy = "set_loan_strategy"

var _ sdk.Msg = &MsgSetLoanStrategy{}

func NewMsgSetLoanStrategy(creator string, loanStrategy int32) *MsgSetLoanStrategy {
	return &MsgSetLoanStrategy{
		Creator:      creator,
		LoanStrategy: loanStrategy,
	}
}

func (msg *MsgSetLoanStrategy) Route() string {
	return RouterKey
}

func (msg *MsgSetLoanStrategy) Type() string {
	return TypeMsgSetLoanStrategy
}

func (msg *MsgSetLoanStrategy) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetLoanStrategy) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetLoanStrategy) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
