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
			Name:    "Lorena Morena",
			CPF:     "86419560004",
			Balance: 2000,
		}
		mockRepo.On("FindOne").Return(account, nil)
		mockRepo.On("UpdateBalance").Return(nil)
		accountUC := usecase.NewAccountUc(mockRepo, nil)

		//test
		err := accountUC.Deposit("86419560004", 100)

		//assert
		assert.Nil(t, err)
	})
	t.Run("the deposit must not be made because the account cannot be found", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)

		mockRepo.On("FindOne").Return(entities.Account{}, errors.New("error"))
		mockRepo.On("UpdateBalance").Return(nil)
		accountUC := usecase.NewAccountUc(mockRepo, log)

		//test
		err := accountUC.Deposit("86419560004", 100)

		//assert
		assert.Equal(t, err.Error(), "failed to recover account")

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

		//test
		accountUC := usecase.NewAccountUc(mockRepo, log)

		//assert
		err := accountUC.Deposit("86419560004", 0)
		assert.Equal(t, err.Error(), "invalid value")

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
		err := accountUC.Deposit("86419560004", 100)

		//assert
		assert.Equal(t, err.Error(), "failed to update balance")

	})

}
