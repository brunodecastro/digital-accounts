package service

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
)

// MockAccountService - struct that mocks the account service
type MockAccountService struct {
	ResultCreateAccount output.CreateAccountOutputVO
	ResultGetAll        []output.FindAllAccountOutputVO
	ResultGetBalance    output.FindAccountBalanceOutputVO
	Err                 error
}

// Create - creates a new account mock
func (m MockAccountService) Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (output.CreateAccountOutputVO, error) {
	return m.ResultCreateAccount, m.Err
}

// FindAll - list all accounts mock
func (m MockAccountService) FindAll(ctx context.Context) ([]output.FindAllAccountOutputVO, error) {
	return m.ResultGetAll, m.Err
}

// GetBalance - gets the account balance mock
func (m MockAccountService) GetBalance(ctx context.Context, accountId string) (output.FindAccountBalanceOutputVO, error) {
	return m.ResultGetBalance, m.Err
}
