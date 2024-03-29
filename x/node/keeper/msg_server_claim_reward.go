package keeper

import (
	"context"

	markettypes "github.com/SaoNetwork/sao/x/market/types"
	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ClaimReward(goCtx context.Context, msg *types.MsgClaimReward) (*types.MsgClaimRewardResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	// claim pledge reward
	pledge, found := k.GetPledge(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrPledgeNotFound, "pledge not found")
	}

	logger := k.Logger(ctx)

	k.ShardRelease(ctx, msg.GetSigners()[0], nil)

	pledge, _ = k.GetPledge(ctx, msg.Creator)

	moduleaddr := k.ak.GetModuleAccount(ctx, types.ModuleName)

	logger.Debug("module ", "balance", k.bank.GetAllBalances(ctx, moduleaddr.GetAddress()))
	logger.Debug("pledge", "reward", pledge.Reward)

	claimReward, remainReward := pledge.Reward.TruncateDecimal()

	logger.Debug("PledgeTrace: claim reward",
		"sp", msg.Creator,
		"reward", pledge.Reward.String(),
		"remainReward", remainReward.String())

	pledge.Reward = remainReward

	// claim worker reward
	workerReward, err := k.market.Claim(ctx, claimReward.Denom, msg.Creator)
	if err != nil {
		return nil, err
	}

	k.RepayPledgeDebt(ctx, msg.Creator, []*sdk.Coin{&claimReward, &workerReward})

	if !claimReward.IsZero() {
		logger.Debug("CoinTrace: block reward", "from", types.ModuleName, "to", msg.GetSigners()[0], "amount", claimReward.String())

		err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msg.GetSigners()[0], sdk.Coins{claimReward})
		if err != nil {
			return nil, err
		}
	}
	if !workerReward.IsZero() {
		logger.Debug("CoinTrace: worker reward", "from", markettypes.ModuleName, "to", msg.GetSigners()[0], "amount", workerReward.String())
		err = k.bank.SendCoinsFromModuleToAccount(ctx, markettypes.ModuleName, msg.GetSigners()[0], sdk.Coins{workerReward})
		if err != nil {
			return nil, err
		}
	}

	k.SetPledge(ctx, pledge)

	claimReward = claimReward.Add(workerReward)

	return &types.MsgClaimRewardResponse{
		ClaimedReward: claimReward.Amount.Uint64(),
	}, nil
}
