package service

import (
	"context"
	"errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/common/vo/output"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/brunodecastro/digital-accounts/app/util"
)

type AuthenticateService interface {
	Authenticate(ctx context.Context, credentialInputVO input.CredentialInputVO) (output.CredentialOutputVO, error)
}

type authenticateServiceImpl struct {
	accountRepository repository.AccountRepository
}

func NewAuthenticateService(accountRepository repository.AccountRepository) AuthenticateService {
	return &authenticateServiceImpl{
		accountRepository: accountRepository,
	}
}

func (serviceImpl authenticateServiceImpl) Authenticate(ctx context.Context, credentialInputVO input.CredentialInputVO) (output.CredentialOutputVO, error) {
	account, err := serviceImpl.accountRepository.FindByCpf(ctx, util.NumbersOnly(credentialInputVO.Cpf))
	if err != nil {
		return output.CredentialOutputVO{}, errors.New("account not found with this cpf")
	}

	// Check the password with secret
	if !util.CheckPasswordHash(credentialInputVO.Password, account.Secret) {
		return output.CredentialOutputVO{}, errors.New("wrong password")
	}

	return output.CredentialOutputVO{
		AccountId: string(account.Id),
		Cpf:       account.Cpf,
		Username:  account.Name,
	}, nil
}
