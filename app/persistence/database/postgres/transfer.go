package postgres

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

type transferRepositoryImpl struct {
	dataBasePool *pgxpool.Pool
}

func NewTransferRepository(dataBasePool *pgxpool.Pool) repository.TransferRepository {
	return &transferRepositoryImpl{
		dataBasePool: dataBasePool,
	}
}

func (repositoryImpl transferRepositoryImpl) Create(ctx context.Context, transfer model.Transfer) (model.Transfer, error) {
	panic("implement me")
}
