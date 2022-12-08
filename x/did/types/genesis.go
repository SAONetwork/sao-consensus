package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DidBingingProofList:    []DidBingingProof{},
		AccountListList:        []AccountList{},
		AccountAuthList:        []AccountAuth{},
		SidDocumentList:        []SidDocument{},
		SidDocumentVersionList: []SidDocumentVersion{},
		PastSeedsList:          []PastSeeds{},
		PaymentAddressList:     []PaymentAddress{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in DidBingingProof
	DidBingingProofIndexMap := make(map[string]struct{})

	for _, elem := range gs.DidBingingProofList {
		index := string(DidBingingProofKey(elem.AccountId))
		if _, ok := DidBingingProofIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for DidBingingProof")
		}
		DidBingingProofIndexMap[index] = struct{}{}
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
	// Check for duplicated index in pastSeeds
	pastSeedsIndexMap := make(map[string]struct{})

	for _, elem := range gs.PastSeedsList {
		index := string(PastSeedsKey(elem.Did))
		if _, ok := pastSeedsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pastSeeds")
		}
		pastSeedsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in paymentAddress
	paymentAddressIndexMap := make(map[string]struct{})

	for _, elem := range gs.PaymentAddressList {
		index := string(PaymentAddressKey(elem.Did))
		if _, ok := paymentAddressIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for paymentAddress")
		}
		paymentAddressIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
