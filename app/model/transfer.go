package model

import (
	"time"
)

type TransferID string

type Transfer struct {
	ID                   TransferID `json:"id"`
	Amount               int        `json:"amount"`
	AccountOriginID      string     `json:"account_origin_id"`
	AccountDestinationID string     `json:"account_destination_id"`
	CreateAt             time.Time  `json:"created-at"`
}
