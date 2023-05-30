package keeper

import (
	"context"
	nodetypes "github.com/SaoNetwork/sao/x/node/types"

	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amount := msg.Amount

	decAmount := sdk.NewDecCoinFromCoin(amount)

	err := k.loan.Withdraw(ctx, msg.Creator, decAmount)
	if err != nil {
		return nil, err
	}

	err = k.bank.SendCoinsFromModuleToAccount(ctx, nodetypes.ModuleName, msg.GetSigners()[0], sdk.Coins{amount})
	if err != nil {
		return nil, err
	}

	return &types.MsgWithdrawResponse{}, nil
}
