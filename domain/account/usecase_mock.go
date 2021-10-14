package account

import (
	"context"

	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct {
	mock.Mock
}

//implementation of usecases account interfaces

/*CreateAccount(account entities.AccountInput) (entities.Account, error)
ListAll() ([]entities.Account, error)
Deposit(CPF string, amount float64) error
Withdraw(CPF string, amount float64) error
GetBalance(accountID string) (float64, error)
GetAccountByID(ID string) (entities.Account, error)
GetAccountByCPF(CPF string) (entities.Account, error)*/

func (mock *UsecaseMock) CreateAccount(account entities.AccountInput) (entities.Account, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entities.Account), args.Error(1)
}

func (mock *UsecaseMock) ListAll() ([]entities.Account, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.Account), args.Error(1)
}

func (mock *UsecaseMock) Deposit(CPF string, amount float64) (entities.TransactionOutput, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entities.TransactionOutput), args.Error(1)
}

func (mock *UsecaseMock) Withdraw(CPF string, amount float64) (entities.TransactionOutput, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entities.TransactionOutput), args.Error(1)
}

func (mock *UsecaseMock) GetBalance(accountID string) (float64, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(float64), args.Error(1)
}
func (mock *UsecaseMock) GetAccountByID(ID string) (entities.Account, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entities.Account), args.Error(1)
}
func (mock *UsecaseMock) GetAccountByCPF(CPF string) (entities.Account, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entities.Account), args.Error(1)
}
func (mock *UsecaseMock) GetCPF(ctx context.Context) (string, bool) {
	args := mock.Called()
	result := args.Get(0)
	return result.(string), result.(bool)
}
