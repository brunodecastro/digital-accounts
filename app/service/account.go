package service

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"time"
)

type AccountService interface {
	Create(ctx context.Context, account model.Account) (model.Account, error)
	GetAll(ctx context.Context) ([]model.Account, error)
	GetBalance(ctx context.Context, accountId string) (model.Account, error)
}

type accountServiceImpl struct {
	repository repository.AccountRepository
}

func NewAccountService(repository repository.AccountRepository) AccountService {
	return &accountServiceImpl{
		repository: repository,
	}
}

func (serviceImpl accountServiceImpl) Create(ctx context.Context, account model.Account) (model.Account, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()
	return serviceImpl.repository.Create(ctx, account)
}

func (serviceImpl accountServiceImpl) GetAll(ctx context.Context, ) ([]model.Account, error) {
	panic("implement me")
}

func (serviceImpl accountServiceImpl) GetBalance(ctx context.Context, accountId string) (model.Account, error) {
	panic("implement me")
}
