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
	Balance   types.Money
	CreatedAt time.Time
}

func NewAccount(id AccountID, name string, cpf string, balance types.Money, createdAt time.Time) Account {
	return Account{
		Id:        id,
		Name:      name,
		Cpf:       cpf,
		Balance:   balance,
		CreatedAt: createdAt,
	}
}
