package usecase

import "errors"

var (
	ErrorAccountAlreadyExists error = errors.New("The account already exists")
)
