package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {

	pool := Pool{
		TotalPledged:       sdk.NewInt64Coin("sao", 0),
		TotalReward:        sdk.NewInt64Coin("sao", 0),
		AccRewardPerByte:   sdk.NewInt64DecCoin("sao", 0),
		TotalStorage:       0,
		RewardedBlockCount: 0,
	}
	return &GenesisState{
		Pool:      &pool,
		NodeList:  []Node{},
		ShardList: []Shard{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in node
	nodeIndexMap := make(map[string]struct{})

	for _, elem := range gs.NodeList {
		index := string(NodeKey(elem.Creator))
		if _, ok := nodeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for node")
		}
		nodeIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in shard
	shardIndexMap := make(map[string]struct{})

	for _, elem := range gs.ShardList {
		index := string(ShardKey(elem.Idx))
		if _, ok := shardIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for shard")
		}
		shardIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
