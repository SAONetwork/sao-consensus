package keeper_test

import (
	"errors"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/x/node/keeper"
	"github.com/SaoNetwork/sao/x/node/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice         = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	bob           = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
	carol         = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
	peerId        = "/ip4/127.0.0.1/tcp/4001/p2p/12D3KooWMVpqMrAYwPH78BowbhDFKjwTgbBuRtzLXWd4PHTAFY4E"
	newPeerId     = "/ip4/192.168.1.1/tcp/4001/p2p/12D3KooWMVpqMrAYwPH78BowbhDFKjwTgbBuRtzLXWd4PHTAFY4E"
	invalidPeerId = "/ip4/127.0.0.1/tcp/4001/p2p"
)

func TestLogin(t *testing.T) {
	k, ctx := keepertest.NodeKeeper(t)
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	// should success
	registryResponse, err := server.Login(context, &types.MsgLogin{
		Creator: bob,
		Peer:    peerId,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgLoginResponse{}, *registryResponse)

}

func TestLoginWithInvalidPeer(t *testing.T) {
	k, ctx := keepertest.NodeKeeper(t)
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	// should failed, invalid peer
	_, err := server.Login(context, &types.MsgLogin{
		Creator: alice,
		Peer:    invalidPeerId,
	})

	require.Equal(t, errors.Is(err, types.ErrInvalidPeer), true)
}

func TestNodeAlreadyLogin(t *testing.T) {
	k, ctx := keepertest.NodeKeeper(t)
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	server.Login(context, &types.MsgLogin{
		Creator: bob,
		Peer:    peerId,
	})
	// should failed, already registered
	_, err := server.Login(context, &types.MsgLogin{
		Creator: bob,
		Peer:    peerId,
	})
	require.Equal(t, errors.Is(err, types.ErrAlreadyRegistered), true)
}
