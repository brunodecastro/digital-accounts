package repository

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
)

// TransferRepository - interface of transfer repository
type TransferRepository interface {
	Create(ctx context.Context, transfer model.Transfer) (*model.Transfer, error)
	FindAll(ctx context.Context, accountId string) ([]model.Transfer, error)
}
