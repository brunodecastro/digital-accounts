package validator

import (
	"github.com/brunodecastro/digital-accounts/app/common/errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/util"
)

// ValidateCreateAccountInput - validates the input.CreateAccountInputVO to create a new account
func ValidateCreateAccountInput(createAccountInputVO input.CreateAccountInputVO) error {

	if createAccountInputVO.CPF == "" {
		return errors.ErrorAccountCpfRequired
	}

	if createAccountInputVO.Name == "" {
		return errors.ErrorAccountNameRequired
	}

	if createAccountInputVO.Secret == "" {
		return errors.ErrorAccountSecretRequired
	}

	if !util.IsCpfValid(createAccountInputVO.CPF) {
		return errors.ErrorCpfInvalid
	}

	if createAccountInputVO.Balance < 0 {
		return errors.ErrorAccountBalanceValue
	}

	return nil
}
