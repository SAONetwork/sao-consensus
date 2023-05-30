package keeper

import (
	"context"
	nodetypes "github.com/SaoNetwork/sao/x/node/types"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount := msg.Amount

	err := k.bank.SendCoinsFromAccountToModule(ctx, msg.GetSigners()[0], nodetypes.ModuleName, sdk.Coins{amount})
	if err != nil {
		return nil, err
	}

	decAmount := sdk.NewDecCoinFromCoin(amount)

	err = k.loan.Deposit(ctx, msg.Creator, decAmount)
	if err != nil {
		return nil, err
	}

	return &types.MsgDepositResponse{}, nil
}
