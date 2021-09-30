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

func TestListAll(t *testing.T) {
	t.Run("all accounts are recovered without error", func(t *testing.T) {
		//prepare
		mockRepo := new(account.MockRepository)
		accounts := []entities.Account{

			{
				Name:    "Lorena Morena",
				CPF:     "86419560004",
				Balance: 2000,
			},
			{
				Name:    "Margarete de Albuquerque",
				CPF:     "47708141001",
				Balance: 900,
			},
		}
		mockRepo.On("ListAllAccounts").Return(accounts, nil)
		accountUC := usecase.NewAccountUc(mockRepo, nil)

		//test
		receivedAccounts, err := accountUC.ListAll()

		//assert
		assert.Nil(t, err)
		assert.Equal(t, accounts, receivedAccounts)
	})
	t.Run("accounts are not listed and the error is displayed", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		mockRepo.On("ListAllAccounts").Return([]entities.Account{}, errors.New("error"))
		accountUC := usecase.NewAccountUc(mockRepo, log)

		//test
		receivedAccounts, err := accountUC.ListAll()

		//assert
		assert.Equal(t, err.Error(), "failed to list accounts")
		assert.Equal(t, []entities.Account{}, receivedAccounts)
	})

}
