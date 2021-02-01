package validator

import (
	"github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
)

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

	if createAccountInputVO.Balance < 0 {
		return custom_errors.ErrorAccountBalanceValue
	}

	return nil
}
