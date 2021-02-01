package repository

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/model"
)

type MockAccountRepositoryImpl struct {
	Result          model.Account
	Results         []model.Account
	ResultBalance   int64
	ResultFindByCpf *model.Account
	ResultFindById  *model.Account
	Err             error
}

func (m MockAccountRepositoryImpl) Create(ctx context.Context, account model.Account) (*model.Account, error) {
	return &m.Result, m.Err
}

func (m MockAccountRepositoryImpl) FindAll(ctx context.Context) ([]model.Account, error) {
	return m.Results, m.Err
}

func (m MockAccountRepositoryImpl) GetBalance(ctx context.Context, accountId string) (*model.Account, error) {
	return &m.Result, m.Err
}

func (m MockAccountRepositoryImpl) UpdateBalance(ctx context.Context, accountOriginId model.AccountID, balance types.Money) (int64, error) {
	return m.ResultBalance, m.Err
}

func (m MockAccountRepositoryImpl) FindByCpf(ctx context.Context, cpf string) (*model.Account, error) {
	return m.ResultFindByCpf, m.Err
}

func (m MockAccountRepositoryImpl) FindById(ctx context.Context, accountId string) (*model.Account, error) {
	return m.ResultFindById, m.Err
}
