package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MetadataList: []Metadata{},
		ModelList:    []Model{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in metadata
	metadataIndexMap := make(map[string]struct{})

	for _, elem := range gs.MetadataList {
		index := string(MetadataKey(elem.DataId))
		if _, ok := metadataIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for metadata")
		}
		metadataIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in model
	modelIndexMap := make(map[string]struct{})

	for _, elem := range gs.ModelList {
		index := string(ModelKey(elem.Key))
		if _, ok := modelIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for model")
		}
		modelIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
