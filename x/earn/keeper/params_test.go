package keeper_test

import (
	"testing"

	testkeeper "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/x/earn/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EarnKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.BlockReward, k.BlockReward(ctx))
	require.EqualValues(t, params.EarnDenom, k.EarnDenom(ctx))
}