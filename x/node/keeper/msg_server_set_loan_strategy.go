package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetLoanStrategy(goCtx context.Context, msg *types.MsgSetLoanStrategy) (*types.MsgSetLoanStrategyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denom := k.staking.BondDenom(ctx)

	pledge, found := k.GetPledge(ctx, msg.Creator)
	if !found {
		pledge = types.Pledge{
			Creator:             msg.Creator,
			TotalStoragePledged: sdk.NewCoin(denom, sdk.NewInt(0)),
			Reward:              sdk.NewDecCoin(denom, sdk.NewInt(0)),
			RewardDebt:          sdk.NewDecCoin(denom, sdk.NewInt(0)),
			TotalStorage:        0,
			LoanStrategy:        msg.LoanStrategy,
			LoanPledged:         sdk.NewCoin(denom, sdk.NewInt(0)),
			LastRewardAt:        uint64(ctx.BlockHeight()),
		}
	} else {
		pledge.LoanStrategy = msg.LoanStrategy
	}

	k.SetPledge(ctx, pledge)

	return &types.MsgSetLoanStrategyResponse{}, nil
}
