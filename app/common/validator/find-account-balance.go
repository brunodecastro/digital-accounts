package validator

import (
	"github.com/brunodecastro/digital-accounts/app/common/custom-errors"
)

func ValidateFindAccountBalanceInput(accountId string) error {
	if accountId == "" {
		return custom_errors.ErrorAccountIdRequired
	}
	return nil
}
