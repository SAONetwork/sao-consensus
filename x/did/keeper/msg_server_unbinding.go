package keeper

import (
	"context"
	"strings"

	"github.com/SaoNetwork/sao/x/did/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Unbinding(goCtx context.Context, msg *types.MsgUnbinding) (*types.MsgUnbindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	accountId := msg.GetAccountId()
	proof, found := k.GetDidBindingProof(ctx, accountId)
	if !found {
		return nil, types.ErrBindingNotFound
	}

	payAddr, found := k.GetPaymentAddress(ctx, proof.Proof.Did)
	accIdSplits := strings.Split(accountId, ":")
	if found &&
		len(accIdSplits) == 3 &&
		accIdSplits[0] == "cosmos" &&
		accIdSplits[1] == ctx.ChainID() &&
		accIdSplits[2] == payAddr.Address {
		return nil, types.ErrUnbindPayAddr
	}

	k.RemoveDidBindingProof(ctx, accountId)

	return &types.MsgUnbindingResponse{}, nil
}
