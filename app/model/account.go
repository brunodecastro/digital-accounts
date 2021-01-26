package model

import (
	"github.com/google/uuid"
)

type Account struct {
	ID uuid.UUID
	Name string
	User User
	Amount float32
}