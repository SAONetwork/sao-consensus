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

		DidBindingProofsList: []types.DidBindingProofs{
			{
				AccountId: "0",
			},
			{
				AccountId: "1",
			},
		},
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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DidKeeper(t)
	did.InitGenesis(ctx, *k, genesisState)
	got := did.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DidBindingProofsList, got.DidBindingProofsList)
	require.ElementsMatch(t, genesisState.AccountListList, got.AccountListList)
	require.ElementsMatch(t, genesisState.AccountAuthList, got.AccountAuthList)
	// this line is used by starport scaffolding # genesis/test/assert
}
