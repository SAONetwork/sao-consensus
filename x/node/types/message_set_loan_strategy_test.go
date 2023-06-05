package types

import (
	"testing"

	"github.com/SaoNetwork/sao/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSetLoanStrategy_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSetLoanStrategy
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSetLoanStrategy{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSetLoanStrategy{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
