package validator

import (
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/util"
)

// ValidateAuthenticate -  validates the input.CredentialInputVO for authenticate
func ValidateAuthenticate(credentialInput input.CredentialInputVO) error {

	if credentialInput.CPF == "" {
		return custom_errors.ErrorCredentialCPFRequired
	}

	if credentialInput.Secret == "" {
		return custom_errors.ErrorCredentialSecretRequired
	}

	if !util.IsCpfValid(credentialInput.CPF) {
		return custom_errors.ErrorCpfInvalid
	}

	return nil
}
