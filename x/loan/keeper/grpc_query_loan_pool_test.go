package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/loan/types"
)

func TestLoanPoolQuery(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestLoanPool(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLoanPoolRequest
		response *types.QueryGetLoanPoolResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLoanPoolRequest{},
			response: &types.QueryGetLoanPoolResponse{LoanPool: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.LoanPool(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
