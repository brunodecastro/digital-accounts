package service

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
)

type MockAccountService struct {
	ResultCreateAccount output.CreateAccountOutputVO
	ResultGetAll        []output.FindAllAccountOutputVO
	ResultGetBalance    output.FindAccountBalanceOutputVO
	Err                 error
}

func (m MockAccountService) Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (output.CreateAccountOutputVO, error) {
	return m.ResultCreateAccount, m.Err
}

func (m MockAccountService) FindAll(ctx context.Context, ) ([]output.FindAllAccountOutputVO, error) {
	return m.ResultGetAll, m.Err
}

func (m MockAccountService) GetBalance(ctx context.Context, accountId string) (output.FindAccountBalanceOutputVO, error) {
	return m.ResultGetBalance, m.Err
}
