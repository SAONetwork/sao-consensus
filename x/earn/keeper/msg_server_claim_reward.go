package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/earn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ClaimReward(goCtx context.Context, msg *types.MsgClaimReward) (*types.MsgClaimRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	logger := k.Logger(ctx)

	pool, pool_found := k.GetPool(ctx)

	if !pool_found {
		return nil, sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	pledge, pledge_found := k.GetPledge(ctx, msg.Creator)
	if !pledge_found {
		return nil, sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	reward := uint64(pledge.Pledged.Amount.Int64())*pool.CoinPerShare/1e12 - uint64(pledge.RewardDebt.Amount.Int64())

	if reward == 0 {
		return nil, sdkerrors.Wrap(types.ErrInsufficientCoin, "")
	}

	logger.Info("#########", types.ModuleName, msg.GetSigners()[0])
	addr := k.ak.GetModuleAddress(types.ModuleName)
	balance := k.bank.GetAllBalances(ctx, addr)
	rewardCoin := sdk.NewInt64Coin(pool.Denom.Denom, int64(reward))
	logger.Info("####", "balance", balance, "pledge", reward, "rewardcoin", rewardCoin)

	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msg.GetSigners()[0], sdk.NewCoins(rewardCoin))

	if err != nil {
		logger.Info(err.Error())
		return nil, err
	}

	pledge.RewardDebt = pledge.RewardDebt.Add(rewardCoin)

	k.SetPledge(ctx, pledge)

	return &types.MsgClaimRewardResponse{}, nil
}
