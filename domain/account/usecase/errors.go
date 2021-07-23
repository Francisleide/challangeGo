package usecase

import "errors"

var (
	ErrorAccountAlreadyExists error = errors.New("A conta já existe!")
)