package keeper_test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"encoding/hex"

	"github.com/SaoNetwork/sao/x/model/keeper"
)

func TestVersion(t *testing.T) {

	commit := "abc"
	height := int64(10)
	version := keeper.Version(commit, height)
	commitInfo := strings.Split(version, "\032")
	fmt.Println(hex.EncodeToString([]byte("\032")))
	fmt.Println(version)
	if len(commitInfo) != 2 || len(commitInfo[1]) == 0 {
		t.Error("invalid commit information")
	}
	height, err := strconv.ParseInt(commitInfo[1], 10, 64)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(commitInfo[0], height)
}
