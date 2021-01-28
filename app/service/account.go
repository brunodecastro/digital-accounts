package service

import (
	"context"
	"errors"
	"github.com/brunodecastro/digital-accounts/app/common/converter"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
)

type AccountService interface {
	Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (output.CreateAccountOutputVO, error)
	GetAll(ctx context.Context) ([]output.FindAllAccountOutputVO, error)
	GetBalance(ctx context.Context, accountId string) (output.FindAccountBalanceOutputVO, error)
}

type accountServiceImpl struct {
	repository repository.AccountRepository
}

func NewAccountService(repository repository.AccountRepository) AccountService {
	return &accountServiceImpl{
		repository: repository,
	}
}

func (serviceImpl accountServiceImpl) Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (output.CreateAccountOutputVO, error) {
	//ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	//defer cancelFunc()

	accountCreated, err := serviceImpl.repository.Create(ctx, converter.CreateAccountInputVOToModel(accountInputVO))
	if err != nil {
		return output.CreateAccountOutputVO{}, errors.New("error on create accounts")
	}

	return converter.AccountModelToCreateAccountOutputVO(accountCreated), err
}

func (serviceImpl accountServiceImpl) GetAll(ctx context.Context, ) ([]output.FindAllAccountOutputVO, error) {
	accounts, err := serviceImpl.repository.GetAll(ctx)
	if err != nil {
		return []output.FindAllAccountOutputVO{}, errors.New("error listing all accounts")
	}

	return converter.AccountModelToFindAllAccountOutputVO(accounts), err
}

func (serviceImpl accountServiceImpl) GetBalance(ctx context.Context, accountId string) (output.FindAccountBalanceOutputVO, error) {
	account, err := serviceImpl.repository.GetBalance(ctx, accountId)
	if err != nil {
		return output.FindAccountBalanceOutputVO{}, errors.New("error getting account balance")
	}

	return converter.AccountModelToFindAccountBalanceOutputVO(account), err
}
