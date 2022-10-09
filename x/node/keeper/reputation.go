package keeper

import (
	"math"
	"math/big"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
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

func (k Keeper) RandomSP(ctx sdk.Context, count int) []types.Node {
	header := new(big.Int).SetBytes(ctx.HeaderHash().Bytes())
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	// return all nodes
	nodes := k.GetAllNode(ctx)
	if len(nodes) <= count {
		return nodes
	}

	// nodes with reputation weight
	nodes = make([]types.Node, 0)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Node
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		reputation := val.Reputation
		for reputation > 100 {
			nodes = append(nodes, val)
			reputation -= 100
		}
	}
	total := len(nodes)
	sps := make([]types.Node, 0)
	for _, idx := range k.RandomIndex(header, total, count) {
		sps = append(sps, nodes[idx])
	}
	return sps
}
