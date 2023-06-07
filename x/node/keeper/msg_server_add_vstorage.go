package keeper

import (
	"context"

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

	amount := price.MulInt64(int64(msg.Size_)).TruncateInt()

	size := sdk.NewDecFromInt(amount).Quo(price).TruncateInt()

	coin := sdk.NewCoin(param.Baseline.Denom, size)

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
		}
	}

	pledge.TotalStorage += size.Int64()

	k.SetPledge(ctx, pledge)

	pool.TotalPledged.Amount = pool.TotalPledged.Amount.Add(amount)

	pool.TotalStorage += size.Int64()

	k.SetPool(ctx, pool)

	return &types.MsgAddVstorageResponse{}, nil
}
