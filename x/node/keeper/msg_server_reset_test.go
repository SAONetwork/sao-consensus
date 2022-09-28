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

func TestRestNode(t *testing.T) {
	server, _, context := login(t)
	// should success
	resetResponse, err := server.Reset(context, &types.MsgReset{
		Creator: alice,
		Peer:    newPeerId,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgResetResponse{}, *resetResponse)

}

func TestRestNodeWithOtherAddress(t *testing.T) {
	k, ctx := keepertest.NodeKeeper(t)
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	// should failed, invalid peer
	_, err := server.Reset(context, &types.MsgReset{
		Creator: bob,
		Peer:    newPeerId,
	})
	require.Equal(t, errors.Is(err, types.ErrNodeNotFound), true)
}

func TestRestNodeWithInvalidPeer(t *testing.T) {
	k, ctx := keepertest.NodeKeeper(t)
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)

	// should failed, invalid peer
	_, err := server.Reset(context, &types.MsgReset{
		Creator: bob,
		Peer:    invalidPeerId,
	})
	require.Equal(t, errors.Is(err, types.ErrNodeNotFound), true)
}
