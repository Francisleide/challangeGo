package account_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/account"
	a "github.com/francisleide/ChallengeGo/gateways/http/account"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestDeposit(t *testing.T) {
	r := mux.NewRouter()
	t.Run("the amount to be deposited is a valid amount and 200 is returned", func(t *testing.T) {
		//prepare
		depositInput := a.DepositInput{
			Amount: 200,
		}
		requestBody, _ := json.Marshal(depositInput)
		req := bytes.NewReader(requestBody)
		usecaseFake := new(account.UsecaseMock)
		usecaseFake.On("Deposit").Return(nil)
		log := logrus.NewEntry(logrus.New())
		handler := a.Accounts(r, usecaseFake, log)
		request := httptest.NewRequest("Post", "/deposit", req)
		response := httptest.NewRecorder()

		//test
		http.HandlerFunc(handler.Deposit).ServeHTTP(response, request)

		//assert
		assert.Equal(t, http.StatusOK, response.Result().StatusCode)

	})
}
