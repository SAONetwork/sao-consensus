package types

import (
	fmt "fmt"
)

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ExpiredOrderList: []ExpiredOrder{},
		// this line is used by starport scaffolding # genesis/types/default
		OrderList: []Order{},
		Params:    DefaultParams(),
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
	// Check for duplicated index in expiredOrder
	expiredOrderIndexMap := make(map[string]struct{})

	for _, elem := range gs.ExpiredOrderList {
		index := string(ExpiredOrderKey(elem.Height))
		if _, ok := expiredOrderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for expiredOrder")
		}
		expiredOrderIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
