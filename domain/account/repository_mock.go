package account
import (
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

//implementation of repository interfaces

func (mock *MockRepository) FindOne(CPF string) (entities.Account, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entities.Account), args.Error(1)
}
func (mock *MockRepository) UpdateBalance(ID string, balance float64) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockRepository) FindByID(accountID string) (entities.Account, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entities.Account), args.Error(1)
}
func (mock *MockRepository) InsertAccount(entities.Account) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockRepository) ListAllAccounts() ([]entities.Account, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.Account), args.Error(1)
}