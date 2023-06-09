package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// 1 byte to record which current super node is
// is 256 enough to track super nodes?
func (k Keeper) SetNodeRound(ctx sdk.Context, round uint8) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeRoundKeyPrefix))
	store.Set(types.NodeRoundKey(), []byte{round})
}

func (k Keeper) GetNodeRound(ctx sdk.Context) (round uint8, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeRoundKeyPrefix))
	b := store.Get(types.NodeRoundKey())
	if b == nil {
		return round, false
	}
	return b[0], true
}

// SetNode set a specific node in the store from its index
func (k Keeper) SetNode(ctx sdk.Context, node types.Node) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
	b := k.cdc.MustMarshal(&node)
	store.Set(types.NodeKey(
		node.Creator,
	), b)
}

// GetNode returns a node from its index
func (k Keeper) GetNode(
	ctx sdk.Context,
	creator string,

) (val types.Node, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))

	b := store.Get(types.NodeKey(
		creator,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNode removes a node from the store
func (k Keeper) RemoveNode(
	ctx sdk.Context,
	creator string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
	store.Delete(types.NodeKey(
		creator,
	))
}

// GetAllNodes returns all nodes
func (k Keeper) GetAllNode(ctx sdk.Context) (list []types.Node) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Node
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllNodesByStatus returns all nodes with the expected status
func (k Keeper) GetAllNodesByStatus(ctx sdk.Context, status uint32) (list []types.Node) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var n types.Node
		k.cdc.MustUnmarshal(iterator.Value(), &n)
		if status&n.Status == status {
			list = append(list, n)
		}
	}

	return
}

/**
 * Find all valid super nodes
 */
func (k Keeper) GetAllSuperNodes(ctx sdk.Context) (list []types.Node) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var n types.Node
		k.cdc.MustUnmarshal(iterator.Value(), &n)
		if n.Role == types.NODE_SUPER {
			list = append(list, n)
		}
	}
	return
}

/**
 * Get next super node to accept order
 */
func (k Keeper) GetNextSuperNodes(ctx sdk.Context, status uint32, reputation float32, ignore []string) types.Node {
	roundStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeRoundKeyPrefix))
	round := roundStore.Get(types.NodeRoundKey())

	if len(round) == 0 {
		// first time after upgrade
		round = []byte{0}
		roundStore.Set(types.NodeRoundKey(), round)
	}

	snodes := k.GetAllSuperNodes(ctx)
	i := uint8(round[0])
	if len(snodes) > 0 {
		for {
			if i >= uint8(len(snodes)) {
				i = 0
			}
			toIgnore := false
			for _, ig := range ignore {
				if ig == snodes[i].Creator {
					toIgnore = true
					break
				}
			}
			if !toIgnore {
				if status&snodes[i].Status == status && snodes[i].Reputation >= reputation {
					// update next round
					if int(i+1) >= len(snodes) {
						roundStore.Set(types.NodeRoundKey(), []byte{0})
					} else {
						roundStore.Set(types.NodeRoundKey(), []byte{i + 1})
					}
					return snodes[i]
				}
			}
			// if all super nodes don't satify, quit
			if round[0] == 0 {
				if i == uint8(len(snodes)-1) {
					break
				}
			} else {
				if i == uint8(round[0]-1) {
					break
				}
			}
			i++
		}
	}
	return types.Node{}
}

// GetAllNodesByStatus returns all nodes with the expected status and reputation
func (k Keeper) GetAllNodesByStatusAndReputationAndRole(ctx sdk.Context, role uint32, status uint32, reputation float32) (list []types.Node) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var n types.Node
		k.cdc.MustUnmarshal(iterator.Value(), &n)
		if status&n.Status == status && n.Reputation >= reputation && n.Role == role {
			list = append(list, n)
		}
	}

	return
}
