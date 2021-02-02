package repository

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/model"
)

// AccountRepository - interface of account repository
type AccountRepository interface {
	Create(ctx context.Context, account model.Account) (*model.Account, error)
	UpdateBalance(ctx context.Context, accountOriginId model.AccountID, balance types.Money) (int64, error)
	FindAll(ctx context.Context) ([]model.Account, error)
	GetBalance(ctx context.Context, accountId string) (*model.Account, error)
	FindByCpf(ctx context.Context, cpf string) (*model.Account, error)
	FindById(ctx context.Context, accountId string) (*model.Account, error)
}
