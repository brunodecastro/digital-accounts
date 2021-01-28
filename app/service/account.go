package service

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/brunodecastro/digital-accounts/app/util"
	"time"
)

type AccountService interface {
	Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (*model.Account, error)
	GetAll(ctx context.Context) ([]model.Account, error)
	GetBalance(ctx context.Context, accountId string) (*model.Account, error)
}

type accountServiceImpl struct {
	repository repository.AccountRepository
}

func NewAccountService(repository repository.AccountRepository) AccountService {
	return &accountServiceImpl{
		repository: repository,
	}
}

func (serviceImpl accountServiceImpl) Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (*model.Account, error) {
	//ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	//defer cancelFunc()

	return serviceImpl.repository.Create(ctx, accountVOToModel(accountInputVO))
}

func (serviceImpl accountServiceImpl) GetAll(ctx context.Context, ) ([]model.Account, error) {
	return serviceImpl.repository.GetAll(ctx)
}

func (serviceImpl accountServiceImpl) GetBalance(ctx context.Context, accountId string) (*model.Account, error) {
	return serviceImpl.repository.GetBalance(ctx, accountId)
}

func accountVOToModel(accountInputVO input.CreateAccountInputVO) model.Account {
	return model.Account{
		ID:        model.AccountID(common.NewUUID()),
		Name:      accountInputVO.Name,
		Cpf:       util.NumbersOnly(accountInputVO.Cpf),
		Secret:    util.EncryptPassword(accountInputVO.Secret),
		Balance:   common.Money(accountInputVO.Balance),
		CreatedAt: time.Now(),
	}
}
