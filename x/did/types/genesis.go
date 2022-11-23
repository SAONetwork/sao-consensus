package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DidBindingProofsList:   []DidBindingProofs{},
		AccountListList:        []AccountList{},
		AccountAuthList:        []AccountAuth{},
		SidDocumentList:        []SidDocument{},
		SidDocumentVersionList: []SidDocumentVersion{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in didBindingProofs
	didBindingProofsIndexMap := make(map[string]struct{})

	for _, elem := range gs.DidBindingProofsList {
		index := string(DidBindingProofsKey(elem.AccountId))
		if _, ok := didBindingProofsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for didBindingProofs")
		}
		didBindingProofsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in accountList
	accountListIndexMap := make(map[string]struct{})

	for _, elem := range gs.AccountListList {
		index := string(AccountListKey(elem.Did))
		if _, ok := accountListIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for accountList")
		}
		accountListIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in accountAuth
	accountAuthIndexMap := make(map[string]struct{})

	for _, elem := range gs.AccountAuthList {
		index := string(AccountAuthKey(elem.AccountDid))
		if _, ok := accountAuthIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for accountAuth")
		}
		accountAuthIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in sidDocument
	sidDocumentIndexMap := make(map[string]struct{})

	for _, elem := range gs.SidDocumentList {
		index := string(SidDocumentKey(elem.VersionId))
		if _, ok := sidDocumentIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for sidDocument")
		}
		sidDocumentIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in sidDocumentVersion
	sidDocumentVersionIndexMap := make(map[string]struct{})

	for _, elem := range gs.SidDocumentVersionList {
		index := string(SidDocumentVersionKey(elem.DocId))
		if _, ok := sidDocumentVersionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for sidDocumentVersion")
		}
		sidDocumentVersionIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
