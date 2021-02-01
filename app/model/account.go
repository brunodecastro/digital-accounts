package model

import (
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"time"
)

type AccountID string

type Account struct {
	Id        AccountID
	Cpf       string
	Name      string
	Secret    string
	Balance   types.Money // Brazilian real cents
	CreatedAt time.Time
}
