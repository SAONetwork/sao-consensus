package did_test

import (
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/did"
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AccountListList: []types.AccountList{
			{
				Did: "0",
			},
			{
				Did: "1",
			},
		},
		AccountAuthList: []types.AccountAuth{
			{
				AccountDid: "0",
			},
			{
				AccountDid: "1",
			},
		},
		SidDocumentList: []types.SidDocument{
			{
				VersionId: "0",
			},
			{
				VersionId: "1",
			},
		},
		SidDocumentVersionList: []types.SidDocumentVersion{
			{
				DocId: "0",
			},
			{
				DocId: "1",
			},
		},
		PastSeedsList: []types.PastSeeds{
			{
				Did: "0",
			},
			{
				Did: "1",
			},
		},
		PaymentAddressList: []types.PaymentAddress{
			{
				Did: "0",
			},
			{
				Did: "1",
			},
		},
		AccountIdList: []types.AccountId{
			{
				AccountDid: "0",
			},
			{
				AccountDid: "1",
			},
		},
		DidList: []types.Did{
			{
				AccountId: "0",
			},
			{
				AccountId: "1",
			},
		},
		KidList: []types.Kid{
			{
				AccountId: "0",
			},
			{
				AccountId: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DidKeeper(t)
	did.InitGenesis(ctx, *k, genesisState)
	got := did.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AccountListList, got.AccountListList)
	require.ElementsMatch(t, genesisState.AccountAuthList, got.AccountAuthList)
	require.ElementsMatch(t, genesisState.SidDocumentList, got.SidDocumentList)
	require.ElementsMatch(t, genesisState.SidDocumentVersionList, got.SidDocumentVersionList)
	require.ElementsMatch(t, genesisState.PastSeedsList, got.PastSeedsList)
	require.ElementsMatch(t, genesisState.PaymentAddressList, got.PaymentAddressList)
	require.ElementsMatch(t, genesisState.AccountIdList, got.AccountIdList)
	require.ElementsMatch(t, genesisState.DidList, got.DidList)
	require.ElementsMatch(t, genesisState.KidList, got.KidList)
	// this line is used by starport scaffolding # genesis/test/assert
}
