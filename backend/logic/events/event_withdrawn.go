package events

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type EventWithdrawn struct {
	VCId         string         `json:"VCId"`
	TokenAddress common.Address `json:"tokenAddress"`
	Amount       *big.Int       `json:"amount"`
	Vault        string         `json:"vault"`
	Requester    common.Address `json:"requester"`
}

func (e *EventWithdrawn) MarshalBinary() (data []byte, err error) {
	return json.Marshal(e)
}

func (e *EventWithdrawn) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, e)
}

func (e *EventWithdrawn) UniqId() string {
	return e.VCId
}
