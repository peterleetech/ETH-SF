package model

import (
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
	"veric-backend-mvp/logic/eth"
)

var issuePryKey *eth.PrivateKey

func init() {
	var err error
	issuePryKey, err = eth.NewPrivateKey("...")
	if err != nil {
		panic(err)
	}
}

func randomUserKey() *eth.PrivateKey {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	return eth.NewPrivateKeyFromECDSA(key)
}

func randomUserDID() (*eth.PrivateKey, *DIDDocument) {
	key := randomUserKey()
	did, err := CreateUserDID(key.PublicKey())
	if err != nil {
		panic(err)
	}

	return key, did
}

func TestCreateUserDID(t *testing.T) {
	_, oriDID := randomUserDID()

	toByte := oriDID.ToJson()

	toDID, err := ParseUserDIDFromJsonStr(toByte)
	if err != nil {
		panic(err)
	}

	if toDID.ID != oriDID.ID {
		panic("toDID.ID != oriDID.ID")
	}
}

func Test222(t *testing.T) {
	key := randomUserKey()
	t.Log(key.HexString())
	signature, err := key.ETHSignature([]byte("sign in"))
	if err != nil {
		panic(err)
	}

	sign, err := eth.GetPublicKeyUseEthSign("sign in", signature)
	if err != nil {
		panic(err)
	}

	if sign.Address().String() != key.Address().String() {
		panic("in")
	}

	t.Log(key.PublicKey().Address().String(), signature)
}
