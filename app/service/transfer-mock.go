package service

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
)

// MockTransferService - struct of Transfer Service mock
type MockTransferService struct {
	ResultCreateTransfer output.CreateTransferOutputVO
	ResultFindAll        []output.FindAllTransferOutputVO
	Err                  error
}

// Create - creates a new transfer mock
func (m MockTransferService) Create(ctx context.Context, transferInputVO input.CreateTransferInputVO) (output.CreateTransferOutputVO, error) {
	return m.ResultCreateTransfer, m.Err
}

// FindAll - list all transfers mock
func (m MockTransferService) FindAll(ctx context.Context, accountOriginID string) ([]output.FindAllTransferOutputVO, error) {
	return m.ResultFindAll, m.Err
}
