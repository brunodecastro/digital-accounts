package model


import (
	"github.com/brunodecastro/digital-accounts/app/model/constants"
	"github.com/google/uuid"
)

type Transaction struct {
	ID uuid.UUID
	Amount float32
	TransactionType constants.TransactionType
}