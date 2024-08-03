package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/jack139/ganymede/ganymede/x/exchange/keeper"
	"github.com/jack139/ganymede/ganymede/x/exchange/types"
)

func SimulateMsgNewReply(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgNewReply{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the NewReply simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "NewReply simulation not implemented"), nil, nil
	}
}
