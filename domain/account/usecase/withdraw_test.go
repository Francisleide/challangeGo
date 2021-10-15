package usecase_test

import (
	"testing"

	ac "github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/account/usecase"
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestWithdraw(t *testing.T) {
	t.Run("withdraw should take place without errors", func(t *testing.T) {
		//prepare
		mockRepo := new(ac.MockRepository)
		account := entities.Account{
			Name:    "Lorena Morena",
			CPF:     "86419560004",
			Balance: 2000,
		}
		accountExpected := ac.TransactionOutput{
			PreviousBalance: account.Balance,
			ActualBalance:   1500,
		}
		mockRepo.On("FindOne").Return(account, nil)
		mockRepo.On("UpdateBalance").Return(nil)
		accountUC := usecase.NewAccountUc(mockRepo, nil)

		//test
		withdrawOutput, err := accountUC.Withdraw(account.CPF, 500.0)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, accountExpected.ActualBalance, withdrawOutput.ActualBalance)
		assert.Equal(t, accountExpected.PreviousBalance, withdrawOutput.PreviousBalance)

	})
	t.Run("the withdrawal is not made due to lack of balance", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(ac.MockRepository)
		account := entities.Account{
			Name:    "Lorena Morena",
			CPF:     "86419560004",
			Balance: 2000,
		}
		mockRepo.On("FindOne").Return(account, nil)
		mockRepo.On("UpdateBalance").Return(nil)
		accountUC := usecase.NewAccountUc(mockRepo, log)

		//test
		_, err := accountUC.Withdraw(account.CPF, 3000)

		//assert
		assert.ErrorIs(t, usecase.ErrorInsufficientBalance, err)
	})

}
