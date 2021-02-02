package repository

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/types"
	"github.com/brunodecastro/digital-accounts/app/model"
)

// MockAccountRepositoryImpl - struct of account repository mock
type MockAccountRepositoryImpl struct {
	Result          model.Account
	Results         []model.Account
	ResultBalance   int64
	ResultFindByCPF *model.Account
	ResultFindByID  *model.Account
	Err             error
	ErrFindByCPF    error
}

// Create - creates a new account mock
func (m MockAccountRepositoryImpl) Create(ctx context.Context, account model.Account) (*model.Account, error) {
	return &m.Result, m.Err
}

// FindAll - list all accounts mock
func (m MockAccountRepositoryImpl) FindAll(ctx context.Context) ([]model.Account, error) {
	return m.Results, m.Err
}

// GetBalance - gets the account balance mock
func (m MockAccountRepositoryImpl) GetBalance(ctx context.Context, accountID string) (*model.Account, error) {
	return &m.Result, m.Err
}

// UpdateBalance - updates the account balance mock
func (m MockAccountRepositoryImpl) UpdateBalance(ctx context.Context, accountOriginID model.AccountID, balance types.Money) (int64, error) {
	return m.ResultBalance, m.Err
}

// FindByCPF - find an account by cpf mock
func (m MockAccountRepositoryImpl) FindByCPF(ctx context.Context, cpf string) (*model.Account, error) {
	return m.ResultFindByCPF, m.ErrFindByCPF
}

// FindByID - find an account by the id mock
func (m MockAccountRepositoryImpl) FindByID(ctx context.Context, accountOriginID string) (*model.Account, error) {
	return m.ResultFindByID, m.Err
}
