package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/earn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ClaimReward(goCtx context.Context, msg *types.MsgClaimReward) (*types.MsgClaimRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	pledge, found := k.GetPledge(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrPoolNotFound, "")
	}

	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msg.GetSigners()[0], sdk.NewCoins(pledge.Reward))
	if err != nil {
		return nil, err
	}

	pledge.Reward.Sub(pledge.Reward)

	return &types.MsgClaimRewardResponse{}, nil
}
