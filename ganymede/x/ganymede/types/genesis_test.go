package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
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

				UsersList: []types.Users{
					{
						ChainAddr: "0",
					},
					{
						ChainAddr: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated users",
			genState: &types.GenesisState{
				UsersList: []types.Users{
					{
						ChainAddr: "0",
					},
					{
						ChainAddr: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
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
