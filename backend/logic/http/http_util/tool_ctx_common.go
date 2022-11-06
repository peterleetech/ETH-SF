package http_util

import (
	lru "github.com/hashicorp/golang-lru"
	"veric-backend-mvp/logic/eth"
)

const SignContent = "sign in"

var userCache, _ = lru.New(100)

func GetUserPublicKeyFromSign(sign string) (*eth.PublicKey, error) {
	if cached, ok := userCache.Get(sign); ok {
		return cached.(*eth.PublicKey), nil
	} else {
		publicKey, err := eth.GetPublicKeyUseEthSign(SignContent, sign)
		if err != nil {
			return nil, err
		}

		userCache.Add(sign, publicKey)
		return publicKey, nil
	}
}
