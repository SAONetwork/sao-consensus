package keeper

import (
	"github.com/SaoNetwork/sao/x/node/types"
)

var _ types.QueryServer = Keeper{}
