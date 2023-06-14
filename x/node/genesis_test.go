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
		PledgeDebtList: []types.PledgeDebt{
			{
				Sp: "0",
			},
			{
				Sp: "1",
			},
		},
		VstorageList: []types.Vstorage{
			{
				Sp: "0",
			},
			{
				Sp: "1",
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
	require.ElementsMatch(t, genesisState.PledgeDebtList, got.PledgeDebtList)
	require.ElementsMatch(t, genesisState.VstorageList, got.VstorageList)
	// this line is used by starport scaffolding # genesis/test/assert
}
