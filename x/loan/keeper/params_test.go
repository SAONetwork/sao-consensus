package keeper_test

import (
	"testing"

	testkeeper "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/x/loan/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LoanKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.InterestRatePerBlock, k.LoanInterest(ctx))
	require.EqualValues(t, params.MinLiquidityRatio, k.MinLiquidityRatio(ctx))
}
