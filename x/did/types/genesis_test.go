package types_test

import (
	"testing"

	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
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
				DidBalancesList: []types.DidBalances{
					{
						Did: "0",
					},
					{
						Did: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc:     "duplicated DidBindingProof",
			genState: &types.GenesisState{},
			valid:    false,
		},
		{
			desc: "duplicated accountList",
			genState: &types.GenesisState{
				AccountListList: []types.AccountList{
					{
						Did: "0",
					},
					{
						Did: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated accountAuth",
			genState: &types.GenesisState{
				AccountAuthList: []types.AccountAuth{
					{
						AccountDid: "0",
					},
					{
						AccountDid: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated sidDocument",
			genState: &types.GenesisState{
				SidDocumentList: []types.SidDocument{
					{
						VersionId: "0",
					},
					{
						VersionId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated sidDocumentVersion",
			genState: &types.GenesisState{
				SidDocumentVersionList: []types.SidDocumentVersion{
					{
						DocId: "0",
					},
					{
						DocId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated pastSeeds",
			genState: &types.GenesisState{
				PastSeedsList: []types.PastSeeds{
					{
						Did: "0",
					},
					{
						Did: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated paymentAddress",
			genState: &types.GenesisState{
				PaymentAddressList: []types.PaymentAddress{
					{
						Did: "0",
					},
					{
						Did: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated accountId",
			genState: &types.GenesisState{
				AccountIdList: []types.AccountId{
					{
						AccountDid: "0",
					},
					{
						AccountDid: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated did",
			genState: &types.GenesisState{
				DidList: []types.Did{
					{
						AccountId: "0",
					},
					{
						AccountId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated kid",
			genState: &types.GenesisState{
				KidList: []types.Kid{
					{
						AccountId: "0",
					},
					{
						AccountId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated didBalances",
			genState: &types.GenesisState{
				DidBalancesList: []types.DidBalances{
					{
						Did: "0",
					},
					{
						Did: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
