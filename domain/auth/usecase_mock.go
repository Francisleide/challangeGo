package auth

import "github.com/stretchr/testify/mock"

type UsecaseMock struct {
	mock.Mock
}

func (mock *UsecaseMock) CreateToken(CPF string, secret string) (string, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
func (mock *UsecaseMock) Login(CPF, secret string) error {
	args := mock.Called()
	return args.Error(0)
}
