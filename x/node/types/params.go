package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyBlockReward = []byte("BlockReward")
	// TODO: Determine the default value
	DefaultBlockReward sdk.Coin
)

var (
	KeyBaseLine = []byte("Baseline")
	// TODO: Determine the default value
	DefaultBaseline sdk.Coin
)

var (
	KeyAPY = []byte("AnnualPercentageYield")
	// TODO: Determine the default value
	DefaultAPY = sdk.NewDecWithPrec(50, 2)
)

var (
	KeyHalvingPeriod     = []byte("HalvingPeriod")
	DefaultHalvingPeriod = int64(32000000)
)

var (
	KeyAdjustmentPeriod     = []byte("AdjustmentPeriod")
	DefaultAdjustmentPeriod = int64(2000)
)

var (
	KeyFishmenInfo     = []byte("FishmenInfo")
	DefaultFishmenInfo = ""
)

var (
	KeyPenaltyBase     = []byte("PenaltyBase")
	DefaultPenaltyBase = 1
)

var (
	KeyMaxPenalty         = []byte("MaxPenalty")
	DefaultMaxPenalty     = 10000
	KeyShareThreshold     = []byte("ShareThreshold")
	DefaultShareThreshold = sdk.NewDecWithPrec(10, 2)
)

var (
	KeyVstorageThreshold     = []byte("VstorageThreshold")
	DefaultVstorageThreshold = int64(10 << 30) // 10 Gb
)

var (
	KeyOfflineTriggerHeight     = []byte("OfflineTriggerHeight")
	DefaultOfflineTriggerHeight = int64(1800) // 1 hour
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	blockReward sdk.Coin,
	baseline sdk.Coin,
	apy sdk.Dec,
	halving int64,
	adjustment int64,
	fishmenInfo string,
	penaltyBase uint64,
	maxPenalty uint64,
	threshold sdk.Dec,
	minVstorage int64,
	offlineTriggerHeight int64,
) Params {
	return Params{
		BlockReward:           blockReward,
		Baseline:              baseline,
		AnnualPercentageYield: apy.String(),
		HalvingPeriod:         halving,
		AdjustmentPeriod:      adjustment,
		FishmenInfo:           fishmenInfo,
		PenaltyBase:           penaltyBase,
		MaxPenalty:            maxPenalty,
		ShareThreshold:        threshold.String(),
		VstorageThreshold:     minVstorage,
		OfflineTriggerHeight:  offlineTriggerHeight,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultBlockReward,
		DefaultBaseline,
		DefaultAPY,
		DefaultHalvingPeriod,
		DefaultAdjustmentPeriod,
		DefaultFishmenInfo,
		uint64(DefaultPenaltyBase),
		uint64(DefaultMaxPenalty),
		DefaultShareThreshold,
		DefaultVstorageThreshold,
		DefaultOfflineTriggerHeight,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBlockReward, &p.BlockReward, validateBlockReward),
		paramtypes.NewParamSetPair(KeyBaseLine, &p.Baseline, validateBaseline),
		paramtypes.NewParamSetPair(KeyAPY, &p.AnnualPercentageYield, validateAPY),
		paramtypes.NewParamSetPair(KeyHalvingPeriod, &p.HalvingPeriod, validatePeriod),
		paramtypes.NewParamSetPair(KeyAdjustmentPeriod, &p.AdjustmentPeriod, validatePeriod),
		paramtypes.NewParamSetPair(KeyFishmenInfo, &p.FishmenInfo, validateFishmenInfo),
		paramtypes.NewParamSetPair(KeyPenaltyBase, &p.PenaltyBase, validatePenaltyBase),
		paramtypes.NewParamSetPair(KeyMaxPenalty, &p.MaxPenalty, validateMaxPenalty),
		paramtypes.NewParamSetPair(KeyShareThreshold, &p.ShareThreshold, validateShareThreshold),
		paramtypes.NewParamSetPair(KeyVstorageThreshold, &p.VstorageThreshold, validateVstorageThreshold),
		paramtypes.NewParamSetPair(KeyOfflineTriggerHeight, &p.OfflineTriggerHeight, validateOfflineTriggerHeight),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateBlockReward(p.BlockReward); err != nil {
		return err
	}

	if err := validateBaseline(p.Baseline); err != nil {
		return err
	}

	if err := validateAPY(p.AnnualPercentageYield); err != nil {
		return err
	}

	if err := validatePeriod(p.HalvingPeriod); err != nil {
		return err
	}

	if err := validatePeriod(p.AdjustmentPeriod); err != nil {
		return err
	}

	if err := validateFishmenInfo(p.FishmenInfo); err != nil {
		return err
	}

	if err := validatePenaltyBase(p.PenaltyBase); err != nil {
		return err
	}

	if err := validateMaxPenalty(p.MaxPenalty); err != nil {
		return err
	}

	if err := validateShareThreshold(p.ShareThreshold); err != nil {
		return err
	}

	if err := validateVstorageThreshold(p.VstorageThreshold); err != nil {
		return err
	}

	if err := validateOfflineTriggerHeight(p.OfflineTriggerHeight); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateBlockReward validates the BlockReward param
func validateBlockReward(v interface{}) error {
	_ = v.(sdk.Coin)

	return nil
}

// validateBaseline validates the BlockReward param
func validateBaseline(v interface{}) error {
	_ = v.(sdk.Coin)

	return nil
}

func validatePeriod(v interface{}) error {
	p := v.(int64)
	if p > 10 {
		return nil
	}
	return errors.New("invalid period")
}

// validateAPY validates the BlockReward param
func validateAPY(v interface{}) error {
	_, err := sdk.NewDecFromStr(v.(string))
	return err
}

// validateFishmenInfo validates the Fishmen list
func validateFishmenInfo(v interface{}) error {
	return nil
}

// validatePenaltyBase validates penalty base
func validatePenaltyBase(v interface{}) error {
	p := v.(uint64)
	if p > 0 {
		return nil
	}
	return errors.New("invalid penalty base")
}

// validateMaxPenalty validates max penalty
func validateMaxPenalty(v interface{}) error {
	p := v.(uint64)
	if p > 10 {
		return nil
	}
	return errors.New("invalid max penalty")
}

func validateShareThreshold(v interface{}) error {
	t, err := sdk.NewDecFromStr(v.(string))
	if t.MustFloat64() < 0.01 {
		return errors.New("invalid share threshold")
	}
	return err
}

func validateVstorageThreshold(v interface{}) error {
	p := v.(int64)
	if p > 0 {
		return nil
	}
	return errors.New("invalid vstorage threshold")
}

func validateOfflineTriggerHeight(v interface{}) error {
	p := v.(int64)
	if p > 0 {
		return nil
	}
	return errors.New("invalid offline trigger height")
}
