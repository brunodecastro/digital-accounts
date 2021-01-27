package postgres

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

type accountRepositoryImpl struct {
	dataBasePool *pgxpool.Pool
}

func NewAccountRepository(dataBasePool *pgxpool.Pool) repository.AccountRepository {
	return &accountRepositoryImpl{
		dataBasePool: dataBasePool,
	}
}

func (repositoryImpl accountRepositoryImpl) Create(ctx context.Context, account model.Account) (model.Account, error) {
	panic("implement me")
}

func (repositoryImpl accountRepositoryImpl) GetAll(ctx context.Context) ([]model.Account, error) {
	panic("implement me")
}

func (repositoryImpl accountRepositoryImpl) GetBalance(ctx context.Context, accountId string) (model.Account, error) {
	panic("implement me")
}
