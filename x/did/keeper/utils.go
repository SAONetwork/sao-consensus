package keeper

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/tendermint/tendermint/crypto"
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
