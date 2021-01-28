package repository

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/model"
)

type MockAccountRepositoryImpl struct {
	Result model.Account
	Results []model.Account
	Err    error
}

func (m MockAccountRepositoryImpl) Create(ctx context.Context, account model.Account) (*model.Account, error) {
	return &m.Result, m.Err
}

func (m MockAccountRepositoryImpl) GetAll(ctx context.Context) ([]model.Account, error) {
	return m.Results, m.Err
}

func (m MockAccountRepositoryImpl) GetBalance(ctx context.Context, accountId string) (*model.Account, error) {
	return &m.Result, m.Err
}