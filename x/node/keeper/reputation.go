package keeper

import (
	"math"
	"math/big"

	"github.com/SaoNetwork/sao/x/node/types"
	ordertypes "github.com/SaoNetwork/sao/x/order/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) IncreaseReputation(ctx sdk.Context, nodeId string, value float32) error {
	node, found := k.GetNode(ctx, nodeId)
	if !found {
		return types.ErrNodeNotFound
	}

	node.Reputation += value
	k.SetNode(ctx, node)

	return nil
}

func (k Keeper) DecreaseReputation(ctx sdk.Context, nodeId string, value float32) error {
	node, found := k.GetNode(ctx, nodeId)
	if !found {
		return types.ErrNodeNotFound
	}

	node.Reputation -= value
	k.SetNode(ctx, node)

	return nil
}

func (k Keeper) RandomIndex(seed *big.Int, total, count int) []int {
	idx := make([]int, 0)
	mod := math.Pow10(int(math.Ceil(math.Log10(float64(total)))))
	if total <= count {
		return idx
	}
	for count > 0 {
		rs := int(new(big.Int).Mod(seed, big.NewInt(int64(mod))).Int64()) % total
		seed = new(big.Int).Div(seed, big.NewInt(10))
		duplicate := false
		for _, v := range idx {
			if rs == v {
				duplicate = true
			}
		}
		if duplicate {
			continue
		}
		idx = append(idx, rs)
		count -= 1
	}
	return idx
}

func (k Keeper) RandomSP(ctx sdk.Context, order ordertypes.Order) []types.Node {
	header := new(big.Int).SetBytes(ctx.HeaderHash().Bytes())

	// return all avaliable storage nodes
	var status = types.NODE_STATUS_SERVE_STORAGE | types.NODE_STATUS_ACCEPT_ORDER
	nodes := k.GetAllNodesByStatusAndReputation(ctx, status, 10000.0)
	if len(nodes) <= int(order.Replica) {
		return nodes
	}

	maxCandicates := len(nodes)
	if int(order.Replica)*2 < maxCandicates {
		maxCandicates = int(order.Replica) * 2
	}

	for round := 0; round < maxCandicates; round++ {
		for index := 1; index < len(nodes); index++ {
			if nodes[index].LastAliveHeigh > nodes[index-1].LastAliveHeigh {
				node := nodes[index]
				nodes[index] = nodes[index-1]
				nodes[index-1] = node
			}
		}
	}

	sps := make([]types.Node, 0)
	for _, idx := range k.RandomIndex(header, maxCandicates, int(order.Replica)) {
		sps = append(sps, nodes[idx])
	}
	return sps
}
