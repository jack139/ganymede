package simulation

import (
	"math/rand"
	"strconv"

	simappparams "cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jack139/ganymede/ganymede/x/zoo/keeper"
	"github.com/jack139/ganymede/ganymede/x/zoo/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateKvzoo(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateKvzoo{
			Creator: simAccount.Address.String(),
			Owner:   strconv.Itoa(i),
			ZooKey:  strconv.Itoa(i),
		}

		_, found := k.GetKvzoo(ctx, msg.Owner, msg.ZooKey)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Kvzoo already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateKvzoo(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			kvzoo      = types.Kvzoo{}
			msg        = &types.MsgUpdateKvzoo{}
			allKvzoo   = k.GetAllKvzoo(ctx)
			found      = false
		)
		for _, obj := range allKvzoo {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				kvzoo = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "kvzoo creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Owner = kvzoo.Owner
		msg.ZooKey = kvzoo.ZooKey

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteKvzoo(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			kvzoo      = types.Kvzoo{}
			msg        = &types.MsgUpdateKvzoo{}
			allKvzoo   = k.GetAllKvzoo(ctx)
			found      = false
		)
		for _, obj := range allKvzoo {
			simAccount, found = FindAccount(accs, obj.Creator)
			if found {
				kvzoo = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "kvzoo creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.Owner = kvzoo.Owner
		msg.ZooKey = kvzoo.ZooKey

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
