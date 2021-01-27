package model

import (
	"time"
)

type Account struct {
	ID       string    `json:"id"`
	Cpf      string    `json:"cpf"`
	Name     string    `json:"name"`
	Secret   string    `json:"-"`
	Balance  int       `json:"balance"`
	CreateAt time.Time `json:"created-at"`
}
