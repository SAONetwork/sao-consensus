package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		OrderList: []Order{},
		ShardList: []Shard{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in order
	orderIdMap := make(map[uint64]bool)
	orderCount := gs.GetOrderCount()
	for _, elem := range gs.OrderList {
		if _, ok := orderIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for order")
		}
		if elem.Id >= orderCount {
			return fmt.Errorf("order id should be lower or equal than the last id")
		}
		orderIdMap[elem.Id] = true
	}
	// Check for duplicated ID in shard
	shardIdMap := make(map[uint64]bool)
	shardCount := gs.GetShardCount()
	for _, elem := range gs.ShardList {
		if _, ok := shardIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for shard")
		}
		if elem.Id >= shardCount {
			return fmt.Errorf("shard id should be lower or equal than the last id")
		}
		shardIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
