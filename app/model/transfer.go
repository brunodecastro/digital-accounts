package model

import (
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"time"
)

type TransferID string

type Transfer struct {
	Id                   TransferID
	AccountOriginId      AccountID
	AccountDestinationId AccountID
	Amount               types.Money
	CreatedAt            time.Time
}
