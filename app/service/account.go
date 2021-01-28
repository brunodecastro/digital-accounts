package service

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	output "github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/model"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/brunodecastro/digital-accounts/app/util"
	"time"
)

type AccountService interface {
	Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (output.CreateAccountOutputVO, error)
	GetAll(ctx context.Context) ([]model.Account, error)
	GetBalance(ctx context.Context, accountId string) (*model.Account, error)
}

type accountServiceImpl struct {
	repository repository.AccountRepository
}

func NewAccountService(repository repository.AccountRepository) AccountService {
	return &accountServiceImpl{
		repository: repository,
	}
}

func (serviceImpl accountServiceImpl) Create(ctx context.Context, accountInputVO input.CreateAccountInputVO) (output.CreateAccountOutputVO, error) {
	//ctx, cancelFunc := context.WithTimeout(ctx, 5*time.Second)
	//defer cancelFunc()

	accountCreated, err := serviceImpl.repository.Create(ctx, accountVOToModel(accountInputVO))
	util.MaybeError(err, "error on create accounts")

	return accountModelToVO(accountCreated), err
}

func (serviceImpl accountServiceImpl) GetAll(ctx context.Context, ) ([]model.Account, error) {
	return serviceImpl.repository.GetAll(ctx)
}

func (serviceImpl accountServiceImpl) GetBalance(ctx context.Context, accountId string) (*model.Account, error) {
	return serviceImpl.repository.GetBalance(ctx, accountId)
}

func accountVOToModel(accountInputVO input.CreateAccountInputVO) model.Account {
	return model.Account{
		ID:        model.AccountID(common.NewUUID()),
		Name:      accountInputVO.Name,
		Cpf:       util.NumbersOnly(accountInputVO.Cpf),
		Secret:    util.EncryptPassword(accountInputVO.Secret),
		Balance:   common.Money(accountInputVO.Balance),
		CreatedAt: time.Now(),
	}
}

func accountModelToVO(account *model.Account) output.CreateAccountOutputVO {
	return output.CreateAccountOutputVO{
		Name:      account.Name,
		Cpf:       util.CpfFormat(account.Cpf),
		Balance:   account.Balance.GetFloat(),
		CreatedAt: account.CreatedAt.Format(time.RFC3339), // "2006-01-02T15:04:05Z07:00"
	}
}
