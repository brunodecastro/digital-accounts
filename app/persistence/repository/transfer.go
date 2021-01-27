package repository

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
)

type TransferRepository interface {
	Create(ctx context.Context, transfer model.Transfer) (model.Transfer, error)
}
