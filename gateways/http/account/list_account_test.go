package account_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/entities"
	a "github.com/francisleide/ChallengeGo/gateways/http/account"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestListAccount(t *testing.T) {
	r := mux.NewRouter()
	t.Run("the list of accounts is displayed and 200 is returned", func(t *testing.T) {
		//prepare
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
		usecaseFake.On("ListAll").Return(accounts, nil)
		log := logrus.NewEntry(logrus.New())
		handler := a.Accounts(r, usecaseFake, log)
		request := httptest.NewRequest("Get", "/accounts", nil)
		response := httptest.NewRecorder()
		accounts2, _ := json.Marshal(accounts)
		//resp := bytes.NewReader(accounts2)

		//test
		http.HandlerFunc(handler.ListAllAccounts).ServeHTTP(response, request)

		//assert
		assert.Equal(t, http.StatusOK, response.Result().StatusCode)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
		assert.Equal(t, string(accounts2), strings.TrimSpace(response.Body.String()))
	})
}
