package keeper

import (
	"github.com/SaoNetwork/sao/x/earn/types"
)

var _ types.QueryServer = Keeper{}
