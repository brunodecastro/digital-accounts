package model

import (
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"time"
)

// TransferID entity transfer ID type
type TransferID string

// Transfer - struct of transfer entity
type Transfer struct {
	ID                   TransferID
	AccountOriginID      AccountID
	AccountDestinationID AccountID
	Amount               types.Money // Brazilian real cents
	CreatedAt            time.Time
}
