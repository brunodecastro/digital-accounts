package repository

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
)

type MockATransferRepositoryImpl struct {
	Result  model.Transfer
	Results []model.Transfer
	Err     error
}

func (m MockATransferRepositoryImpl) Create(ctx context.Context, transfer model.Transfer) (*model.Transfer, error) {
	return &m.Result, m.Err
}

func (m MockATransferRepositoryImpl) FindAll(ctx context.Context, accountId string) ([]model.Transfer, error) {
	return m.Results, m.Err
}
