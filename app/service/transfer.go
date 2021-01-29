package service

import (
	"context"
	"errors"
	"github.com/brunodecastro/digital-accounts/app/common/converter"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
)

type TransferService interface {
	Create(ctx context.Context, transferInputVO input.CreateTransferInputVO) (output.CreateTransferOutputVO, error)
	FindAll(ctx context.Context) ([]output.FindAllTransferOutputVO, error)
}

type transferServiceImpl struct {
	repository repository.TransferRepository
}

func NewTransferService(repository repository.TransferRepository) TransferService {
	return &transferServiceImpl{
		repository: repository,
	}
}

func (serviceImpl transferServiceImpl) Create(ctx context.Context, transferInputVO input.CreateTransferInputVO) (output.CreateTransferOutputVO, error) {
	transferCreated, err := serviceImpl.repository.Create(ctx, converter.CreateTransferInputVOToModel(transferInputVO))
	if err != nil {
		return output.CreateTransferOutputVO{}, errors.New("error on create transfer")
	}

	return converter.ModelToCreateTransferOutputVO(transferCreated), nil
}

func (serviceImpl transferServiceImpl) FindAll(ctx context.Context) ([]output.FindAllTransferOutputVO, error) {
	transfers, err := serviceImpl.repository.FindAll(ctx)
	if err != nil {
		return []output.FindAllTransferOutputVO{}, errors.New("error listing all transfers")
	}

	return converter.ModelToFindAllTransferOutputVO(transfers), nil
}
