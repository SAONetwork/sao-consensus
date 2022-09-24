package keeper

import (
	"github.com/SaoNetwork/sao/x/sao/types"
)

var _ types.QueryServer = Keeper{}
