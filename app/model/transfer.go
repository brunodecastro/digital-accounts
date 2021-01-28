package model

import (
	"github.com/google/uuid"
	"time"
)

type Transfer struct {
	ID                   uuid.UUID    `json:"id"`
	Amount               int       `json:"amount"`
	AccountOriginID      string    `json:"account_origin_id"`
	AccountDestinationID string    `json:"account_destination_id"`
	CreateAt             time.Time `json:"created-at"`
}
