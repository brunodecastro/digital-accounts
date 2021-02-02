package service

import (
	"context"
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/errors"
	"github.com/brunodecastro/digital-accounts/app/common/logger"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/persistence/repository"
	"github.com/brunodecastro/digital-accounts/app/util"
	"go.uber.org/zap"
)

// AuthenticationService - interface of Authentication Service
type AuthenticationService interface {
	Authenticate(ctx context.Context, credentialInputVO input.CredentialInputVO) (vo.CredentialVO, error)
}

type authenticationServiceImpl struct {
	accountRepository repository.AccountRepository
}

// NewAuthenticationService - new instance of Authentication Service impl
func NewAuthenticationService(accountRepository repository.AccountRepository) AuthenticationService {
	return &authenticationServiceImpl{
		accountRepository: accountRepository,
	}
}

// Authenticate - authenticate the user with credentials
func (serviceImpl authenticationServiceImpl) Authenticate(ctx context.Context, credentialInputVO input.CredentialInputVO) (vo.CredentialVO, error) {
	logAPI := logger.GetLogger().With(
		zap.String("resource", "AuthenticationService"),
		zap.String("method", "Authenticate"))

	var emptyCredentialVO = vo.CredentialVO{}

	if !util.IsCpfValid(credentialInputVO.CPF) {
		return emptyCredentialVO, custom_errors.ErrorCpfInvalid
	}

	account, err := serviceImpl.accountRepository.FindByCPF(ctx, util.NumbersOnly(credentialInputVO.CPF))
	if err != nil {
		logAPI.Error(err.Error())
		return emptyCredentialVO, custom_errors.ErrorUnexpected
	}

	// Check if the account exists
	if account == nil || account.ID == "" {
		return emptyCredentialVO, custom_errors.ErrorAccountNotFound
	}

	// Check the password with secret
	if !util.CheckPasswordHash(credentialInputVO.Secret, account.Secret) {
		return emptyCredentialVO, custom_errors.ErrorCredentialWrongSecret
	}

	return vo.CredentialVO{
		AccountID: string(account.ID),
		CPF:       account.CPF,
		Username:  account.Name,
	}, nil
}
