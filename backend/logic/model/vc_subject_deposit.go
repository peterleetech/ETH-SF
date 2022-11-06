package model

import (
	"bytes"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type VCSubjectDeposit struct {
	TokenAddress common.Address `json:"tokenAddress"`
	Amount       *big.Int       `json:"amount"`
	Vault        string         `json:"vault"`
}

func ParseVCSubjectDepositJsonStr(jsonDoc []byte) (*VCSubjectDeposit, error) {
	vc := &VCSubjectDeposit{}
	err := json.Unmarshal(jsonDoc, vc)
	if err != nil {
		return nil, err
	}
	return vc, nil
}

func (v *VCSubjectDeposit) ToJson() []byte {
	jsonDoc, err := json.Marshal(v)
	if err != nil {
		return nil
	}

	return jsonDoc
}

func (v *VCSubjectDeposit) ToByte() []byte {
	var convertedBytes bytes.Buffer

	convertedBytes.WriteString(v.TokenAddress.String())
	convertedBytes.WriteString(v.Amount.String())
	convertedBytes.WriteString(v.Vault)

	return convertedBytes.Bytes()
}
