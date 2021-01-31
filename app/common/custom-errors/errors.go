package custom_errors

import (
	"errors"
)

var (
	ErrInvalidAccessCredentials          = errors.New("invalid access credentials")
	ErrorCreateTransfer                  = errors.New("error on create transfer")
	ErrorInsufficientBalance             = errors.New("insufficient balance in the origin account")
	ErrorAmountValue                     = errors.New("the transfer amount must be greater than 0")
	ErrorAccountOriginNotFound           = errors.New("account origin not found")
	ErrorAccountDestinationNotFound      = errors.New("account destination not found")
	ErrorUpdateAccountOriginBalance      = errors.New("error on update account origin balance")
	ErrorUpdateAccountDestinationBalance = errors.New("error on update account destination balance")
)
