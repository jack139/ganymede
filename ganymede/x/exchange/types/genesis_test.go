package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
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

				AskList: []types.Ask{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				AskCount: 2,
				ReplyList: []types.Reply{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ReplyCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated ask",
			genState: &types.GenesisState{
				AskList: []types.Ask{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid ask count",
			genState: &types.GenesisState{
				AskList: []types.Ask{
					{
						Id: 1,
					},
				},
				AskCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated reply",
			genState: &types.GenesisState{
				ReplyList: []types.Reply{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid reply count",
			genState: &types.GenesisState{
				ReplyList: []types.Reply{
					{
						Id: 1,
					},
				},
				ReplyCount: 0,
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
