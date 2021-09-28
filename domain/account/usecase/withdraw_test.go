package usecase_test

import (
	"testing"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/account/usecase"
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestWithdraw(t *testing.T) {
	t.Run("withdraw should take place without errors", func(t *testing.T) {
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		account := entities.Account{
			Name:    "Lorena Morena",
			CPF:     "86419560004",
			Balance: 2000,
		}
		mockRepo.On("FindOne").Return(account, nil)
		mockRepo.On("UpdateBalance").Return(nil)
		accountUC := usecase.NewAccountUc(mockRepo, log)
		err := accountUC.Withdraw(account.CPF, 500.0)

		assert.Nil(t, err)
	})
	t.Run("the withdrawal is not made due to lack of balance", func(t *testing.T) {
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		account := entities.Account{
			Name:    "Lorena Morena",
			CPF:     "86419560004",
			Balance: 2000,
		}
		mockRepo.On("FindOne").Return(account, nil)
		mockRepo.On("UpdateBalance").Return(nil)
		accountUC := usecase.NewAccountUc(mockRepo, log)
		err := accountUC.Withdraw(account.CPF, 3000)
		assert.Equal(t, err.Error(), "insufficient balance")
	})
	
}
