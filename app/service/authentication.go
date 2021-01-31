package service

import (
	"context"
	"errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/brunodecastro/digital-accounts/app/util"
)

type AuthenticationService interface {
	Authenticate(ctx context.Context, credentialInputVO input.CredentialInputVO) (vo.CredentialVO, error)
}

type authenticationServiceImpl struct {
	accountRepository repository.AccountRepository
}

func NewAuthenticationService(accountRepository repository.AccountRepository) AuthenticationService {
	return &authenticationServiceImpl{
		accountRepository: accountRepository,
	}
}

func (serviceImpl authenticationServiceImpl) Authenticate(ctx context.Context, credentialInputVO input.CredentialInputVO) (vo.CredentialVO, error) {
	account, err := serviceImpl.accountRepository.FindByCpf(ctx, util.NumbersOnly(credentialInputVO.Cpf))
	if err != nil {
		return vo.CredentialVO{}, errors.New("account not found with this cpf")
	}

	// Check the password with secret
	if !util.CheckPasswordHash(credentialInputVO.Password, account.Secret) {
		return vo.CredentialVO{}, errors.New("wrong password")
	}

	return vo.CredentialVO{
		AccountId: string(account.Id),
		Cpf:       account.Cpf,
		Username:  account.Name,
	}, nil
}
