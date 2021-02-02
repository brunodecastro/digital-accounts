package validator

import (
	"github.com/brunodecastro/digital-accounts/app/common/errors"
)

// ValidateFindAccountBalanceInput - validate the accountId to find account balance
func ValidateFindAccountBalanceInput(accountID string) error {
	if accountID == "" {
		return errors.ErrorAccountIDRequired
	}
	return nil
}
