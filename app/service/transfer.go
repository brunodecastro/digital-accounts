package service

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/converter"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/logger"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"go.uber.org/zap"
)

type TransferService interface {
	Create(ctx context.Context, transferInputVO input.CreateTransferInputVO) (output.CreateTransferOutputVO, error)
	FindAll(ctx context.Context, accountOriginId string) ([]output.FindAllTransferOutputVO, error)
}

type transferServiceImpl struct {
	transferRepository repository.TransferRepository
	accountRepository  repository.AccountRepository
	transactionHelper  postgres.TransactionHelper
}

func NewTransferService(
	transferRepository repository.TransferRepository,
	accountRepository repository.AccountRepository,
	transactionHelper postgres.TransactionHelper) TransferService {
	return &transferServiceImpl{
		transferRepository: transferRepository,
		accountRepository:  accountRepository,
		transactionHelper:  transactionHelper,
	}
}

func (serviceImpl transferServiceImpl) Create(ctx context.Context, transferInputVO input.CreateTransferInputVO) (output.CreateTransferOutputVO, error) {
	logApi := logger.GetLogger().With(
		zap.String("resource", "TransferService"),
		zap.String("method", "Create"))

	ctx, err := serviceImpl.transactionHelper.StartTransaction(ctx)
	if err != nil {
		logApi.Error(err.Error())
		return output.CreateTransferOutputVO{}, custom_errors.ErrorStartTransaction
	}

	// Transfer amount between accounts
	err = serviceImpl.transferAmountBetweenAccounts(ctx, transferInputVO)
	if err != nil {
		if errT := serviceImpl.transactionHelper.RollbackTransaction(ctx); errT != nil {
			logApi.Error(errT.Error())
			return output.CreateTransferOutputVO{}, custom_errors.ErrorRollbackTransaction
		}
		logApi.Error(err.Error())
		return output.CreateTransferOutputVO{}, err
	}

	// Create the transfer record
	transferCreated, err := serviceImpl.transferRepository.Create(ctx, converter.CreateTransferInputVOToModel(transferInputVO))
	if err != nil {
		logApi.Error(err.Error())
		return output.CreateTransferOutputVO{}, custom_errors.ErrorCreateTransfer
	}

	if err := serviceImpl.transactionHelper.CommitTransaction(ctx); err != nil {
		logApi.Error(err.Error())
		return output.CreateTransferOutputVO{}, custom_errors.ErrorCommitTransaction
	}

	return converter.ModelToCreateTransferOutputVO(transferCreated), nil
}

func (serviceImpl transferServiceImpl) transferAmountBetweenAccounts(ctx context.Context, transferInputVO input.CreateTransferInputVO) error {
	if transferInputVO.Amount <= 0 {
		return custom_errors.ErrorTransferAmountValue
	}

	// Check if the transfer is to the same account
	if transferInputVO.AccountOriginId == transferInputVO.AccountDestinationId {
		return custom_errors.ErrorTransferSameAccount
	}

	// Chek if the account origin exists
	accountOrigin, err := serviceImpl.accountRepository.FindById(ctx, transferInputVO.AccountOriginId)
	if err != nil || accountOrigin == nil {
		return custom_errors.ErrorAccountOriginNotFound
	}

	// Chek if the account origin has sufficient balance
	if transferInputVO.Amount > accountOrigin.Balance.GetInt64() {
		return custom_errors.ErrorInsufficientBalance
	}

	// Chek if the account destination exists
	accountDestination, err := serviceImpl.accountRepository.FindById(ctx, transferInputVO.AccountDestinationId)
	if err != nil || accountDestination == nil {
		return custom_errors.ErrorAccountDestinationNotFound
	}

	var newAccountOriginBalance = types.Money(accountOrigin.Balance.GetInt64() - transferInputVO.Amount)
	var newAccountDestinationBalance = types.Money(accountDestination.Balance.GetInt64() + transferInputVO.Amount)

	var rowsAffecteds int64
	// Update origin account balance
	rowsAffecteds, err = serviceImpl.accountRepository.UpdateBalance(ctx, accountOrigin.Id, newAccountOriginBalance)
	if err != nil || rowsAffecteds < 1 {
		return custom_errors.ErrorUpdateAccountOriginBalance
	}

	// Update destination account balance
	rowsAffecteds, err = serviceImpl.accountRepository.UpdateBalance(ctx, accountDestination.Id, newAccountDestinationBalance)
	if err != nil || rowsAffecteds < 1 {
		return custom_errors.ErrorUpdateAccountDestinationBalance
	}

	return nil
}

func (serviceImpl transferServiceImpl) FindAll(ctx context.Context, accountOriginId string) ([]output.FindAllTransferOutputVO, error) {
	logApi := logger.GetLogger().With(
		zap.String("resource", "TransferService"),
		zap.String("method", "FindAll"))

	transfers, err := serviceImpl.transferRepository.FindAll(ctx, accountOriginId)
	if err != nil {
		logApi.Error(err.Error())
		return []output.FindAllTransferOutputVO{}, custom_errors.ErrorListingAllTransfers
	}

	return converter.ModelToFindAllTransferOutputVO(transfers), nil
}
