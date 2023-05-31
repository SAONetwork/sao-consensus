package types_test

import (
	"testing"

	"github.com/SaoNetwork/sao/x/sao/types"
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
				TimeoutOrderList: []types.TimeoutOrder{
					{
						Height: 0,
					},
					{
						Height: 1,
					},
				},
				ExpiredShardList: []types.ExpiredShard{
					{
						Height: 0,
					},
					{
						Height: 1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc:     "duplicated order",
			genState: &types.GenesisState{},
			valid:    false,
		},
		{
			desc:     "invalid order count",
			genState: &types.GenesisState{},
			valid:    false,
		},
		{
			desc:     "duplicated shard",
			genState: &types.GenesisState{},
			valid:    false,
		},
		{
			desc:     "invalid shard count",
			genState: &types.GenesisState{},
			valid:    false,
		},
		{
			desc: "duplicated timeoutOrder",
			genState: &types.GenesisState{
				TimeoutOrderList: []types.TimeoutOrder{
					{
						Height: 0,
					},
					{
						Height: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated expiredShard",
			genState: &types.GenesisState{
				ExpiredShardList: []types.ExpiredShard{
					{
						Height: 0,
					},
					{
						Height: 0,
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
