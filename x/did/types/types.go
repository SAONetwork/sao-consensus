package types

import (
	"strings"
)

type Caip10AccountId struct {
	Network string
	Chain   string
	Address string
}

func ParseToCaip10(accountId string) Caip10AccountId {
	splited := strings.Split(accountId, ":")
	return Caip10AccountId{
		Network: splited[0],
		Chain:   splited[1],
		Address: splited[2],
	}
}

func (c Caip10AccountId) ToString() string {
	return c.Network + ":" + c.Chain + ":" + c.Address
}
