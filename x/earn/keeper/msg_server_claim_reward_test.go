package keeper_test

import (
	"context"
	"fmt"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/x/earn"
	"github.com/SaoNetwork/sao/x/earn/keeper"
	"github.com/SaoNetwork/sao/x/earn/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	alice = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob   = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
)

func setupPledge(t *testing.T) (types.MsgServer, *keeper.Keeper, context.Context, sdk.Context) {

	genesisState := types.DefaultGenesis()

	k, ctx := keepertest.EarnKeeper(t)

	k.SetPool(ctx, *genesisState.Pool)

	sp, _ := sdk.AccAddressFromBech32(bob)
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	coin := sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)
	k.OrderPledge(ctx, sp, coin)

	return server, k, context, ctx
}

func TestClaimReward(t *testing.T) {

	server, k, context, ctx := setupPledge(t)

	earn.BeginBlocker(ctx, *k)

	_, err := server.ClaimReward(context, &types.MsgClaimReward{
		Creator: bob,
	})

	fmt.Println(err)
}
