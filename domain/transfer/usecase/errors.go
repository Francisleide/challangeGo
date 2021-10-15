package usecase

import "errors"

var (
	ErrorInvalidTransfer error = errors.New("invalid transfer")
	ErrorSaveTransfer error = errors.New("unable to save transfer")
	ErrorInsufficientFunds error = errors.New("insufficient funds")
)