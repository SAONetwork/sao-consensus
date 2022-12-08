package types_test

import (
	"testing"

	"github.com/SaoNetwork/sao/x/model/types"
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated metadata",
			genState: &types.GenesisState{
				MetadataList: []types.Metadata{
					{
						DataId: "0",
					},
					{
						DataId: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated model",
			genState: &types.GenesisState{
				ModelList: []types.Model{
					{
						Key: "0",
					},
					{
						Key: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated expiredData",
			genState: &types.GenesisState{
				ExpiredDataList: []types.ExpiredData{
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
