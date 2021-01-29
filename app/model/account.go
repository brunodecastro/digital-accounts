package model

import (
	"github.com/brunodecastro/digital-accounts/app/common"
	"time"
)

type AccountID string

type Account struct {
	Id        AccountID
	Cpf       string
	Name      string
	Secret    string
	Balance   common.Money
	CreatedAt time.Time
}

func NewAccount(id AccountID, name string, cpf string, balance common.Money, createdAt time.Time) Account {
	return Account{
		Id:        id,
		Name:      name,
		Cpf:       cpf,
		Balance:   balance,
		CreatedAt: createdAt,
	}
}
