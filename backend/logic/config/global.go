package config

import (
	"veric-backend-mvp/logic/eth"
	"veric-backend-mvp/logic/model"
)

var IssuePriKey *eth.PrivateKey
var IssueDID *model.DIDDocument

func init() {
	var err error
	IssuePriKey, err = eth.NewPrivateKey("...")
	if err != nil {
		panic(err)
	}

	IssueDID, err = model.CreateUserDID(IssuePriKey.PublicKey())
}
