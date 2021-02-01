package service

import (
	"context"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/logger"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/brunodecastro/digital-accounts/app/util"
	"go.uber.org/zap"
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
	logApi := logger.GetLogger().With(
		zap.String("resource", "AuthenticationService"),
		zap.String("method", "Authenticate"))

	var emptyCredentialVO = vo.CredentialVO{}

	if !util.IsCpfValid(credentialInputVO.Cpf) {
		return emptyCredentialVO, custom_errors.ErrorCpfInvalid
	}

	account, err := serviceImpl.accountRepository.FindByCpf(ctx, util.NumbersOnly(credentialInputVO.Cpf))
	if err != nil {
		logApi.Error(err.Error())
		return emptyCredentialVO, custom_errors.ErrorUnexpected
	}

	// Check if the account exists
	if account == nil || account.Id == "" {
		return emptyCredentialVO, custom_errors.ErrorAccountNotFound
	}

	// Check the password with secret
	if !util.CheckPasswordHash(credentialInputVO.Password, account.Secret) {
		return emptyCredentialVO, custom_errors.ErrorCredentialWrongPassword
	}

	return vo.CredentialVO{
		AccountId: string(account.Id),
		Cpf:       account.Cpf,
		Username:  account.Name,
	}, nil
}
