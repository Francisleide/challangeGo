package usecase_test

import (
	"errors"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/account/usecase"
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateAccount(t *testing.T) {
	t.Run("cpf does not exist in the bank and the account is created", func(t *testing.T) {
		//prepare
		mockRepo := new(account.MockRepository)
		accountUC := usecase.NewAccountUc(mockRepo, nil)
		accountInput := entities.AccountInput{
			CPF:    "47708141001",
			Secret: "abc123",
			Name:   "Paulo Saulo",
		}
		mockRepo.On("FindOne").Return(entities.Account{}, errors.New(""))
		mockRepo.On("InsertAccount").Return(nil)

		//test
		accountReceived, err := accountUC.CreateAccount(accountInput)

		//assert
		assert.NoError(t, err)
		assert.Equal(t, accountInput.CPF, accountReceived.CPF)
		assert.Equal(t, accountInput.Name, accountReceived.Name)
		err = bcrypt.CompareHashAndPassword([]byte(accountReceived.Secret), []byte(accountInput.Secret))
		assert.NoError(t, err)

	})

	t.Run("cpf exists in the bank and the account is not created", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		accountUC := usecase.NewAccountUc(mockRepo, log)
		account := entities.Account{
			CPF:    "47708141001",
			Name:   "Pereira Silveira",
			Secret: "abc123",
		}
		accountInput := entities.AccountInput{
			CPF:    account.CPF,
			Name:   account.Name,
			Secret: account.Secret,
		}
		mockRepo.On("FindOne").Return(account, nil)
		mockRepo.On("InsertAccount").Return(nil)

		//test
		accountReceived, err := accountUC.CreateAccount(accountInput)

		//assert
		assert.Equal(t, "the account already exists", err.Error())
		assert.Equal(t, entities.Account{}, accountReceived)
	})
	t.Run("the cpf is invalid and the account is not created", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		accountUC := usecase.NewAccountUc(mockRepo, log)
		accountInput := entities.AccountInput{
			CPF:    "12345678901",
			Name:   "Ronaldo Furtado",
			Secret: "abc123",
		}
		mockRepo.On("FindOne").Return(entities.Account{}, errors.New(""))
		mockRepo.On("InsertAccount").Return(nil)

		//test
		accountReceived, err := accountUC.CreateAccount(accountInput)

		//assert
		assert.Equal(t, "invalid cpf", err.Error())
		assert.Equal(t, entities.Account{}, accountReceived)
	})
	t.Run("the cpf is invalid and the account is not created", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		accountUC := usecase.NewAccountUc(mockRepo, log)

		accountInput := entities.AccountInput{
			CPF:    "63597331025",
			Name:   "Ronaldo Furtado",
			Secret: "123",
		}
		mockRepo.On("FindOne").Return(entities.Account{}, errors.New(""))
		mockRepo.On("InsertAccount").Return(nil)

		//test
		accountReceived, err := accountUC.CreateAccount(accountInput)

		//assert
		assert.Equal(t, "invalid secret", err.Error())
		assert.Equal(t, entities.Account{}, accountReceived)
	})
	t.Run("the cpf is invalid and the account is not created", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(account.MockRepository)
		accountUC := usecase.NewAccountUc(mockRepo, log)
		accountInput := entities.AccountInput{
			CPF:    "63597331025",
			Name:   "",
			Secret: "abc123",
		}
		mockRepo.On("FindOne").Return(entities.Account{}, errors.New(""))
		mockRepo.On("InsertAccount").Return(nil)

		//test
		accountReceived, err := accountUC.CreateAccount(accountInput)

		//assert
		assert.Equal(t, "the name cannot be null", err.Error())
		assert.Equal(t, entities.Account{}, accountReceived)
	})

}
