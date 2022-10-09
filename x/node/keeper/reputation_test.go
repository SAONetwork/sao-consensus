package keeper_test

import (
	"math/big"
	"testing"

	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestRandomIndex(t *testing.T) {
	k, _ := keepertest.NodeKeeper(t)
	blockhash := "4118E59A555FB33D1F401B4CD7050096DDD2743368235218069FCE9697C23AE7"
	header, _ := new(big.Int).SetString(blockhash, 16)
	idx := k.RandomIndex(header, 16, 5)
	require.Equal(t, idx, []int{3, 6, 9, 1, 10})
}
