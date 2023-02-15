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
				ShardList: []types.Shard{
					{
						Idx: "0",
					},
					{
						Idx: "1",
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
			desc: "duplicated shard",
			genState: &types.GenesisState{
				ShardList: []types.Shard{
					{
						Idx: "0",
					},
					{
						Idx: "0",
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
