package open_api

import (
	"veric-backend-mvp/logic/db"
	"veric-backend-mvp/logic/events"
	"veric-backend-mvp/logic/http/http_util"
	"veric-backend-mvp/logic/model"
)

var event *events.Events

func SetEvent(e *events.Events) {
	event = e
}

func FindUser(r *http_util.HTTPContext) (*db.User, error) {
	pk, err := r.GetHeaderPublicKey()
	if err != nil {
		return nil, err
	}

	user, err := db.FindUserByAddress(pk.Address().String())
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		did, err := model.CreateUserDID(pk)
		if err != nil {
			return nil, err
		}
		user = &db.User{
			Address: pk.Address().String(),
			DID:     string(did.ToJson()),
			DIDId:   did.ID,
		}
		return user, db.SaveUser(user)
	}

	if user.DID == "" {
		did, err := model.CreateUserDID(pk)
		if err != nil {
			return nil, err
		}
		user.DID = string(did.ToJson())
		user.DIDId = did.ID
		return user, db.SaveUser(user)
	}

	return user, err
}

func RecordEvent() {

}
