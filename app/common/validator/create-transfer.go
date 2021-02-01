package validator

import (
	custom_errors "github.com/brunodecastro/digital-accounts/app/common/custom-errors"
	"github.com/brunodecastro/digital-accounts/app/common/vo/input"
)

func ValidateCreateTransferInput(createTransferInputVO input.CreateTransferInputVO) error {

	if createTransferInputVO.AccountDestinationId == "" {
		return custom_errors.ErrorAccountDestinationIdRequired
	}

	if createTransferInputVO.Amount <= 0 {
		return custom_errors.ErrorTransferAmountValue
	}

	return nil
}
