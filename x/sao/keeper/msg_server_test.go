package keeper_test

import (
	"context"
	"fmt"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/x/sao/keeper"
	"github.com/SaoNetwork/sao/x/sao/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SaoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestAddress(t *testing.T) {
	fmt.Println(sdk.MustAccAddressFromBech32("cosmos1r33rtwtgak2erkwq2462l3ed2ry2q0p0427eu9"))
}
