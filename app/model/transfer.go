package model

import (
	"github.com/brunodecastro/digital-accounts/app/common"
	"time"
)

type TransferID string

type Transfer struct {
	Id                   TransferID
	AccountOriginId      AccountID
	AccountDestinationId AccountID
	Amount               common.Money
	CreatedAt            time.Time
}
