package errors

import (
	"errors"
)

// custom errors constants
var (
	ErrorCredentialCPFRequired    = errors.New("credential cpf is required")
	ErrorCredentialSecretRequired = errors.New("credential secret is required")
	ErrorCredentialWrongSecret    = errors.New("wrong secret")
	ErrInvalidAccessCredentials   = errors.New("invalid access credentials")
	ErrInvalidToken               = errors.New("invalid token")
	ErrAuthorizationHeader        = errors.New("an Authorization header is required")
	ErrInvalidAuthorizationToken  = errors.New("invalid authorization token")

	ErrorAccountIDRequired               = errors.New("account id is required")
	ErrorAccountCpfRequired              = errors.New("account cpf is required")
	ErrorAccountNameRequired             = errors.New("account name is required")
	ErrorAccountSecretRequired           = errors.New("account secret is required")
	ErrorAccountBalanceRequired          = errors.New("account balance is required")
	ErrorAccountNotFound                 = errors.New("account not found")
	ErrorAccountBalanceValue             = errors.New("account balance must be greater than or equal to 0")
	ErrorAccountCpfExists                = errors.New("there is already an account created with the same cpf")
	ErrorCreateAccount                   = errors.New("error on create account")
	ErrorListingAllAccounts              = errors.New("error listing all accounts")
	ErrorGettingAccountBalance           = errors.New("error getting account balance")
	ErrorUpdateAccountOriginBalance      = errors.New("error on update account origin balance")
	ErrorUpdateAccountDestinationBalance = errors.New("error on update account destination balance")

	ErrorCpfInvalid                   = errors.New("invalid cpf")
	ErrorCreateTransfer               = errors.New("error on create transfer")
	ErrorInsufficientBalance          = errors.New("insufficient balance in the origin account")
	ErrorAccountOriginIDRequired      = errors.New("account origin id is required")
	ErrorAccountDestinationIDRequired = errors.New("account destination id is required")
	ErrorAccountOriginNotFound        = errors.New("account origin not found")
	ErrorAccountDestinationNotFound   = errors.New("account destination not found")
	ErrorTransferAmountValue          = errors.New("the transfer amount must be greater than to 0")
	ErrorListingAllTransfers          = errors.New("error listing all transfers")
	ErrorTransferSameAccount          = errors.New("cannot transfer to the same account")

	ErrorStartTransaction    = errors.New("error when trying to start transaction")
	ErrorCommitTransaction   = errors.New("error when trying to commit transaction")
	ErrorRollbackTransaction = errors.New("error when trying to rollback transaction")

	ErrorInvalidJSONFormat = errors.New("invalid JSON format")
	ErrorUnexpected        = errors.New("unexpected error")
)
