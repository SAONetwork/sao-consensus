package keeper_test

import (
	"fmt"
	"math/big"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/x/node/keeper"
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/stretchr/testify/require"
)

func TestRandomIndex(t *testing.T) {
	k, _ := keepertest.NodeKeeper(t)
	blockhash := "4118E59A555FB33D1F401B4CD7050096DDD2743368235218069FCE9697C23AE7"
	header, _ := new(big.Int).SetString(blockhash, 16)
	idx := k.RandomIndex(header, 16, 5)
	require.Equal(t, idx, []int{3, 6, 9, 1, 10})
}

func TestRandomSP(t *testing.T) {
	setupMsgServer(t)
	nodes := make([]types.Node, 0)
	for i := 0; i < 10000; i++ {
		nodes = append(nodes, types.Node{
			Creator:         fmt.Sprintf("creator_%d", i),
			LastAliveHeight: int64(i * 100),
		})
	}

	for i, node := range keeper.SelectNodes(10, nodes) {
		fmt.Printf("node[%d]: %s\n", i, node.Creator)
		require.Equal(t, node.Creator, fmt.Sprintf("creator_%d", len(nodes)-1-i))
	}
}

func TestRandomSP2(t *testing.T) {
	setupMsgServer(t)
	nodes := make([]types.Node, 0)
	for i := 0; i < 10; i++ {
		nodes = append(nodes, types.Node{
			Creator:         fmt.Sprintf("creator_%d", i),
			LastAliveHeight: int64(i * 100),
		})
	}

	for i, node := range keeper.SelectNodes(20, nodes) {
		fmt.Printf("node[%d]: %s\n", i, node.Creator)
		require.Equal(t, node.Creator, fmt.Sprintf("creator_%d", len(nodes)-1-i))
	}
}

func TestRandomSP3(t *testing.T) {
	setupMsgServer(t)
	nodes := make([]types.Node, 0)
	for i := 0; i < 10000; i++ {
		nodes = append(nodes, types.Node{
			Creator:         fmt.Sprintf("creator_%d", i),
			Reputation:      float32(i * 10000.0),
			LastAliveHeight: int64(100),
		})
	}

	for i, node := range keeper.SelectNodes(5, nodes) {
		fmt.Printf("node[%d]: %s\n", i, node.Creator)
		require.Equal(t, node.Creator, fmt.Sprintf("creator_%d", len(nodes)-1-i))
	}
}
