package service

import (
	"context"
	"errors"
	"github.com/brunodecastro/digital-accounts/app/common/converter"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
)

type AccountService interface {
	Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (output.CreateAccountOutputVO, error)
	FindAll(ctx context.Context) ([]output.FindAllAccountOutputVO, error)
	GetBalance(ctx context.Context, accountId string) (output.FindAccountBalanceOutputVO, error)
}

type accountServiceImpl struct {
	repository        repository.AccountRepository
	transactionHelper postgres.TransactionHelper
}

func NewAccountService(repository repository.AccountRepository,
	transactionHelper postgres.TransactionHelper) AccountService {
	return &accountServiceImpl{
		repository:        repository,
		transactionHelper: transactionHelper,
	}
}

func (serviceImpl accountServiceImpl) Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (output.CreateAccountOutputVO, error) {
	ctx, err := serviceImpl.transactionHelper.StartTransaction(ctx)
	if err != nil {
		return output.CreateAccountOutputVO{}, custom_errors.ErrorStartTransaction
	}

	accountCreated, err := serviceImpl.repository.Create(ctx, converter.CreateAccountInputVOToModel(accountInputVO))
	if err != nil {
		errT := serviceImpl.transactionHelper.RollbackTransaction(ctx)
		if errT != nil {
			return output.CreateAccountOutputVO{}, custom_errors.ErrorRollbackTransaction
		}
		return output.CreateAccountOutputVO{}, errors.New("error on create accounts")
	}

	serviceImpl.transactionHelper.CommitTransaction(ctx)
	if err != nil {
		return output.CreateAccountOutputVO{}, custom_errors.ErrorCommitTransaction
	}

	return converter.ModelToCreateAccountOutputVO(accountCreated), nil
}

func (serviceImpl accountServiceImpl) FindAll(ctx context.Context, ) ([]output.FindAllAccountOutputVO, error) {
	accounts, err := serviceImpl.repository.FindAll(ctx)
	if err != nil {
		return []output.FindAllAccountOutputVO{}, errors.New("error listing all accounts")
	}

	return converter.AccountModelToFindAllAccountOutputVO(accounts), nil
}

func (serviceImpl accountServiceImpl) GetBalance(ctx context.Context, accountId string) (output.FindAccountBalanceOutputVO, error) {
	account, err := serviceImpl.repository.GetBalance(ctx, accountId)
	if err != nil {
		return output.FindAccountBalanceOutputVO{}, errors.New("error getting account balance")
	}

	return converter.ModelToFindAccountBalanceOutputVO(account), nil
}
