package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ClaimReward(goCtx context.Context, msg *types.MsgClaimReward) (*types.MsgClaimRewardResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	pledge, found := k.GetPledge(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	logger := k.Logger(ctx)

	k.OrderPledge(ctx, msg.GetSigners()[0], sdk.NewInt64Coin(sdk.DefaultBondDenom, int64(0)))

	pledge, _ = k.GetPledge(ctx, msg.Creator)

	moduleaddr := k.ak.GetModuleAccount(ctx, types.ModuleName)

	logger.Debug("module ", "balance", k.bank.GetAllBalances(ctx, moduleaddr.GetAddress()))
	logger.Debug("pledge", "reward", pledge.Reward)

	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msg.GetSigners()[0], sdk.NewCoins(pledge.Reward))

	if err != nil {
		return nil, err
	}

	return &types.MsgClaimRewardResponse{}, nil
}
