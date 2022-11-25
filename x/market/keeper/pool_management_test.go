package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestCoin(t *testing.T) {
	coin := sdk.NewInt64Coin("stake", 1000)
	deccoin := sdk.NewDecCoinFromCoin(coin)
	fmt.Println(coin)
	fmt.Println(deccoin)
	fmt.Println(deccoin.Amount.Abs().TruncateInt())

}
