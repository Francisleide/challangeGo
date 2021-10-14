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

func TestDeposit(t *testing.T) {
	t.Run("deposit should take place without errors", func(t *testing.T) {
		//prepare
		mockRepo := new(account.MockRepository)
		account := entities.Account{
			ID:      "8aecf60b-b549-41a2-b9b9-143d2d513c87",
			Name:    "Lorena Morena",
			CPF:     "86419560004",
			Balance: 2000,
		}
		mockRepo.On("FindOne").Return(account, nil)
		mockRepo.On("UpdateBalance").Return(nil)
		accountUC := usecase.NewAccountUc(mockRepo, nil)
		expectedDepositOut := entities.TransactionOutput{
			ID:              account.ID,
			PreviousBalance: account.Balance,
			ActualBalance:   account.Balance + 100,
		}

		//test
		depositOutReceived, err := accountUC.Deposit("86419560004", 100)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, expectedDepositOut, depositOutReceived)

	})
	t.Run("the deposit must not be made because the account cannot be found", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)

		mockRepo.On("FindOne").Return(entities.Account{}, errors.New("error"))
		mockRepo.On("UpdateBalance").Return(nil)
		accountUC := usecase.NewAccountUc(mockRepo, log)

		//test
		depositOutReceived, err := accountUC.Deposit("86419560004", 100)

		//assert
		assert.ErrorIs(t, usecase.ErrorRetrieveAccount, err)
		assert.Equal(t, entities.TransactionOutput{}, depositOutReceived)

	})
	t.Run("the deposit must not be made because the amount is not valid", func(t *testing.T) {
		//prepare
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

		//test
		depositOutReceived, err := accountUC.Deposit("86419560004", 0)

		//assert
		assert.ErrorIs(t, usecase.ErrorInvalidValue, err)
		assert.Equal(t, entities.TransactionOutput{}, depositOutReceived)

	})
	t.Run("the deposit does not take place as there was a failure to update the balance", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		account := entities.Account{
			Name:    "Lorena Morena",
			CPF:     "86419560004",
			Balance: 2000,
		}
		mockRepo.On("FindOne").Return(account, nil)
		mockRepo.On("UpdateBalance").Return(errors.New("Error"))
		accountUC := usecase.NewAccountUc(mockRepo, log)

		//test
		depositOutReceived, err := accountUC.Deposit("86419560004", 100)

		//assert
		assert.ErrorIs(t, usecase.ErrorUpdateBalance, err)
		assert.Equal(t, entities.TransactionOutput{}, depositOutReceived)

	})

}
