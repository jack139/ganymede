package ganymede

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jack139/ganymede/ganymede/testutil/sample"
	ganymedesimulation "github.com/jack139/ganymede/ganymede/x/ganymede/simulation"
	"github.com/jack139/ganymede/ganymede/x/ganymede/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = ganymedesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateUsers = "op_weight_msg_users"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUsers int = 100

	opWeightMsgUpdateUsers = "op_weight_msg_users"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateUsers int = 100

	opWeightMsgDeleteUsers = "op_weight_msg_users"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteUsers int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	ganymedeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		UsersList: []types.Users{
			{
				Creator:   sample.AccAddress(),
				ChainAddr: "0",
			},
			{
				Creator:   sample.AccAddress(),
				ChainAddr: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&ganymedeGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateUsers int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateUsers, &weightMsgCreateUsers, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUsers = defaultWeightMsgCreateUsers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUsers,
		ganymedesimulation.SimulateMsgCreateUsers(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateUsers int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateUsers, &weightMsgUpdateUsers, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUsers = defaultWeightMsgUpdateUsers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUsers,
		ganymedesimulation.SimulateMsgUpdateUsers(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteUsers int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteUsers, &weightMsgDeleteUsers, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUsers = defaultWeightMsgDeleteUsers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUsers,
		ganymedesimulation.SimulateMsgDeleteUsers(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateUsers,
			defaultWeightMsgCreateUsers,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ganymedesimulation.SimulateMsgCreateUsers(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateUsers,
			defaultWeightMsgUpdateUsers,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ganymedesimulation.SimulateMsgUpdateUsers(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteUsers,
			defaultWeightMsgDeleteUsers,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ganymedesimulation.SimulateMsgDeleteUsers(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
