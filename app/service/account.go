package service

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/converter"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/logger"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/brunodecastro/digital-accounts/app/util"
	"go.uber.org/zap"
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

func NewAccountService(repository repository.AccountRepository, transactionHelper postgres.TransactionHelper) AccountService {
	return &accountServiceImpl{
		repository:        repository,
		transactionHelper: transactionHelper,
	}
}

func (serviceImpl accountServiceImpl) Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (output.CreateAccountOutputVO, error) {
	logApi := logger.GetLogger().With(
		zap.String("resource", "AccountService"),
		zap.String("method", "Create"))

	var emptyCreateAccountOutputVO = output.CreateAccountOutputVO{}

	ctx, err := serviceImpl.transactionHelper.StartTransaction(ctx)
	if err != nil {
		logApi.Error(err.Error())
		return emptyCreateAccountOutputVO, custom_errors.ErrorStartTransaction
	}

	accountExists, err := serviceImpl.repository.FindByCpf(ctx, util.NumbersOnly(accountInputVO.Cpf))
	if err != nil {
		logApi.Error(err.Error())
		return emptyCreateAccountOutputVO, custom_errors.ErrorUnexpected
	}
	// Check if there is an account with the same cpf
	if accountExists != nil && accountExists.Id != "" {
		return emptyCreateAccountOutputVO, custom_errors.ErrorAccountCpfExists
	}

	accountCreated, err := serviceImpl.repository.Create(ctx, converter.CreateAccountInputVOToModel(accountInputVO))
	if err != nil {
		if errT := serviceImpl.transactionHelper.RollbackTransaction(ctx); errT != nil {
			logApi.Error(errT.Error())
			return emptyCreateAccountOutputVO, custom_errors.ErrorRollbackTransaction
		}
		logApi.Error(err.Error())
		return emptyCreateAccountOutputVO, custom_errors.ErrorCreateAccount
	}

	if err := serviceImpl.transactionHelper.CommitTransaction(ctx); err != nil {
		logApi.Error(err.Error())
		return emptyCreateAccountOutputVO, custom_errors.ErrorCommitTransaction
	}

	return converter.ModelToCreateAccountOutputVO(accountCreated), nil
}

func (serviceImpl accountServiceImpl) FindAll(ctx context.Context, ) ([]output.FindAllAccountOutputVO, error) {
	logApi := logger.GetLogger().With(
		zap.String("resource", "AccountService"),
		zap.String("method", "FindAll"))

	accounts, err := serviceImpl.repository.FindAll(ctx)
	if err != nil {
		logApi.Error(err.Error())
		return []output.FindAllAccountOutputVO{}, custom_errors.ErrorListingAllAccounts
	}

	return converter.AccountModelToFindAllAccountOutputVO(accounts), nil
}

func (serviceImpl accountServiceImpl) GetBalance(ctx context.Context, accountId string) (output.FindAccountBalanceOutputVO, error) {
	logApi := logger.GetLogger().With(
		zap.String("resource", "AccountService"),
		zap.String("method", "GetBalance"))

	account, err := serviceImpl.repository.GetBalance(ctx, accountId)
	if err != nil {
		logApi.Error(err.Error())
		return output.FindAccountBalanceOutputVO{}, custom_errors.ErrorGettingAccountBalance
	}

	if account == nil || account.Id == "" {
		return output.FindAccountBalanceOutputVO{}, custom_errors.ErrorAccountNotFound
	}

	return converter.ModelToFindAccountBalanceOutputVO(account), nil
}
