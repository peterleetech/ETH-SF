package model

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"sort"
	"strings"
	"time"
	"veric-backend-mvp/logic/eth"
)

type VerifiableCredential struct {
	Context           []string        `json:"@context" mapstructure:"@context"`
	ID                string          `json:"id"`
	Type              []string        `json:"type"`
	Issuer            string          `json:"issuer"`
	IssuanceDate      string          `json:"issuanceDate"`
	ExpirationDate    string          `json:"expirationDate"`
	Description       string          `json:"description"`
	CredentialSubject json.RawMessage `json:"credentialSubject"`
	Proof             VCProof         `json:"proof"`
}

func CreateVerifiableCredential(issueDID *DIDDocument, credentialSubject CanToJson) *VerifiableCredential {
	documentId := uuid.NewString()
	loc, _ := time.LoadLocation("UTC")

	return &VerifiableCredential{
		ID:                documentId,
		Context:           []string{ContextCredential, ContextSecp256k1},
		Type:              []string{TypeCredential, TypeVericDeposit},
		Issuer:            issueDID.ID,
		IssuanceDate:      time.Now().In(loc).Format(time.RFC3339),
		ExpirationDate:    time.Now().In(loc).AddDate(10, 0, 0).Format(time.RFC3339),
		Description:       "Veric Deposit",
		CredentialSubject: credentialSubject.ToJson(),
		Proof:             CreateProof(issueDID.Authentication),
	}
}

func ParseVerifiableCredentialFromJsonStr(jsonDoc []byte) (*VerifiableCredential, error) {
	vc := &VerifiableCredential{}
	err := json.Unmarshal(jsonDoc, vc)
	if err != nil {
		return nil, err
	}
	return vc, nil
}

func (vc *VerifiableCredential) ToJson() []byte {
	jsonDoc, err := json.Marshal(vc)
	if err != nil {
		return nil
	}

	return jsonDoc
}

func (vc *VerifiableCredential) ToByte() []byte {
	var convertedBytes bytes.Buffer

	sort.Strings(vc.Context)
	sort.Strings(vc.Type)

	convertedBytes.WriteString(strings.Join(vc.Context, ","))
	convertedBytes.WriteString(vc.ID)
	convertedBytes.WriteString(strings.Join(vc.Type, ","))
	convertedBytes.WriteString(vc.Issuer)
	convertedBytes.WriteString(vc.IssuanceDate)
	convertedBytes.WriteString(vc.ExpirationDate)
	convertedBytes.WriteString(vc.Description)
	convertedBytes.Write(vc.CredentialSubject)
	convertedBytes.Write(vc.Proof.ToByte())

	return convertedBytes.Bytes()
}

func (vc *VerifiableCredential) Signature(issuerPrivKey *eth.PrivateKey) error {
	vc.Proof.JWSSignature = ""

	signatureData, err := issuerPrivKey.Sha256JWSSignature(vc.ToByte())
	if err != nil {
		return err
	}

	vc.Proof.JWSSignature = signatureData
	return nil
}

func (vc *VerifiableCredential) Verify(issuerPubKey *eth.PublicKey) (bool, error) {
	copiedVC := *vc
	copiedVC.Proof.JWSSignature = ""

	result, err := issuerPubKey.VerifySha256JWSSignature(copiedVC.ToByte(), vc.Proof.JWSSignature)
	if err != nil {
		return false, err
	}
	return result, nil
}
