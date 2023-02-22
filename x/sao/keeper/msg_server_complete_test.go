package keeper_test

import (
	"fmt"
	"math/big"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestCompleteOrder(t *testing.T) {

	price := new(big.Int).Quo(big.NewInt(1000000000000000000), big.NewInt(1048576))
	amount := new(big.Int).Mul(price, big.NewInt(1024))
	decamount := sdk.NewDecFromBigIntWithPrec(amount, 18)
	coin, decoin := sdk.NewDecCoinFromDec("sao", decamount).TruncateDecimal()
	fmt.Println(coin, decoin)
}
