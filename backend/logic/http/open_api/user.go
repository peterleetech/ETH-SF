package open_api

import (
	"log"
	"veric-backend-mvp/logic/http/http_util"
)

type LoginUseEthSignatureRequest struct {
	EthAddress string `json:"eth_address" validate:"required"`
	SignData   string `json:"sign_data" validate:"required"`
}

func LoginUseEthSignature(typ interface{}, r *http_util.HTTPContext) (resp interface{}, err error) {
	req := typ.(*LoginUseEthSignatureRequest)
	pk, err := http_util.GetUserPublicKeyFromSign(req.SignData)
	if err != nil {
		return nil, http_util.NewHttpError(0xE001001, "sign data is invalid.")
	}
	if pk.Address().String() != req.EthAddress {
		log.Println(pk.Address().String(), req.EthAddress)
		return nil, http_util.NewHttpError(0xE001002, "sign address is invalid.")
	}

	return nil, err
}
