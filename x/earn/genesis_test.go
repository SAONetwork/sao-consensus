package earn_test

import (
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/earn"
	"github.com/SaoNetwork/sao/x/earn/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Pool: &types.Pool{},
		PledgeList: []types.Pledge{
			{
				Creator: "0",
			},
			{
				Creator: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EarnKeeper(t)
	earn.InitGenesis(ctx, *k, genesisState)
	got := earn.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Pool, got.Pool)
	require.ElementsMatch(t, genesisState.PledgeList, got.PledgeList)
	// this line is used by starport scaffolding # genesis/test/assert
}
