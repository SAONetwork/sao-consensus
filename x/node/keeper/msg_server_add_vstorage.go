package keeper

import (
	"context"
	"math"

	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k msgServer) AddVstorage(goCtx context.Context, msg *types.MsgAddVstorage) (*types.MsgAddVstorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetNode(ctx, msg.Creator)

	if !found {
		return nil, status.Errorf(codes.NotFound, "node %d not found", msg.Creator)
	}

	pool, found := k.GetPool(ctx)

	if !found {
		return nil, status.Errorf(codes.NotFound, "pool not found")
	}

	price := sdk.NewDecWithPrec(1, 6)

	param := k.GetParams(ctx)

	amount := price.MulInt64(int64(msg.Size_)).Ceil().TruncateInt()

	size := sdk.NewDecFromInt(amount).Quo(price).TruncateInt()

	coin := sdk.NewCoin(param.Baseline.Denom, amount)

	err := k.bank.SendCoinsFromAccountToModule(ctx, msg.GetSigners()[0], types.ModuleName, sdk.Coins{coin})

	if err != nil {
		return nil, err
	}

	pledge, found := k.GetPledge(ctx, msg.Creator)
	if !found {
		pledge = types.Pledge{
			Creator:             msg.Creator,
			TotalStorage:        0,
			UsedStorage:         0,
			TotalStoragePledged: coin,
			TotalShardPledged:   sdk.NewInt64Coin(coin.Denom, 0),
			Reward:              sdk.NewInt64DecCoin(coin.Denom, 0),
			RewardDebt:          sdk.NewInt64DecCoin(coin.Denom, 0),
		}
	} else {
		pledge.TotalStoragePledged = pledge.TotalStoragePledged.Add(coin)
	}

	if pledge.TotalStorage > 0 {
		pending := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage).Sub(pledge.RewardDebt.Amount)
		pledge.Reward.Amount = pledge.Reward.Amount.Add(pending)
	}

	pledge.TotalStorage += size.Int64()

	rewardDebt := pool.AccRewardPerByte.Amount.MulInt64(pledge.TotalStorage)

	pledge.RewardDebt.Amount = rewardDebt

	pool.TotalPledged.Amount = pool.TotalPledged.Amount.Add(amount)

	pool.TotalStorage += size.Int64()

	// check super node
	if pledge.TotalStorage >= k.VstorageThreshold(ctx) {
		node, found := k.GetNode(ctx, msg.Creator)
		if !found {
			return nil, types.ErrNodeNotFound
		}
		if node.Role == types.NODE_NORMAL && node.Status&types.NODE_STATUS_SUPER_REQUIREMENT == types.NODE_STATUS_SUPER_REQUIREMENT {
			if node.Validator != "" {
				err := k.CheckDelegationShare(ctx, msg.Creator, node.Validator, sdk.NewDec(0))
				if err == nil && node.Role == types.NODE_NORMAL {
					node.Role = types.NODE_SUPER
					k.SetNode(ctx, node)
				}
			} else {
				accAddr := sdk.MustAccAddressFromBech32(msg.Creator)
				dels := k.staking.GetDelegatorDelegations(ctx, accAddr, math.MaxUint16)
				for _, del := range dels {
					err := k.CheckDelegationShare(ctx, msg.Creator, del.ValidatorAddress, sdk.NewDec(0))
					if err == nil && node.Role == types.NODE_NORMAL {
						node.Validator = del.ValidatorAddress
						node.Role = types.NODE_SUPER
						k.SetNode(ctx, node)
						break
					}
				}
			}
		}
	}

	k.SetPledge(ctx, pledge)

	k.SetPool(ctx, pool)

	return &types.MsgAddVstorageResponse{}, nil
}
