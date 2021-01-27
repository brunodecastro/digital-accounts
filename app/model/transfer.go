package model

import (
	"time"
)

type Transfer struct {
	ID                   string    `json:"id"`
	Amount               int       `json:"amount"`
	AccountOriginID      string    `json:"account_origin_id"`
	AccountDestinationID string    `json:"account_destination_id"`
	CreateAt             time.Time `json:"created-at"`
}
