package transfer

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct {
	mock.Mock
}

func (mock *UsecaseMock) CreateTransfer(accountOrigin, accountDestination entities.Account, amount float64) (entities.Transfer, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entities.Transfer), args.Error(1)
}

func (mock *UsecaseMock) ListUserTransfers(CPF string) ([]entities.Transfer, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.Transfer), args.Error(1)
}
