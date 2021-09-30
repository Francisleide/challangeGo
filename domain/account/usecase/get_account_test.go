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

func TestGetAccountByID(t *testing.T) {
	t.Run("the account was found via id and no error occurred", func(t *testing.T) {
		//prepare
		mockRepo := new(account.MockRepository)
		account := entities.Account{
			Name:      "João Paixão",
			ID:        "8aecf60b-b549-41a2-b9b9-143d2d513c87",
			Secret:    "123abc",
			Balance:   100,
			CreatedAt: "",
		}
		mockRepo.On("FindByID").Return(account, nil)
		accountUC := usecase.NewAccountUc(mockRepo, nil)

		//test
		accountReceived, err := accountUC.GetAccountByID(account.ID)

		//assert
		assert.Equal(t, account, accountReceived)
		assert.Nil(t, err)
	})
	t.Run("the id was sent for the account to be found, but there was an error finding it", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		mockRepo.On("FindByID").Return(entities.Account{}, errors.New(""))
		accountUC := usecase.NewAccountUc(mockRepo, log)
		account := entities.Account{
			Name:      "João Paixão",
			ID:        "8aecf60b-b549-41a2-b9b9-143d2d513c87",
			Secret:    "123abc",
			Balance:   100,
			CreatedAt: "",
		}

		//test
		accountReceived, err := accountUC.GetAccountByID(account.ID)

		//assert
		assert.Equal(t, entities.Account{}, accountReceived)
		assert.Error(t, err, "failed to retrieve the account from repository")
	})

}

func TestGetAccountByCPF(t *testing.T) {
	t.Run("the account was found via cpf and no error occurred", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		account := entities.Account{
			Name:      "João Paixão",
			ID:        "8aecf60b-b549-41a2-b9b9-143d2d513c87",
			Secret:    "123abc",
			CPF:       "86419560004",
			Balance:   100,
			CreatedAt: "",
		}
		mockRepo.On("FindOne").Return(account, nil)
		accountUC := usecase.NewAccountUc(mockRepo, log)

		//test
		accountReceived, err := accountUC.GetAccountByCPF(account.CPF)
		
		//assert
		assert.Equal(t, account, accountReceived)
		assert.Nil(t, err)
	})
	t.Run("the cpf was sent for the account to be found, but there was an error finding it", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		mockRepo.On("FindOne").Return(entities.Account{}, errors.New(""))
		accountUC := usecase.NewAccountUc(mockRepo, log)
		account := entities.Account{
			Name:      "João Paixão",
			ID:        "8aecf60b-b549-41a2-b9b9-143d2d513c87",
			Secret:    "123abc",
			CPF:       "86419560004",
			Balance:   100,
			CreatedAt: "",
		}
		//test
		accountReceived, err := accountUC.GetAccountByCPF(account.CPF)

		//assert
		assert.Equal(t, entities.Account{}, accountReceived)
		assert.Error(t, err, "failed to retrieve the account from repository")
	})
}
