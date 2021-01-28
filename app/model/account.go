package model

import (
	"github.com/brunodecastro/digital-accounts/app/common"
	"time"
)

type AccountID string

type Account struct {
	ID        AccountID    `json:"id"`
	Cpf       string       `json:"cpf"`
	Name      string       `json:"name"`
	Secret    string       `json:"secret"`
	Balance   common.Money `json:"balance"`
	CreatedAt time.Time    `json:"created-at"`
}

func NewAccount(id AccountID, name string, cpf string, balance common.Money, createdAt time.Time) Account {
	return Account{
		ID:        id,
		Name:      name,
		Cpf:       cpf,
		Balance:   balance,
		CreatedAt: createdAt,
	}
}
