package usecase_test

import (
	"errors"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/auth"
	"github.com/francisleide/ChallengeGo/domain/auth/usecase"
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	t.Run("username and password are correct and no error is returned", func(t *testing.T) {
		//prepare
		mockRepo := new(auth.MockRepository)
		authUC := usecase.NewAuthenticationUC(mockRepo)
		accountOrigin := usecase.Credentials{
			CPF:    "47708141001",
			Secret: "abc123",
		}
		secret, _ := bcrypt.GenerateFromPassword([]byte(accountOrigin.Secret), bcrypt.DefaultCost)
		accountReceived := entities.Account{
			CPF:    accountOrigin.CPF,
			Name:   "Pereira Silveira",
			Secret: string(secret),
		}

		mockRepo.On("FindOne").Return(accountReceived, nil)

		//test
		err := authUC.Login(accountOrigin.CPF, accountOrigin.Secret)

		//assert
		assert.Equal(t, accountOrigin.CPF, accountReceived.CPF)
		assert.NoError(t, err)

	})
	t.Run("incorrect cpf and login is not performed", func(t *testing.T) {
		//prepare
		mockRepo := new(auth.MockRepository)
		authUC := usecase.NewAuthenticationUC(mockRepo)
		accountOrigin := usecase.Credentials{
			CPF:    "12345678910",
			Secret: "abc123",
		}

		mockRepo.On("FindOne").Return(entities.Account{}, errors.New(""))

		//test
		err := authUC.Login(accountOrigin.CPF, accountOrigin.Secret)

		//assert
		assert.Equal(t, "invalid CPF", err.Error())
	})
	t.Run("incorrect password and login is not performed", func(t *testing.T) {
		//prepare
		mockRepo := new(auth.MockRepository)
		authUC := usecase.NewAuthenticationUC(mockRepo)
		accountOrigin := usecase.Credentials{
			CPF:    "12345678910",
			Secret: "abc123",
		}
		secret, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		accountReceived := entities.Account{
			CPF:    accountOrigin.CPF,
			Name:   "Pereira Silveira",
			Secret: string(secret),
		}

		mockRepo.On("FindOne").Return(accountReceived, nil)

		//test
		err := authUC.Login(accountOrigin.CPF, accountOrigin.Secret)

		//assert
		assert.Equal(t, "incorrect password", err.Error())
	})
}

func TestCreateToken(t *testing.T) {
	t.Run("login is valid and token is created", func(t *testing.T) {
		//prepare
		mockRepo := new(auth.MockRepository)
		authUC := usecase.NewAuthenticationUC(mockRepo)
		accountOrigin := usecase.Credentials{
			CPF:    "12345678910",
			Secret: "abc123",
		}
		secret, _ := bcrypt.GenerateFromPassword([]byte(accountOrigin.Secret), bcrypt.DefaultCost)
		accountReceived := entities.Account{
			CPF:    accountOrigin.CPF,
			Name:   "Pereira Silveira",
			Secret: string(secret),
		}
		mockRepo.On("FindOne").Return(accountReceived, nil)

		//test
		token, err := authUC.CreateToken(accountOrigin.CPF, accountOrigin.Secret)

		//assert
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})
	t.Run("invalid login and token is not created", func(t *testing.T) {
		//prepare
		mockRepo := new(auth.MockRepository)
		authUC := usecase.NewAuthenticationUC(mockRepo)
		accountOrigin := usecase.Credentials{
			CPF:    "12345678910",
			Secret: "abc123",
		}
		mockRepo.On("FindOne").Return(entities.Account{}, errors.New(""))

		//test
		token, err := authUC.CreateToken(accountOrigin.CPF, accountOrigin.Secret)

		//assert
		assert.Equal(t, "invalid login", err.Error())
		assert.Empty(t, token)
	})

}
