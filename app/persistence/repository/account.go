package repository

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
)

type AccountRepository interface {
	Create(ctx context.Context, account model.Account) (*model.Account, error)
	FindAll(ctx context.Context) ([]model.Account, error)
	GetBalance(ctx context.Context, accountId string) (*model.Account, error)
	FindByCpf(ctx context.Context, cpf string) (*model.Account, error)
}
