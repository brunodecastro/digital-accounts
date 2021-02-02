package repository

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/model"
)

// AccountRepository - interface of account repository
type AccountRepository interface {
	Create(ctx context.Context, account model.Account) (*model.Account, error)
	UpdateBalance(ctx context.Context, accountID model.AccountID, balance types.Money) (int64, error)
	FindAll(ctx context.Context) ([]model.Account, error)
	GetBalance(ctx context.Context, accountID string) (*model.Account, error)
	FindByCPF(ctx context.Context, cpf string) (*model.Account, error)
	FindByID(ctx context.Context, accountID string) (*model.Account, error)
}
