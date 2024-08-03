package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/jack139/ganymede/ganymede/x/postoffice/types"
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
				PortId: types.PortID,
				PostList: []types.Post{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				PostCount: 2,
				SentPostList: []types.SentPost{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				SentPostCount: 2,
				TimedoutPostList: []types.TimedoutPost{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				TimedoutPostCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated post",
			genState: &types.GenesisState{
				PostList: []types.Post{
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
			desc: "invalid post count",
			genState: &types.GenesisState{
				PostList: []types.Post{
					{
						Id: 1,
					},
				},
				PostCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated sentPost",
			genState: &types.GenesisState{
				SentPostList: []types.SentPost{
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
			desc: "invalid sentPost count",
			genState: &types.GenesisState{
				SentPostList: []types.SentPost{
					{
						Id: 1,
					},
				},
				SentPostCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated timedoutPost",
			genState: &types.GenesisState{
				TimedoutPostList: []types.TimedoutPost{
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
			desc: "invalid timedoutPost count",
			genState: &types.GenesisState{
				TimedoutPostList: []types.TimedoutPost{
					{
						Id: 1,
					},
				},
				TimedoutPostCount: 0,
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
