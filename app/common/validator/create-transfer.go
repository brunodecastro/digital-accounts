package validator

import (
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
)

// ValidateCreateTransferInput - validates the input.CreateTransferInputVO to create a new transfer
func ValidateCreateTransferInput(createTransferInputVO input.CreateTransferInputVO) error {

	if createTransferInputVO.AccountDestinationID == "" {
		return custom_errors.ErrorAccountDestinationIDRequired
	}

	if createTransferInputVO.Amount <= 0 {
		return custom_errors.ErrorTransferAmountValue
	}

	return nil
}
