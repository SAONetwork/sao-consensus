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
	var status = types.NODE_STATUS_ONLINE | types.NODE_STATUS_SERVE_STORAGE | types.NODE_STATUS_ACCEPT_ORDER
	nodes := k.GetAllNodesByStatusAndReputation(ctx, status, 8000.0)
	if len(nodes) <= int(order.Replica) {
		return nodes
	}

	maxCandicates := len(nodes)
	if maxCandicates > int(order.Replica)*2 {
		maxCandicates = int(order.Replica) * 2
	}

	nodes = SelectNodes(maxCandicates, nodes)

	sps := make([]types.Node, 0)
	logger := k.Logger(ctx)
	for _, idx := range k.RandomIndex(header, maxCandicates, int(order.Replica)) {
		sps = append(sps, nodes[idx])
		logger.Error("RandomSP ###################", "Node", nodes[idx].Creator)

	}
	return sps
}

func SelectNodes(size int, nodes []types.Node) []types.Node {
	length := len(nodes)

	if length <= size {
		size = length
	}

	for i := 0; i <= size; i++ {
		buildHeap(nodes[i:])
	}

	return nodes[:size]
}

func buildHeap(nodes []types.Node) {
	size := len(nodes)
	for position := size/2 - 1; position >= 0; position-- {
		heapify(position, size, nodes)
	}
}

func heapify(position int, size int, nodes []types.Node) {
	if position >= size {
		return
	}

	cl := 2*position + 1
	cr := 2*position + 2
	index := position
	if cl < size && nodes[cl].LastAliveHeigh > nodes[index].LastAliveHeigh {
		nodes[index], nodes[cl] = nodes[cl], nodes[index]
	}
	if cr < size && nodes[cr].LastAliveHeigh > nodes[index].LastAliveHeigh {
		nodes[index], nodes[cr] = nodes[cr], nodes[index]
	}
}
