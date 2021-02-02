package repository

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
)

// MockTransferRepositoryImpl struct of Transfer Service mock
type MockTransferRepositoryImpl struct {
	Result  model.Transfer
	Results []model.Transfer
	Err     error
}

// Create - creates a new transfer mock
func (m MockTransferRepositoryImpl) Create(ctx context.Context, transfer model.Transfer) (*model.Transfer, error) {
	return &m.Result, m.Err
}

// FindAll - list all transfers mock
func (m MockTransferRepositoryImpl) FindAll(ctx context.Context, accountOriginID string) ([]model.Transfer, error) {
	return m.Results, m.Err
}
