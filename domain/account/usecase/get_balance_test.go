package usecase_test

import (
	"errors"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/account/usecase"
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T) {
	//prepare
	mockRepo := new(account.MockRepository)
	t.Run("the account is found and the balance must be recovered without errors", func(t *testing.T) {
		account := entities.Account{
			Name:    "Silvia Silva",
			ID:      "8aecf60b-b549-41a2-b9b9-143d2d513c87",
			CPF:     "86419560004",
			Balance: 200,
		}
		mockRepo.On("FindByID").Return(account, nil)
		accountUC := usecase.NewAccountUc(mockRepo, nil)

		//test
		balanceReceived, err := accountUC.GetBalance(account.ID)
		
		//assert
		assert.Nil(t, err)
		assert.Equal(t, account.Balance, balanceReceived)
	})
	t.Run("the account is not found and the error is displayed.", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		account := entities.Account{
			Name:      "Silvia Silva",
			ID:        "8aecf60b-b549-41a2-b9b9-143d2d513c87",
			CPF:       "86419560004",
			Secret:    "abc123",
			Balance:   200,
			CreatedAt: "",
		}
		mockRepo.On("FindByID").Return(entities.Account{}, errors.New(""))
		accountUC := usecase.NewAccountUc(mockRepo, log)
		
		//test
		balanceReceived, err := accountUC.GetBalance(account.ID)

		//assert
		assert.Equal(t, err.Error(), "failed to retrieve the account from repository")
		assert.Equal(t, balanceReceived, float64(0))
	})
}
