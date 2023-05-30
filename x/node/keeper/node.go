package keeper

import (
	"bytes"
	"time"

	"github.com/SaoNetwork/sao/x/node/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

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

// GetAllNodesByStatus returns all nodes with the expected status and reputation
func (k Keeper) GetAllNodesByStatusAndReputation(ctx sdk.Context, status uint32, reputation float32) (list []types.Node) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var n types.Node
		k.cdc.MustUnmarshal(iterator.Value(), &n)
		if status&n.Status == status && n.Reputation >= reputation {
			list = append(list, n)
		}
	}

	return
}

func (k Keeper) EndBlock(ctx sdk.Context) {
	if ctx.BlockHeight()%1800 == 0 {
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NodeKeyPrefix))
		iterator := sdk.KVStorePrefixIterator(store, []byte{})

		defer iterator.Close()

		for ; iterator.Valid(); iterator.Next() {
			var n types.Node
			k.cdc.MustUnmarshal(iterator.Value(), &n)
			if n.Status&types.NODE_STATUS_ONLINE > 0 && ctx.BlockHeight()-n.LastAliveHeight > 3600 {
				n.Status = n.Status & (types.NODE_STATUS_NA ^ types.NODE_STATUS_ONLINE)
				b := k.cdc.MustMarshal(&n)
				store.Set(types.NodeKey(
					n.Creator,
				), b)
			}

			if n.Status&types.NODE_STATUS_ONLINE == 0 || ctx.BlockHeight()-n.LastAliveHeight > 10800 {
				store.Delete(types.NodeKey(n.Creator))
			}
		}
	}

	if ctx.BlockHeight()%600 == 0 {
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FaultKeyPrefix))
		iterator := sdk.KVStorePrefixIterator(store, []byte{})

		defer iterator.Close()

		totalPenaltyMap := make(map[string]uint64)
		for ; iterator.Valid(); iterator.Next() {
			var f types.Fault
			k.cdc.MustUnmarshal(iterator.Value(), &f)
			if f.Status == types.FaultStatusConfirmed {
				f.Penalty = f.Penalty + 1
				k.SetFault(ctx, &f)

				totalPenaltyMap[f.Provider] = totalPenaltyMap[f.Provider] + f.Penalty
			}
		}

		fishmenParamStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FishmenKeyPrefix))
		currentFishmenParamBytes := fishmenParamStore.Get([]byte("FishmenParam"))
		var fishmenParam types.FishmenParam
		err := fishmenParam.Unmarshal(currentFishmenParamBytes)
		if err == nil {
			for provider, totalPenalty := range totalPenaltyMap {
				if totalPenalty > fishmenParam.MaxPenalty {
					n, found := k.GetNode(ctx, provider)
					if found {
						n.Status = n.Status & (types.NODE_STATUS_NA ^ types.NODE_STATUS_ONLINE)
						b := k.cdc.MustMarshal(&n)
						store.Set(types.NodeKey(
							n.Creator,
						), b)
					}
				}
			}
		}
	}

	if ctx.BlockHeight()%(60*3) == 0 {
		// if ctx.BlockHeight()%(60*60*6) == 0 {
		proposals := k.gov.GetProposalsFiltered(ctx, v1.QueryProposalsParams{
			Limit:          1000,
			Depositor:      []byte(types.FISHMEN_LIST_DEPOSITOR),
			ProposalStatus: v1.ProposalStatus_PROPOSAL_STATUS_PASSED,
		})
		fishmenParamVersion := "Fishmen-" + time.Now().Format("2006-01")
		for _, proposal := range proposals {
			if len(proposal.Messages) != 1 {
				continue
			}

			message := proposal.Messages[0]
			if !(message.TypeUrl == "github.com/SaoNetwork/sao/x/node/types/FishmenParam") {
				continue
			}
			var fishmenParam types.FishmenParam
			err := fishmenParam.Unmarshal(message.Value)
			if err != nil {
				continue
			}

			if fishmenParam.Version == fishmenParamVersion {
				store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FishmenKeyPrefix))

				currentFishmenParamValue := store.Get([]byte("FishmenParam"))
				if !bytes.Equal(message.Value, currentFishmenParamValue) {
					store.Set([]byte("FishmenParam"), message.Value)
				}
				break
			}
		}
	}
}
