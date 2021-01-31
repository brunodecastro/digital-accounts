package service

import (
	"context"
	"errors"
	"github.com/brunodecastro/digital-accounts/app/common/converter"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/persistence/database/postgres"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
)

type TransferService interface {
	Create(ctx context.Context, transferInputVO input.CreateTransferInputVO) (output.CreateTransferOutputVO, error)
	FindAll(ctx context.Context) ([]output.FindAllTransferOutputVO, error)
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
	ctx, err := serviceImpl.transactionHelper.StartTransaction(ctx)
	if err != nil {
		return output.CreateTransferOutputVO{}, custom_errors.ErrorStartTransaction
	}

	// Transfer amount between accounts
	err = serviceImpl.transferAmountBetweenAccount(ctx, transferInputVO)
	if err != nil {
		errT := serviceImpl.transactionHelper.RollbackTransaction(ctx)
		if errT != nil {
			return output.CreateTransferOutputVO{}, custom_errors.ErrorRollbackTransaction
		}
		return output.CreateTransferOutputVO{}, err
	}

	// Create the transfer record
	transferCreated, err := serviceImpl.transferRepository.Create(ctx, converter.CreateTransferInputVOToModel(transferInputVO))
	if err != nil {
		return output.CreateTransferOutputVO{}, custom_errors.ErrorCreateTransfer
	}

	serviceImpl.transactionHelper.CommitTransaction(ctx)
	if err != nil {
		return output.CreateTransferOutputVO{}, custom_errors.ErrorCommitTransaction
	}

	return converter.ModelToCreateTransferOutputVO(transferCreated), nil
}

func (serviceImpl transferServiceImpl) transferAmountBetweenAccount(ctx context.Context, transferInputVO input.CreateTransferInputVO) error {

	if transferInputVO.Amount <= 0 {
		return custom_errors.ErrorAmountValue
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

func (serviceImpl transferServiceImpl) FindAll(ctx context.Context) ([]output.FindAllTransferOutputVO, error) {
	transfers, err := serviceImpl.transferRepository.FindAll(ctx)
	if err != nil {
		return []output.FindAllTransferOutputVO{}, errors.New("error listing all transfers")
	}

	return converter.ModelToFindAllTransferOutputVO(transfers), nil
}
