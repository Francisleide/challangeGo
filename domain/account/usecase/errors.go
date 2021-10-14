package usecase

import "errors"

var (
	ErrorAccountAlreadyExists error = errors.New("the account already exists")
	ErrorRetrieveAccount      error = errors.New("failed to retrieve account")
	ErrorInvalidValue         error = errors.New("invalid value")
	ErrorUpdateBalance        error = errors.New("failed to update balance")
	ErrorListAccounts         error = errors.New("failed to list accounts")
	ErrorInsufficientBalance  error = errors.New("insufficient balance")
)
