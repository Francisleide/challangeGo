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

func TestWithdraw(t *testing.T) {
	r := mux.NewRouter()
	t.Run("the amount to be deposited is a valid amount and 200 is returned", func(t *testing.T) {
		withdrawInput := a.Withdraw{
			Amount: 200,
		}

		requestBody, _ := json.Marshal(withdrawInput)
		req := bytes.NewReader(requestBody)

		usecaseFake := new(account.UsecaseMock)
		usecaseFake.On("Withdraw").Return(nil)
		log := logrus.NewEntry(logrus.New())
		handler := a.Accounts(r, usecaseFake, log)

		request := httptest.NewRequest("Post", "/withdraw", req)

		response := httptest.NewRecorder()

		http.HandlerFunc(handler.Deposit).ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Result().StatusCode)

	})
}
