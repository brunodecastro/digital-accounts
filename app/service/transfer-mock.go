package service

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
)

type MockTransferService struct {
	ResultCreateTransfer output.CreateTransferOutputVO
	ResultFindAll        []output.FindAllTransferOutputVO
	Err                  error
}

func (m MockTransferService) Create(ctx context.Context, transferInputVO input.CreateTransferInputVO) (output.CreateTransferOutputVO, error) {
	return m.ResultCreateTransfer, m.Err
}

func (m MockTransferService) FindAll(ctx context.Context, accountOriginId string) ([]output.FindAllTransferOutputVO, error) {
	return m.ResultFindAll, m.Err
}
