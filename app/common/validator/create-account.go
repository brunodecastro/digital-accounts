package validator

import (
	"github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
	"github.com/brunodecastro/digital-accounts/app/util"
)

// ValidateCreateAccountInput - validates the input.CreateAccountInputVO to create a new account
func ValidateCreateAccountInput(createAccountInputVO input.CreateAccountInputVO) error {

	if createAccountInputVO.Cpf == "" {
		return custom_errors.ErrorAccountCpfRequired
	}

	if createAccountInputVO.Name == "" {
		return custom_errors.ErrorAccountNameRequired
	}

	if createAccountInputVO.Secret == "" {
		return custom_errors.ErrorAccountSecretRequired
	}

	if !util.IsCpfValid(createAccountInputVO.Cpf) {
		return custom_errors.ErrorCpfInvalid
	}

	if createAccountInputVO.Balance < 0 {
		return custom_errors.ErrorAccountBalanceValue
	}

	return nil
}
