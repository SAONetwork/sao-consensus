package keeper_test

import (
	"context"
	"errors"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/x/node/keeper"
	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func login(t *testing.T) (types.MsgServer, *keeper.Keeper, context.Context) {

	k, ctx := keepertest.NodeKeeper(t)
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	server.Login(context, &types.MsgLogin{
		Creator: alice,
		Peer:    "/ip4/127.0.0.1/tcp/4001/p2p/12D3KooWMVpqMrAYwPH78BowbhDFKjwTgbBuRtzLXWd4PHTAFY4E",
	})
	return server, k, context
}

func TestLogout(t *testing.T) {
	server, _, context := login(t)
	// should success
	logoutResponse, err := server.Logout(context, &types.MsgLogout{
		Creator: alice,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgLogoutResponse{}, *logoutResponse)

}

func TestLogoutWithOtherAddress(t *testing.T) {
	k, ctx := keepertest.NodeKeeper(t)
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	// should failed, invalid peer
	_, err := server.Logout(context, &types.MsgLogout{
		Creator: bob,
	})
	require.Equal(t, errors.Is(err, types.ErrNodeNotFound), true)
}
