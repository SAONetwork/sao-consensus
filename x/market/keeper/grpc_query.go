package keeper

import (
	"github.com/SaoNetwork/sao/x/market/types"
)

var _ types.QueryServer = Keeper{}
