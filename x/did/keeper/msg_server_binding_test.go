package keeper_test

import (
	"encoding/base64"
	"fmt"
	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/did/keeper"
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/cosmos/btcutil/hdkeychain"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/multiformats/go-multibase"
	"github.com/stretchr/testify/require"

	"golang.org/x/crypto/curve25519"
	"testing"
	"time"
)

func TestBinding(t *testing.T) {
	k, ctx := keepertest.DidKeeper(t)
	ctx = ctx.WithChainID("sao")
	server := keeper.NewMsgServerImpl(*k)

	_, _, creator1 := genSecp256k1Account("seed1")
	priv1, pub1, account1 := genSecp256k1Account("seed2")

	kr := NewKeyRing(t)
	pubkeys, err := kr.generatePubKeys()
	if err != nil {
		t.Error(err)
	}

	timestamp := time.Now().Unix()
	ctx = ctx.WithBlockTime(time.Now())
	rootDocId, err := keeper.CalculateDocId(pubkeys, uint64(timestamp))
	if err != nil {
		t.Error(err)
	}

	msg := "Link this account to your did: did:sid:" + rootDocId + "\nTimestamp: " + fmt.Sprint(timestamp)
	signData := keeper.GetSignData(account1, msg)
	sig, err := priv1.Sign(signData)
	if err != nil {
		t.Error(err)
	}

	accountId := "cosmos:sao:" + account1
	did := "did:sid:" + rootDocId
	accountDid := "accountDid1"

	bindingMessage := types.MsgBinding{
		Creator:   creator1,
		AccountId: accountId,
		RootDocId: rootDocId,
		Keys:      pubkeys,
		AccountAuth: &types.AccountAuth{
			AccountDid:           accountDid,
			AccountEncryptedSeed: "aes1",
			SidEncryptedAccount:  "sea1",
		},
		Proof: &types.BindingProof{
			Version: 1,
			Message: msg,
			Signature: "tendermint/PubKeySecp256k1." +
				base64.StdEncoding.EncodeToString(pub1.Bytes()) + "." +
				base64.StdEncoding.EncodeToString(sig),
			Did:       did,
			Timestamp: uint64(timestamp),
		},
	}

	_, err = server.Binding(ctx, &bindingMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrInvalidCreator),
	)

	bindingMessage.Creator = account1

	_, err = server.Binding(ctx, &bindingMessage)
	if err != nil {
		t.Error(err)
	}

	// check storage

	sidDoc, found := k.GetSidDocument(ctx, rootDocId)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(len(pubkeys)),
		nullify.Fill(len(sidDoc.Keys)),
	)
	for i, _ := range pubkeys {
		require.Equal(t,
			nullify.Fill(pubkeys[i]),
			nullify.Fill(sidDoc.Keys[i]),
		)
	}

	sidDocVersion, found := k.GetSidDocumentVersion(ctx, rootDocId)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&[]string{rootDocId}),
		nullify.Fill(&sidDocVersion.VersionList),
	)

	accountAuth, found := k.GetAccountAuth(ctx, accountDid)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&types.AccountAuth{
			AccountDid:           accountDid,
			AccountEncryptedSeed: "aes1",
			SidEncryptedAccount:  "sea1",
		}),
		nullify.Fill(&accountAuth),
	)

	accountList, found := k.GetAccountList(ctx, did)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&[]string{accountDid}),
		nullify.Fill(&accountList.AccountDids),
	)

	storedDid, found := k.GetDid(ctx, accountId)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&did),
		nullify.Fill(&storedDid.Did),
	)

	storedAccountId, found := k.GetAccountId(ctx, accountDid)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&accountId),
		nullify.Fill(&storedAccountId.AccountId),
	)

	payAddr, found := k.GetPaymentAddress(ctx, did)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&account1),
		nullify.Fill(&payAddr.Address),
	)

	// replay binding
	_, err = server.Binding(ctx, &bindingMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrAuthExists),
	)

	// update account auth
	accountDid2 := "accountDid2"
	bindingMessage.AccountAuth = &types.AccountAuth{
		AccountDid:           accountDid2,
		AccountEncryptedSeed: "aes2",
		SidEncryptedAccount:  "sea2",
	}

	_, err = server.Binding(ctx, &bindingMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrBindingExists),
	)

	// update accountId
	priv2, pub2, account2 := genSecp256k1Account("seed3")
	accountId2 := "cosmos:sao:" + account2
	bindingMessage.AccountId = accountId2

	_, err = server.Binding(ctx, &bindingMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrInvalidBindingProof),
	)

	// update binding proof
	timestamp = time.Now().Unix()
	msg2 := "Link this account to your did: " + did + "\nTimestamp: " + fmt.Sprint(timestamp)
	signData2 := keeper.GetSignData(account2, msg2)
	sig2, err := priv2.Sign(signData2)
	if err != nil {
		t.Error(err)
	}

	bindingMessage.Proof = &types.BindingProof{
		Version: 1,
		Message: msg2,
		Signature: "tendermint/PubKeySecp256k1." +
			base64.StdEncoding.EncodeToString(pub2.Bytes()) + "." +
			base64.StdEncoding.EncodeToString(sig2),
		Did:       did,
		Timestamp: uint64(timestamp),
	}

	bindingMessage.Creator = account2

	_, err = server.Binding(ctx, &bindingMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrInvalidCreator),
	)

	// rollback creator to account1
	bindingMessage.Creator = account1
	_, err = server.Binding(ctx, &bindingMessage)
	require.NoError(t, err)

	// check storage

	sidDoc, found = k.GetSidDocument(ctx, rootDocId)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(len(pubkeys)),
		nullify.Fill(len(sidDoc.Keys)),
	)
	for i, _ := range pubkeys {
		require.Equal(t,
			nullify.Fill(pubkeys[i]),
			nullify.Fill(sidDoc.Keys[i]),
		)
	}

	sidDocVersion, found = k.GetSidDocumentVersion(ctx, rootDocId)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&[]string{rootDocId}),
		nullify.Fill(&sidDocVersion.VersionList),
	)

	accountAuth, found = k.GetAccountAuth(ctx, accountDid2)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&types.AccountAuth{
			AccountDid:           accountDid2,
			AccountEncryptedSeed: "aes2",
			SidEncryptedAccount:  "sea2",
		}),
		nullify.Fill(&accountAuth),
	)

	accountList, found = k.GetAccountList(ctx, did)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&[]string{accountDid, accountDid2}),
		nullify.Fill(&accountList.AccountDids),
	)

	storedDid, found = k.GetDid(ctx, accountId2)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&did),
		nullify.Fill(&storedDid.Did),
	)

	storedAccountId, found = k.GetAccountId(ctx, accountDid2)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&accountId2),
		nullify.Fill(&storedAccountId.AccountId),
	)

	payAddr, found = k.GetPaymentAddress(ctx, did)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&account1),
		nullify.Fill(&payAddr.Address),
	)

	bindingMessage.RootDocId = "wrongDocId"
	_, err = server.Binding(ctx, &bindingMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrInconsistentDid),
	)

	bindingMessage.RootDocId = rootDocId
	bindingMessage.Proof.Timestamp = 0
	_, err = server.Binding(ctx, &bindingMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrOutOfDate),
	)

}

