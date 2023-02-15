package node_test

import (
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/node"
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		NodeList: []types.Node{
			{
				Creator: "0",
			},
			{
				Creator: "1",
			},
		},
		ShardList: []types.Shard{
			{
				Idx: "0",
			},
			{
				Idx: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NodeKeeper(t)
	node.InitGenesis(ctx, *k, genesisState)
	got := node.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.NodeList, got.NodeList)
	require.ElementsMatch(t, genesisState.ShardList, got.ShardList)
	// this line is used by starport scaffolding # genesis/test/assert
}
