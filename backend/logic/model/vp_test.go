package model

import (
	"log"
	"testing"
	"veric-backend-mvp/logic/eth"
)

func TestName(t *testing.T) {
	jsonData := `{"@context":["https://ns.did.ai/suites/secp256k1-2019/v1/","https://www.w3.org/2018/credentials/v1"],"type":["VerifiablePresentation"],"verifiableCredential":[{"@context":["https://ns.did.ai/suites/secp256k1-2019/v1/","https://www.w3.org/2018/credentials/v1"],"id":"0","type":["VericDeposit","VerifiableCredential"],"issuer":"did:veric:0xC5BCf228F28a1827Da6C7e576b6d0Dfa5A5168Be","issuanceDate":"2022-06-07T09:43:27Z","expirationDate":"2032-06-07T09:43:27Z","description":"Veric Deposit","credentialSubject":{"tokenAddress":"0xc4860463c59d59a9afac9fde35dff9da363e8425","amount":1000000000000000000,"vault":"0xd3446851deb19bcf700dadef258ba90834c8472a"},"proof":{"type":"EcdsaSecp256k1Signature2019","created":"2022-06-07T09:43:27Z","verificationMethod":"did:veric:0xC5BCf228F28a1827Da6C7e576b6d0Dfa5A5168Be#verification","proofPurpose":"Authentication","jws":"eyJhbGciOiJFUzI1NiJ9..sUOiidTUHzNJ_L93k25EETrzfGcpeqXkDOFxBs2Q4sEtsxRri-Ah5JtCEvQKGNLQYFK2WqIbrmmhbYDggcREGQ"}}],"holder":"did:veric:0x77CBcc0e29E10F1EeA24e0D109aaB26C5b2Abd88","proof":{"type":"EcdsaSecp256k1Signature2019","created":"2022-06-07T18:23:41+08:00","verificationMethod":"did:veric:0x77CBcc0e29E10F1EeA24e0D109aaB26C5b2Abd88#verification","proofPurpose":"Authentication","jws":"eyJhbGciOiJFUzI1NiJ9..Oge0fus9C__fsEk8eUVYZgu47co5aFI8mvVjcjcvc8cUjjor8yEpuWhVjjtO-l0cLxUKTQRWwx0TwCDzulfk2A","nonce":"6666"}}`
	str, err := ParseVerifiablePresentationFromJsonStr([]byte(jsonData))
	if err != nil {
		panic(err)
	}

	key, err := eth.NewPrivateKey("...")
	if err != nil {
		panic(err)
	}

	IssuePriKey, err := eth.NewPrivateKey("...")
	if err != nil {
		panic(err)
	}

	sign, err := eth.GetPublicKeyUseEthSign("sign in", "0x173dd86d64b3ea14cce5a1507256174266b4f0ebfc2ae6db074d4121f60c3f967e8d544ddb36a1aea0bab238ee8b5158edd652049407bdcfb8aa803b25df151d1c")
	if err != nil {
		panic(err)
	}

	t.Log(sign.Address().String(), key.Address().String())
	t.Log(IssuePriKey.Address().String())

	//err = str.Signature(key)
	//if err != nil {
	//	panic(err)
	//}

	verify, err := str.Verify(sign, IssuePriKey.PublicKey())
	if err != nil {
		panic(err)
	}

	log.Println(verify)
}
