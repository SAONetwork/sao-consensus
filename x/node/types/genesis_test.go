package types_test

import (
	"testing"

	"github.com/SaoNetwork/sao/x/node/types"
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated node",
			genState: &types.GenesisState{
				NodeList: []types.Node{
					{
						Creator: "0",
					},
					{
						Creator: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated pledgeDebt",
			genState: &types.GenesisState{
				PledgeDebtList: []types.PledgeDebt{
					{
						Sp: "0",
					},
					{
						Sp: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated vstorage",
			genState: &types.GenesisState{
				VstorageList: []types.Vstorage{
					{
						Sp: "0",
					},
					{
						Sp: "0",
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
