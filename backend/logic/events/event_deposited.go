package events

import (
	"encoding/base64"
	"encoding/json"
	"veric-backend-mvp/logic/sol/vault"
)

type EventDeposited struct {
	Deposited *vault.SolDeposited
}

func (e *EventDeposited) MarshalBinary() (data []byte, err error) {
	return json.Marshal(e.Deposited)
}

func (e *EventDeposited) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &e.Deposited)
}

func (e *EventDeposited) UniqId() string {
	return base64.RawStdEncoding.EncodeToString(e.Deposited.Raw.TxHash.Bytes())
}
