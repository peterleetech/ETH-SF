package model

import "bytes"

type VPProof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	VerificationMethod string `json:"verificationMethod"`
	ProofPurpose       string `json:"proofPurpose"`
	JWSSignature       string `json:"jws"`   //signature is created from a hash of the VP
	Nonce              string `json:"nonce"` //random value generated by verifier that must be included in proof
}

func (v *VPProof) ToByte() []byte {
	var convertedBytes bytes.Buffer
	convertedBytes.WriteString(v.Type)
	convertedBytes.WriteString(v.Created)
	convertedBytes.WriteString(v.VerificationMethod)
	convertedBytes.WriteString(v.ProofPurpose)
	convertedBytes.WriteString(v.JWSSignature)
	convertedBytes.WriteString(v.Nonce)

	return convertedBytes.Bytes()
}