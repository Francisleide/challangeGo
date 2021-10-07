package account_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/account"
	a "github.com/francisleide/ChallengeGo/gateways/http/account"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGetBalance(t *testing.T){
	r := mux.NewRouter()
	t.Run("the account id is valid and the balance and status 200 is returned", func(t *testing.T) {
		//prepare
		accountID := "efb711fa-786e-4d57-9eeb-6fbaca8775a9"
		balance := 200.0
		usecaseFake := new(account.UsecaseMock)
		usecaseFake.On("GetBalance").Return(balance, nil)
		log := logrus.NewEntry(logrus.New())
		handler := a.Accounts(r, usecaseFake, log)
		path := fmt.Sprintf("/accounts/%s/balance", accountID)
		request := httptest.NewRequest("Get", path, nil)
		response := httptest.NewRecorder()

		//test
		http.HandlerFunc(handler.GetBalance).ServeHTTP(response, request)

		//assert
		assert.Equal(t, http.StatusOK, response.Result().StatusCode)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
	})
}