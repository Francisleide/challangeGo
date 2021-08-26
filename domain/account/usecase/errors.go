package usecase

import "errors"

var (
	ErrorAccountAlreadyExists error = errors.New("the account already exists")
)
