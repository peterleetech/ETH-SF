package model

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"
	"veric-backend-mvp/logic/eth"
)

type VerifiableCredentialArr []VerifiableCredential

func (v VerifiableCredentialArr) Len() int {
	return len(v)
}

func (v VerifiableCredentialArr) Less(i, j int) bool {
	return v[i].ID < v[j].ID
}

func (v VerifiableCredentialArr) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v VerifiableCredentialArr) ToByte() []byte {
	var convertedBytes bytes.Buffer
	for _, vc := range v {
		convertedBytes.Write(vc.ToByte())
	}
	return convertedBytes.Bytes()
}

func (v VerifiableCredentialArr) Verify(pubKey *eth.PublicKey) (bool, error) {
	for _, vc := range v {
		verify, err := vc.Verify(pubKey)
		if !verify || err != nil {
			return verify, err
		}
	}

	return true, nil
}

type VerifiablePresentation struct {
	Context              []string                `json:"@context" mapstructure:"@context"`
	Type                 []string                `json:"type"`
	VerifiableCredential VerifiableCredentialArr `json:"verifiableCredential"`
	Holder               string                  `json:"holder"`
	Proof                VPProof                 `json:"proof"`
}

func ParseVerifiablePresentationFromJsonStr(jsonDoc []byte) (*VerifiablePresentation, error) {
	vp := &VerifiablePresentation{}
	err := json.Unmarshal(jsonDoc, vp)
	if err != nil {
		return nil, err
	}
	return vp, nil
}

func (vp *VerifiablePresentation) ToJson() []byte {
	jsonDoc, err := json.Marshal(vp)
	if err != nil {
		return nil
	}

	return jsonDoc
}

func (vp *VerifiablePresentation) ToByte() []byte {
	var convertedBytes bytes.Buffer

	sort.Strings(vp.Context)
	sort.Strings(vp.Type)
	sort.Sort(vp.VerifiableCredential)

	convertedBytes.WriteString(strings.Join(vp.Context, ","))
	convertedBytes.WriteString(strings.Join(vp.Type, ","))
	convertedBytes.Write(vp.VerifiableCredential.ToByte())
	convertedBytes.WriteString(vp.Holder)
	convertedBytes.Write(vp.Proof.ToByte())

	return convertedBytes.Bytes()
}

func (vp *VerifiablePresentation) Signature(holderPrivKey *eth.PrivateKey) error {
	vp.Proof.JWSSignature = ""

	signatureData, err := holderPrivKey.Sha256JWSSignature(vp.ToByte())
	if err != nil {
		return err
	}

	vp.Proof.JWSSignature = signatureData
	return nil
}

func (vp *VerifiablePresentation) Verify(holderPubKey *eth.PublicKey, issuerPubKey *eth.PublicKey) (bool, error) {
	copiedVP := *vp
	copiedVP.Proof.JWSSignature = ""

	isVerifyVP, err := holderPubKey.VerifySha256JWSSignature(copiedVP.ToByte(), vp.Proof.JWSSignature)
	if err != nil {
		return false, err
	}

	if !isVerifyVP {
		return false, nil
	}

	return vp.VerifiableCredential.Verify(issuerPubKey)
}
