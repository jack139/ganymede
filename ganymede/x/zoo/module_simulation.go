package zoo

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jack139/ganymede/ganymede/testutil/sample"
	zoosimulation "github.com/jack139/ganymede/ganymede/x/zoo/simulation"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = zoosimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateKvzoo = "op_weight_msg_kvzoo"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateKvzoo int = 100

	opWeightMsgUpdateKvzoo = "op_weight_msg_kvzoo"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateKvzoo int = 100

	opWeightMsgDeleteKvzoo = "op_weight_msg_kvzoo"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteKvzoo int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	zooGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		KvzooList: []types.Kvzoo{
			{
				Creator: sample.AccAddress(),
				Owner:   "0",
				ZooKey:  "0",
			},
			{
				Creator: sample.AccAddress(),
				Owner:   "1",
				ZooKey:  "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&zooGenesis)
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

	var weightMsgCreateKvzoo int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateKvzoo, &weightMsgCreateKvzoo, nil,
		func(_ *rand.Rand) {
			weightMsgCreateKvzoo = defaultWeightMsgCreateKvzoo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateKvzoo,
		zoosimulation.SimulateMsgCreateKvzoo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateKvzoo int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateKvzoo, &weightMsgUpdateKvzoo, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateKvzoo = defaultWeightMsgUpdateKvzoo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateKvzoo,
		zoosimulation.SimulateMsgUpdateKvzoo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteKvzoo int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteKvzoo, &weightMsgDeleteKvzoo, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteKvzoo = defaultWeightMsgDeleteKvzoo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteKvzoo,
		zoosimulation.SimulateMsgDeleteKvzoo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateKvzoo,
			defaultWeightMsgCreateKvzoo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				zoosimulation.SimulateMsgCreateKvzoo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateKvzoo,
			defaultWeightMsgUpdateKvzoo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				zoosimulation.SimulateMsgUpdateKvzoo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteKvzoo,
			defaultWeightMsgDeleteKvzoo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				zoosimulation.SimulateMsgDeleteKvzoo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
