package exchange

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jack139/ganymede/ganymede/testutil/sample"
	exchangesimulation "github.com/jack139/ganymede/ganymede/x/exchange/simulation"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = exchangesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgNewAsk = "op_weight_msg_new_ask"
	// TODO: Determine the simulation weight value
	defaultWeightMsgNewAsk int = 100

	opWeightMsgNewReply = "op_weight_msg_new_reply"
	// TODO: Determine the simulation weight value
	defaultWeightMsgNewReply int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	exchangeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&exchangeGenesis)
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

	var weightMsgNewAsk int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgNewAsk, &weightMsgNewAsk, nil,
		func(_ *rand.Rand) {
			weightMsgNewAsk = defaultWeightMsgNewAsk
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgNewAsk,
		exchangesimulation.SimulateMsgNewAsk(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgNewReply int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgNewReply, &weightMsgNewReply, nil,
		func(_ *rand.Rand) {
			weightMsgNewReply = defaultWeightMsgNewReply
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgNewReply,
		exchangesimulation.SimulateMsgNewReply(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgNewAsk,
			defaultWeightMsgNewAsk,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				exchangesimulation.SimulateMsgNewAsk(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgNewReply,
			defaultWeightMsgNewReply,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				exchangesimulation.SimulateMsgNewReply(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
