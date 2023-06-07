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

	size := sdk.NewDecFromInt(amount).Quo(price).TruncateInt()

	if size.Int64() > pledge.TotalStorage-pledge.UsedStorage {
		return nil, sdkerrors.Wrap(types.ErrAvailableVstorage, "no enough available vstorage")
	}

	coin := sdk.NewCoin(param.Baseline.Denom, size)

	err := k.bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, msg.GetSigners()[0], sdk.Coins{coin})

	if err != nil {
		return nil, err
	}

	pledge.TotalStorage -= size.Int64()

	k.SetPledge(ctx, pledge)

	pool.TotalPledged.Amount = pool.TotalPledged.Amount.Sub(amount)

	pool.TotalStorage -= size.Int64()

	k.SetPool(ctx, pool)

	return &types.MsgRemoveVstorageResponse{}, nil
}