type Keys struct {
	Encrypt []byte
	Signing []byte
}

type Keyring struct {
	Seed []byte
	Pub  Keys
	Priv Keys
}

func NewKeyRing(t *testing.T) Keyring {
	seed, err := hdkeychain.GenerateSeed(32)
	if err != nil {
		t.Error(err)
	}
	ek, err := hdkeychain.NewMaster(seed, &chaincfg.TestNet3Params)
	if err != nil {
		t.Error(err)
	}
	signingKey, err := ek.Derive(0)
	if err != nil {
		t.Error(err)
	}
	ecryptionKey, err := ek.Derive(1)
	if err != nil {
		t.Error(err)
	}
	ecryptionPrivKey, err := ecryptionKey.ECPrivKey()
	if err != nil {
		t.Error(err)
	}
	var encryptionPub [32]byte
	encryptionPriv := ecryptionPrivKey.D.Bytes()
	curve25519.ScalarBaseMult(&encryptionPub, (*[32]byte)(encryptionPriv))

	signingPrivKey, err := signingKey.ECPrivKey()
	if err != nil {
		t.Error(err)
	}
	signingPriv := signingPrivKey.D.Bytes()
	signingPubKey, err := signingKey.ECPubKey()
	if err != nil {
		t.Error(err)
	}
	signingPub := signingPubKey.SerializeCompressed()

	return Keyring{
		Seed: seed,
		Pub:  Keys{Encrypt: encryptionPub[:], Signing: signingPub},
		Priv: Keys{Encrypt: encryptionPriv, Signing: signingPriv},
	}
}

func (kr Keyring) generatePubKeys() ([]*types.PubKey, error) {
	signing, err := multibase.Encode(multibase.Base58BTC, append([]byte{0xe7, 0x01}, kr.Pub.Signing...))
	if err != nil {
		return nil, err
	}
	encrypt, err := multibase.Encode(multibase.Base58BTC, append([]byte{0xec, 0x01}, kr.Pub.Encrypt...))
	if err != nil {
		return nil, err
	}
	return []*types.PubKey{{
		Name:  KeyName(signing),
		Value: signing,
	}, {
		Name:  KeyName(encrypt),
		Value: encrypt,
	}}, nil
}

func KeyName(key string) string {
	if len(key) <= 10 {
		panic("key length is too short")
	}

	return key[len(key)-10:]
}

func genSecp256k1Account(seed string) (priv secp256k1.PrivKey, pub cryptotypes.PubKey, addr string) {
	priv = *secp256k1.GenPrivKeyFromSecret([]byte(seed))
	pub = priv.PubKey()
	addr = sdk.MustBech32ifyAddressBytes("cosmos", pub.Address().Bytes())
	return
}
