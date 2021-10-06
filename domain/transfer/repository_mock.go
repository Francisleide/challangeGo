package transfer

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

//implementation of repository interfaces

func (mock *MockRepository) InsertTransfer(transfer entities.Transfer) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockRepository) UpdateBalance(ID string, balance float64) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockRepository) ListUserTransfers(CPF string) ([]entities.Transfer, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.Transfer), args.Error(1)
}
