package types

import (
	"strings"
)

type Caip10 struct {
	Network string
	Chain   string
	Address string
}

func ParseToCaip10(accountId string) Caip10 {
	splited := strings.Split(accountId, ":")
	return Caip10{
		Network: splited[0],
		Chain:   splited[1],
		Address: splited[2],
	}
}

func (c Caip10) ToString() string {
	return c.Network + ":" + c.Chain + ":" + c.Address
}
