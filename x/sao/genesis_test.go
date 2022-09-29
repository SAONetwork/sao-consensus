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

		OrderList: []types.Order{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		OrderCount: 2,
		ShardList: []types.Shard{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ShardCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SaoKeeper(t)
	sao.InitGenesis(ctx, *k, genesisState)
	got := sao.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.OrderList, got.OrderList)
	require.Equal(t, genesisState.OrderCount, got.OrderCount)
	require.ElementsMatch(t, genesisState.ShardList, got.ShardList)
	require.Equal(t, genesisState.ShardCount, got.ShardCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
