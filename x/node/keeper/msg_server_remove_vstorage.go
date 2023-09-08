package keeper

import (
	"context"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) RemoveVstorage(goCtx context.Context, msg *types.MsgRemoveVstorage) (*types.MsgRemoveVstorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetNode(ctx, msg.Creator)

	if !found {
		return nil, status.Errorf(codes.NotFound, "node %d not found", msg.Creator)
	}

	pool, found := k.GetPool(ctx)

	if !found {
		return nil, status.Errorf(codes.NotFound, "pool not found")
	}

	pledge, found := k.GetPledge(ctx, msg.Creator)

	if !found {
		return nil, status.Errorf(codes.NotFound, "node %d not pledged yet", msg.Creator)
	}

	price := sdk.NewDecWithPrec(1, 6)

	param := k.GetParams(ctx)

	amount := price.MulInt64(int64(msg.Size_)).TruncateInt()
	if amount.IsZero() {
		return nil, status.Errorf(codes.InvalidArgument, "Removing %d bytes of storage does not release even 1 sao pledge, try increasing the remove size", msg.Size_)
	}

	size := sdk.NewDecFromInt(amount).Quo(price).Ceil().TruncateInt()

	if size.Int64() > pledge.TotalStorage-pledge.UsedStorage {
		return nil, sdkerrors.Wrap(types.ErrAvailableVstorage, "no enough available vstorage")
	}

	coin := sdk.NewCoin(param.Baseline.Denom, amount)

	pledge.TotalStoragePledged = pledge.TotalStoragePledged.Sub(coin)

	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msg.GetSigners()[0], sdk.Coins{coin})

	if err != nil {
		return nil, err
	}

	if pledge.TotalStorage > 0 {
		pending := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage).Sub(pledge.RewardDebt.Amount)
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
	}

	pledge.TotalStorage -= size.Int64()

	rewardDebt := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage)

	pledge.RewardDebt.Amount = rewardDebt

	pool.TotalPledged.Amount = pool.TotalPledged.Amount.Sub(amount)

	pool.TotalStorage -= size.Int64()

	// check super node
	if pledge.TotalStorage < k.VstorageThreshold(ctx) {
		node, found := k.GetNode(ctx, msg.Creator)
		if !found {
			return nil, types.ErrNodeNotFound
		}
		if node.Role == types.NODE_SUPER {
			node.Role = types.NODE_NORMAL
			k.SetNode(ctx, node)
		}
	}

	k.SetPledge(ctx, pledge)

	k.SetPool(ctx, pool)

	return &types.MsgRemoveVstorageResponse{}, nil
}
