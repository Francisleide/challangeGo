package account_test

import (
	"bytes"
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

func TestCreateAccount(t *testing.T) {
	r := mux.NewRouter()
	t.Run("data is passed correctly and 200 is returned", func(t *testing.T) {
		accountInput := a.AccountInput{
			Name:   "Dávila da Vila",
			CPF:    "63597331025",
			Secret: "abc123",
		}
		newAccount, _ := entities.NewAccount(accountInput.Name, accountInput.CPF, accountInput.Secret)
		requestBody, _ := json.Marshal(accountInput)
		req := bytes.NewReader(requestBody)

		resp, _ := json.Marshal(newAccount)

		usecaseFake := new(account.UsecaseMock)
		usecaseFake.On("CreateAccount").Return(newAccount, nil)
		log := logrus.NewEntry(logrus.New())
		handler := a.Accounts(r, usecaseFake, log)

		request := httptest.NewRequest("Post", "/accounts", req)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.CreateAccount).ServeHTTP(response, request)

		assert.Equal(t, string(resp), strings.TrimSpace(response.Body.String()))
		assert.Equal(t, http.StatusOK, response.Result().StatusCode)

	})
}
