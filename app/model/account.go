package model

import (
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"time"
)

// AccountID entity account ID type
type AccountID string

// Account - struct of account entity
type Account struct {
	ID        AccountID
	CPF       string
	Name      string
	Secret    string
	Balance   types.Money // Brazilian real cents
	CreatedAt time.Time
}
