package types

import (
	"github.com/dvsekhvalnov/jose2go/base64url"
	"testing"
)

// use this way to rebuild payload bytes
func TestProposalEncodeDecode(t *testing.T) {
	proposalbase64 := "CkhkaWQ6c2lkOmMxOTc2YjAzMDhkNjY2MzNjMDE5YzU1NzJlM2Q2YmZlM2M3YzM4Y2Q2MTc2NDgwMjAzNDJlYWVmMTMyZDc2MmMSLWNvc21vczEza2p6cDd4ZDlkeHV6YTV0cGFyZGgzamNmNHYyOGo2dG10eHRzdxokMzAyOTNmMGYtM2UwZi00YjNjLWFmZjEtODkwYTJmZGYwNjNiIO0CKAEwgKMFOgp0ZXN0X21vZGVsQiQxMTA5MjhhOC05MzZlLTQ0ZDMtOWYwNS1kZmEwYzJkOWUyNmNKJDExMDkyOGE4LTkzNmUtNDRkMy05ZjA1LWRmYTBjMmQ5ZTI2Y1o7YmFma3JlaWZ4a3Bna29hNXNvcmxudmI2dHVwdGE0N29zNHE1Z2t3NnJlNTU1dGF2Z3BycDR4c3FmeDRwC3gB"
	proposalbytes, err := base64url.Decode(proposalbase64)
	if err != nil {
		t.Error(err)
	}
	var p Proposal
	err = p.Unmarshal(proposalbytes)
	if err != nil {
		t.Error(err)
	}
	bytes, err := p.Marshal()
	if err != nil {
		t.Error(err)
	}
	base64 := base64url.Encode(bytes)
	if base64 != proposalbase64 {
		t.Error("go encoded", base64, "js encoded", proposalbase64)
	}
}
