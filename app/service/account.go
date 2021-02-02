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

// AccountService - interface of account service
type AccountService interface {
	Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (output.CreateAccountOutputVO, error)
	FindAll(ctx context.Context) ([]output.FindAllAccountOutputVO, error)
	GetBalance(ctx context.Context, accountId string) (output.FindAccountBalanceOutputVO, error)
}

type accountServiceImpl struct {
	repository        repository.AccountRepository
	transactionHelper postgres.TransactionHelper
}

// NewAccountService - instance of account service
func NewAccountService(repository repository.AccountRepository, transactionHelper postgres.TransactionHelper) AccountService {
	return &accountServiceImpl{
		repository:        repository,
		transactionHelper: transactionHelper,
	}
}

// Create - creates a new account
func (serviceImpl accountServiceImpl) Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (output.CreateAccountOutputVO, error) {
	logAPI := logger.GetLogger().With(
		zap.String("resource", "AccountService"),
		zap.String("method", "Create"))

	var emptyCreateAccountOutputVO = output.CreateAccountOutputVO{}

	ctx, err := serviceImpl.transactionHelper.StartTransaction(ctx)
	if err != nil {
		logAPI.Error(err.Error())
		return emptyCreateAccountOutputVO, custom_errors.ErrorStartTransaction
	}

	accountExists, err := serviceImpl.repository.FindByCpf(ctx, util.NumbersOnly(accountInputVO.Cpf))
	if err != nil {
		logAPI.Error(err.Error())
		return emptyCreateAccountOutputVO, custom_errors.ErrorUnexpected
	}
	// Check if there is an account with the same cpf
	if accountExists != nil && accountExists.ID != "" {
		return emptyCreateAccountOutputVO, custom_errors.ErrorAccountCpfExists
	}

	accountCreated, err := serviceImpl.repository.Create(ctx, converter.CreateAccountInputVOToModel(accountInputVO))
	if err != nil {
		if errT := serviceImpl.transactionHelper.RollbackTransaction(ctx); errT != nil {
			logAPI.Error(errT.Error())
			return emptyCreateAccountOutputVO, custom_errors.ErrorRollbackTransaction
		}
		logAPI.Error(err.Error())
		return emptyCreateAccountOutputVO, custom_errors.ErrorCreateAccount
	}

	if err := serviceImpl.transactionHelper.CommitTransaction(ctx); err != nil {
		logAPI.Error(err.Error())
		return emptyCreateAccountOutputVO, custom_errors.ErrorCommitTransaction
	}

	return converter.ModelToCreateAccountOutputVO(accountCreated), nil
}

// FindAll - list all accounts
func (serviceImpl accountServiceImpl) FindAll(ctx context.Context) ([]output.FindAllAccountOutputVO, error) {
	logAPI := logger.GetLogger().With(
		zap.String("resource", "AccountService"),
		zap.String("method", "FindAll"))

	accounts, err := serviceImpl.repository.FindAll(ctx)
	if err != nil {
		logAPI.Error(err.Error())
		return []output.FindAllAccountOutputVO{}, custom_errors.ErrorListingAllAccounts
	}

	return converter.AccountModelToFindAllAccountOutputVO(accounts), nil
}

// GetBalance - gets the account balance
func (serviceImpl accountServiceImpl) GetBalance(ctx context.Context, accountId string) (output.FindAccountBalanceOutputVO, error) {
	logAPI := logger.GetLogger().With(
		zap.String("resource", "AccountService"),
		zap.String("method", "GetBalance"))

	account, err := serviceImpl.repository.GetBalance(ctx, accountId)
	if err != nil {
		logAPI.Error(err.Error())
		return output.FindAccountBalanceOutputVO{}, custom_errors.ErrorGettingAccountBalance
	}

	if account == nil || account.ID == "" {
		return output.FindAccountBalanceOutputVO{}, custom_errors.ErrorAccountNotFound
	}

	return converter.ModelToFindAccountBalanceOutputVO(account), nil
}
