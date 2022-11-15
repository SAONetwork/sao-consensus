package keeper

import (
	"github.com/SaoNetwork/sao/x/order/types"
)

var _ types.QueryServer = Keeper{}
