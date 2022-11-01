package model_test

import (
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/model"
	"github.com/SaoNetwork/sao/x/model/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		MetadataList: []types.Metadata{
			{
				DataId: "0",
			},
			{
				DataId: "1",
			},
		},
		ModelList: []types.Model{
			{
				Key: "0",
			},
			{
				Key: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ModelKeeper(t)
	model.InitGenesis(ctx, *k, genesisState)
	got := model.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MetadataList, got.MetadataList)
	require.ElementsMatch(t, genesisState.ModelList, got.ModelList)
	// this line is used by starport scaffolding # genesis/test/assert
}
