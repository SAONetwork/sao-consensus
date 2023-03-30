package sao_test

import (
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/sao"
	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		TimeoutOrderList: []types.TimeoutOrder{
			{
				Height: 0,
			},
			{
				Height: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SaoKeeper(t)
	sao.InitGenesis(ctx, *k, genesisState)
	got := sao.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
	require.ElementsMatch(t, genesisState.TimeoutOrderList, got.TimeoutOrderList)
	// this line is used by starport scaffolding # genesis/test/assert
}
