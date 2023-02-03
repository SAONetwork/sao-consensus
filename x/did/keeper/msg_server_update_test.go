package keeper_test

import (
	"encoding/base64"
	"fmt"
	keepertest "github.com/SaoNetwork/sao/testutil/keeper"
	"github.com/SaoNetwork/sao/testutil/nullify"
	"github.com/SaoNetwork/sao/x/did/keeper"
	"github.com/SaoNetwork/sao/x/did/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {

	k, ctx := keepertest.DidKeeper(t)
	ctx = ctx.WithChainID("sao")
	server := keeper.NewMsgServerImpl(*k)

	priv1, pub1, account1 := genSecp256k1Account("seed1")
	priv2, pub2, account2 := genSecp256k1Account("seed2")
	priv3, pub3, account3 := genSecp256k1Account("seed3")
	priv4, pub4, account4 := genSecp256k1Account("seed4")

	kr := NewKeyRing(t)
	pubkeys, err := kr.generatePubKeys()
	require.NoError(t, err)

	timestamp := time.Now().Unix()
	ctx = ctx.WithBlockTime(time.Now())
	rootDocId, err := keeper.CalculateDocId(pubkeys, uint64(timestamp))
	require.NoError(t, err)

	msg := "Link this account to your did: did:sid:" + rootDocId + "\nTimestamp: " + fmt.Sprint(timestamp)
	signData1 := keeper.GetSignData(account1, msg)
	sig1, err := priv1.Sign(signData1)
	require.NoError(t, err)
	signature1 := "tendermint/PubKeySecp256k1." +
		base64.StdEncoding.EncodeToString(pub1.Bytes()) + "." +
		base64.StdEncoding.EncodeToString(sig1)
	signData2 := keeper.GetSignData(account2, msg)
	sig2, err := priv2.Sign(signData2)
	require.NoError(t, err)
	signature2 := "tendermint/PubKeySecp256k1." +
		base64.StdEncoding.EncodeToString(pub2.Bytes()) + "." +
		base64.StdEncoding.EncodeToString(sig2)
	signData3 := keeper.GetSignData(account3, msg)
	sig3, err := priv3.Sign(signData3)
	require.NoError(t, err)
	signature3 := "tendermint/PubKeySecp256k1." +
		base64.StdEncoding.EncodeToString(pub3.Bytes()) + "." +
		base64.StdEncoding.EncodeToString(sig3)
	signData4 := keeper.GetSignData(account4, msg)
	sig4, err := priv4.Sign(signData4)
	require.NoError(t, err)
	signature4 := "tendermint/PubKeySecp256k1." +
		base64.StdEncoding.EncodeToString(pub4.Bytes()) + "." +
		base64.StdEncoding.EncodeToString(sig4)

	accountId1 := "cosmos:sao:" + account1
	accountId2 := "cosmos:sao:" + account2
	accountId3 := "cosmos:sao:" + account3
	accountId4 := "cosmos:sao:" + account4
	did := "did:sid:" + rootDocId
	accountDid1 := "accountDid1"
	accountDid2 := "accountDid2"
	accountDid3 := "accountDid3"
	accountDid4 := "accountDid4"
	accountAuth1 := types.AccountAuth{
		AccountDid:           accountDid1,
		AccountEncryptedSeed: "aes1",
		SidEncryptedAccount:  "sea1",
	}
	accountAuth2 := types.AccountAuth{
		AccountDid:           accountDid2,
		AccountEncryptedSeed: "aes2",
		SidEncryptedAccount:  "sea2",
	}
	accountAuth3 := types.AccountAuth{
		AccountDid:           accountDid3,
		AccountEncryptedSeed: "aes3",
		SidEncryptedAccount:  "sea3",
	}
	accountAuth4 := types.AccountAuth{
		AccountDid:           accountDid4,
		AccountEncryptedSeed: "aes4",
		SidEncryptedAccount:  "sea4",
	}

	bindingMessage1 := types.MsgBinding{
		Creator:     account1,
		AccountId:   accountId1,
		RootDocId:   rootDocId,
		Keys:        pubkeys,
		AccountAuth: &accountAuth1,
		Proof: &types.BindingProof{
			Version:   1,
			Message:   msg,
			Signature: signature1,
			Did:       did,
			Timestamp: uint64(timestamp),
		},
	}
	bindingMessage2 := types.MsgBinding{
		Creator:     account1,
		AccountId:   accountId2,
		RootDocId:   rootDocId,
		Keys:        pubkeys,
		AccountAuth: &accountAuth2,
		Proof: &types.BindingProof{
			Version:   1,
			Message:   msg,
			Signature: signature2,
			Did:       did,
			Timestamp: uint64(timestamp),
		},
	}
	bindingMessage3 := types.MsgBinding{
		Creator:     account1,
		AccountId:   accountId3,
		RootDocId:   rootDocId,
		Keys:        pubkeys,
		AccountAuth: &accountAuth3,
		Proof: &types.BindingProof{
			Version:   1,
			Message:   msg,
			Signature: signature3,
			Did:       did,
			Timestamp: uint64(timestamp),
		},
	}
	bindingMessage4 := types.MsgBinding{
		Creator:     account1,
		AccountId:   accountId4,
		RootDocId:   rootDocId,
		Keys:        pubkeys,
		AccountAuth: &accountAuth4,
		Proof: &types.BindingProof{
			Version:   1,
			Message:   msg,
			Signature: signature4,
			Did:       did,
			Timestamp: uint64(timestamp),
		},
	}

	_, err = server.Binding(ctx, &bindingMessage1)
	require.NoError(t, err)
	_, err = server.Binding(ctx, &bindingMessage2)
	require.NoError(t, err)
	_, err = server.Binding(ctx, &bindingMessage3)
	require.NoError(t, err)
	_, err = server.Binding(ctx, &bindingMessage4)
	require.NoError(t, err)

	kr1 := NewKeyRing(t)
	pubkeys1, err := kr1.generatePubKeys()
	require.NoError(t, err)

	timestamp1 := time.Now().Unix()
	ctx = ctx.WithBlockTime(time.Now())
	docId1, err := keeper.CalculateDocId(pubkeys1, uint64(timestamp1))
	require.NoError(t, err)

	updateAccountAuth := []*types.AccountAuth{
		{
			AccountDid:           accountDid1,
			AccountEncryptedSeed: "aes1.1",
			SidEncryptedAccount:  "sea1.1",
		},
		{
			AccountDid:           accountDid2,
			AccountEncryptedSeed: "aes2.1",
			SidEncryptedAccount:  "sea2.1",
		},
		{
			AccountDid:           accountDid3,
			AccountEncryptedSeed: "aes3.1",
			SidEncryptedAccount:  "sea3.1",
		}}

	pastSeed1 := "pastSeed1"
	updateMessagePre := types.MsgUpdate{
		Creator:           account4,
		Did:               did,
		NewDocId:          docId1,
		Keys:              pubkeys1,
		Timestamp:         uint64(timestamp1),
		UpdateAccountAuth: updateAccountAuth,
		RemoveAccountDid:  []string{accountDid4},
		PastSeed:          pastSeed1,
	}
	_, err = server.Update(ctx, &updateMessagePre)
	require.NoError(t, err)

	updateMessage := types.MsgUpdate{
		Creator:           "WrongCreator",
		Did:               "WrongDid",
		NewDocId:          docId1,
		Keys:              nil,
		Timestamp:         0,
		UpdateAccountAuth: nil,
		RemoveAccountDid:  nil,
		PastSeed:          pastSeed1,
	}

	_, err = server.Update(ctx, &updateMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrInvalidCreator),
	)

	updateMessage.Creator = account3
	updateMessage.Did = did
	_, err = server.Update(ctx, &updateMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrOutOfDate),
	)

	updateMessage.Timestamp = uint64(timestamp)
	_, err = server.Update(ctx, &updateMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrNoNeedToUpdate),
	)

	updateMessage.RemoveAccountDid = []string{"WrongAccountDid"}
	_, err = server.Update(ctx, &updateMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrUpdateAccAuthEmpty),
	)

	updateMessage.UpdateAccountAuth = []*types.AccountAuth{{
		AccountDid:           "WrongAccountDid",
		AccountEncryptedSeed: "aaa",
		SidEncryptedAccount:  "bbb",
	}}
	_, err = server.Update(ctx, &updateMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrInvalidAuthCount),
	)

	updateMessage.RemoveAccountDid = []string{"WrongAccountDid", "WrongAccountDid"}
	_, err = server.Update(ctx, &updateMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrUnhandledAccountDid),
	)

	updateMessage.RemoveAccountDid = []string{accountDid1, accountDid2}
	updateMessage.UpdateAccountAuth = []*types.AccountAuth{{
		AccountDid:           accountDid3,
		AccountEncryptedSeed: "aes3.2",
		SidEncryptedAccount:  "sea3.2",
	}}
	_, err = server.Update(ctx, &updateMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrSeedExists),
	)

	pastSeed2 := "pastSeed2"
	updateMessage.PastSeed = pastSeed2
	_, err = server.Update(ctx, &updateMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrUnbindPayAddr),
	)

	updateMessage.RemoveAccountDid = []string{accountDid2}
	updateMessage.UpdateAccountAuth = []*types.AccountAuth{{
		AccountDid:           accountDid1,
		AccountEncryptedSeed: "aes1.2",
		SidEncryptedAccount:  "sea1.2",
	},
		{
			AccountDid:           accountDid3,
			AccountEncryptedSeed: "aes3.2",
			SidEncryptedAccount:  "sea3.2",
		}}
	_, err = server.Update(ctx, &updateMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrDocExists),
	)

	updateMessage.NewDocId = "newDocId"
	_, err = server.Update(ctx, &updateMessage)
	require.Equal(t,
		nullify.Fill(err),
		nullify.Fill(types.ErrInconsistentDocId),
	)

	kr2 := NewKeyRing(t)
	pubkeys2, err := kr2.generatePubKeys()
	require.NoError(t, err)
	timestamp2 := time.Now().Unix()
	ctx = ctx.WithBlockTime(time.Now())
	docId2, err := keeper.CalculateDocId(pubkeys2, uint64(timestamp2))
	require.NoError(t, err)
	updateMessage.NewDocId = docId2
	updateMessage.Keys = pubkeys2
	_, err = server.Update(ctx, &updateMessage)
	require.NoError(t, err)

	// check storage
	_, found := k.GetDid(ctx, accountId2)
	require.False(t, found)
	_, found = k.GetDid(ctx, accountId4)
	require.False(t, found)
	_, found = k.GetAccountId(ctx, accountDid2)
	require.False(t, found)
	_, found = k.GetAccountId(ctx, accountDid4)
	require.False(t, found)

	sidDoc, found := k.GetSidDocument(ctx, rootDocId)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(len(pubkeys)),
		nullify.Fill(len(sidDoc.Keys)),
	)
	for i := range pubkeys {
		require.Equal(t,
			nullify.Fill(pubkeys[i]),
			nullify.Fill(sidDoc.Keys[i]),
		)
	}
	sidDoc, found = k.GetSidDocument(ctx, docId1)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(len(pubkeys1)),
		nullify.Fill(len(sidDoc.Keys)),
	)
	for i := range pubkeys1 {
		require.Equal(t,
			nullify.Fill(pubkeys1[i]),
			nullify.Fill(sidDoc.Keys[i]),
		)
	}
	sidDoc, found = k.GetSidDocument(ctx, docId2)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(len(pubkeys2)),
		nullify.Fill(len(sidDoc.Keys)),
	)
	for i := range pubkeys2 {
		require.Equal(t,
			nullify.Fill(pubkeys2[i]),
			nullify.Fill(sidDoc.Keys[i]),
		)
	}

	sidDocVersion, found := k.GetSidDocumentVersion(ctx, rootDocId)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&[]string{rootDocId, docId1, docId2}),
		nullify.Fill(&sidDocVersion.VersionList),
	)

	accountAuth, found := k.GetAccountAuth(ctx, accountDid1)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&types.AccountAuth{
			AccountDid:           accountDid1,
			AccountEncryptedSeed: "aes1.2",
			SidEncryptedAccount:  "sea1.2",
		}),
		nullify.Fill(&accountAuth),
	)
	_, found = k.GetAccountAuth(ctx, accountDid2)
	require.False(t, found)
	accountAuth, found = k.GetAccountAuth(ctx, accountDid3)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&types.AccountAuth{
			AccountDid:           accountDid3,
			AccountEncryptedSeed: "aes3.2",
			SidEncryptedAccount:  "sea3.2",
		}),
		nullify.Fill(&accountAuth),
	)
	_, found = k.GetAccountAuth(ctx, accountDid4)
	require.False(t, found)

	accountList, found := k.GetAccountList(ctx, did)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&[]string{accountDid1, accountDid3}),
		nullify.Fill(&accountList.AccountDids),
	)

	pastSeeds, found := k.GetPastSeeds(ctx, did)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&[]string{pastSeed1, pastSeed2}),
		nullify.Fill(&pastSeeds.Seeds),
	)
}
