package entities

import "errors"

var(
	ErrorInvalidSecret error = errors.New("invalid secret")
	ErrorNameEmpty error = errors.New("the name cannot be null")
	ErrorInvalidCPF error = errors.New("invalid cpf")
)