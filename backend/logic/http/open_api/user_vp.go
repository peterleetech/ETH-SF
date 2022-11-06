package open_api

import (
	"veric-backend-mvp/logic/config"
	"veric-backend-mvp/logic/events"
	"veric-backend-mvp/logic/http/http_util"
	"veric-backend-mvp/logic/model"
)

type VPWithdrawRequest struct {
	VP string `json:"vp" validate:"required"`
}

func VPWithdraw(typ interface{}, r *http_util.HTTPContext) (resp interface{}, err error) {
	pk, err := r.GetHeaderPublicKey()
	if err != nil {
		return nil, err
	}

	user, err := FindUser(r)
	if err != nil {
		return nil, err
	}

	_ = user

	req := typ.(*VPWithdrawRequest)
	vp, err := model.ParseVerifiablePresentationFromJsonStr([]byte(req.VP))
	if err != nil {
		return nil, err
	}

	verify, err := vp.Verify(pk, config.IssuePriKey.PublicKey())
	if err != nil {
		return nil, err
	}

	for _, verifiableCredential := range vp.VerifiableCredential {
		subject, err := model.ParseVCSubjectDepositJsonStr(verifiableCredential.CredentialSubject)
		if err != nil {
			return nil, err
		}

		err = event.SubmitWithdrawn(&events.EventWithdrawn{
			VCId:         verifiableCredential.ID,
			TokenAddress: subject.TokenAddress,
			Amount:       subject.Amount,
			Vault:        subject.Vault,
			Requester:    pk.Address(),
		})
		if err != nil {
			return nil, err
		}
	}

	if verify {
		return "mock Withdraw ok", nil
	} else {
		return "mock Withdraw fail", nil
	}
}
