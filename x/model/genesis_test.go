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
		ExpiredDataList: []types.ExpiredData{
			{
				Height: 0,
			},
			{
				Height: 1,
			},
		},
		OrderFinishList: []types.OrderFinish{
			{
				Height: 0,
			},
			{
				Height: 1,
			},
		},
		ExpiredOrderList: []types.ExpiredOrder{
			{
				Height: 0,
			},
			{
				Height: 1,
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
	require.ElementsMatch(t, genesisState.ExpiredDataList, got.ExpiredDataList)
	require.ElementsMatch(t, genesisState.OrderFinishList, got.OrderFinishList)
	require.ElementsMatch(t, genesisState.ExpiredOrderList, got.ExpiredOrderList)
	// this line is used by starport scaffolding # genesis/test/assert
}
