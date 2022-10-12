package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/x/earn/keeper"
	"github.com/SaoNetwork/sao/x/earn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EarnKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
