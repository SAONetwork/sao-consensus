package keeper

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/tendermint/tendermint/crypto"
	"regexp"
	"strings"
)

func inList(obj string, list []string) bool {
	for _, v := range list {
		if v == obj {
			return true
		}
	}
	return false
}

func inUpdateList(did string, list []*types.AccountAuth) bool {
	for _, v := range list {
		if v.AccountDid == did {
			return true
		}
	}
	return false
}

func getVersionInfo(query string) string {
	// version-id was changed to versionId in the latest did-core spec
	// https://github.com/w3c/did-core/pull/553
	var versionId string
	for _, q := range strings.Split(query, "&") {
		if strings.Contains(q, "versionId") || strings.Contains(q, "version-id") {
			versionId = strings.Split(q, "=")[1]
			break
		}
	}

	return versionId
}

func CalculateDocId(keys []*types.PubKey, timestamp uint64) (string, error) {
	keysmap := make(map[string]string)
	for _, key := range keys {
		keysmap[key.Name] = key.Value
	}

	keysBytes, err := json.Marshal(keysmap)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(crypto.Sha256([]byte(string(keysBytes) + fmt.Sprint(timestamp)))), nil
}

func getSignData(address, message string) []byte {
	// TODO: Amino Sign Doc
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	return []byte(`{"account_number":"0","chain_id":"","fee":{"amount":[],"gas":"0"},"memo":"","msgs":[{"type":"sign/MsgSignData","value":{"data":"` + encodedMessage + `","signer":"` + address + `"}}],"sequence":"0"}`)
}

func parseAcccountId(accountId string) (caip10 types.Caip10, err error) {
	// Check CAIP-10 define
	// check length
	regex := "^[-a-z0-9]{3,8}:[-_a-zA-Z0-9]{1,32}:[-.%a-zA-Z0-9]{1,64}$"
	ok, err := regexp.MatchString(regex, accountId)
	if err != nil {
		return
	}
	if ok {
		caip10 = types.ParseToCaip10(accountId)
		return
	} else {
		return caip10, errors.New("invalid accountId")
	}
}
