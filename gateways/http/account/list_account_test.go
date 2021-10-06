package account_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/entities"
	a "github.com/francisleide/ChallengeGo/gateways/http/account"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestListAccount(t *testing.T){
	r := mux.NewRouter()
	t.Run("the list of accounts is displayed and 200 is returned", func(t *testing.T) {
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

		usecaseFake := new(account.UsecaseMock)
		usecaseFake.On("ListAll").Return(accounts)
		log := logrus.NewEntry(logrus.New())
		handler := a.Accounts(r, usecaseFake, log)

		request := httptest.NewRequest("Get", "/accounts", nil)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Deposit).ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Result().StatusCode)
	})
}