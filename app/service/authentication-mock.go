package service

import (
	"context"
	"github.com/brunodecastro/digital-accounts/app/common/vo"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
)

// MockAuthenticationService - struct that mocks the authentication service
type MockAuthenticationService struct {
	Result vo.CredentialVO
	Err    error
}

// Authenticate - authenticate the user with credentials mock
func (m MockAuthenticationService) Authenticate(ctx context.Context, credentialInputVO input.CredentialInputVO) (vo.CredentialVO, error) {
	return m.Result, m.Err
}
