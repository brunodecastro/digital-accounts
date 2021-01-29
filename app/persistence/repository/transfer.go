package repository

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
)

type TransferRepository interface {
	Create(ctx context.Context, transfer model.Transfer) (*model.Transfer, error)
	FindAll(ctx context.Context) ([]model.Transfer, error)
}
