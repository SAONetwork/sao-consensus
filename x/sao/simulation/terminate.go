package simulation

import (
	"math/rand"

	"github.com/SaoNetwork/sao/x/sao/keeper"
	"github.com/SaoNetwork/sao/x/sao/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgTerminate(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgTerminate{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Terminate simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Terminate simulation not implemented"), nil, nil
	}
}
