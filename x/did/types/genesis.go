package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AccountListList:        []AccountList{},
		AccountAuthList:        []AccountAuth{},
		SidDocumentList:        []SidDocument{},
		SidDocumentVersionList: []SidDocumentVersion{},
		PastSeedsList:          []PastSeeds{},
		PaymentAddressList:     []PaymentAddress{},
		AccountIdList:          []AccountId{},
		DidList:                []Did{},
		KidList:                []Kid{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
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
	// Check for duplicated index in accountId
	accountIdIndexMap := make(map[string]struct{})

	for _, elem := range gs.AccountIdList {
		index := string(AccountIdKey(elem.AccountDid))
		if _, ok := accountIdIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for accountId")
		}
		accountIdIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in did
	didIndexMap := make(map[string]struct{})

	for _, elem := range gs.DidList {
		index := string(DidKey(elem.AccountId))
		if _, ok := didIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for did")
		}
		didIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in kid
	kidIndexMap := make(map[string]struct{})

	for _, elem := range gs.KidList {
		index := string(KidKey(elem.Address))
		if _, ok := kidIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for kid")
		}
		kidIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
