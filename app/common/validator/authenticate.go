package validator

import (
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/util"
)

// ValidateAuthenticate -  validates the input.CredentialInputVO for authenticate
func ValidateAuthenticate(credentialInput input.CredentialInputVO) error {

	if credentialInput.Cpf == "" {
		return custom_errors.ErrorCredentialCpfRequired
	}

	if credentialInput.Secret == "" {
		return custom_errors.ErrorCredentialSecretRequired
	}

	if !util.IsCpfValid(credentialInput.Cpf) {
		return custom_errors.ErrorCpfInvalid
	}

	return nil
}
