package account_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/account"
	a "github.com/francisleide/ChallengeGo/gateways/http/account"
	"github.com/francisleide/ChallengeGo/gateways/http/middleware"
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
		depositOutput := account.TransactionOutput{
			ID:              "2ab7195f-222a-45c3-9189-4f5da5cd745f",
			PreviousBalance: 500,
			ActualBalance:   700,
		}
		usecaseFake := new(account.UsecaseMock)
		usecaseFake.On("Deposit").Return(depositOutput, nil)
		usecaseFake.On("GetCPF").Return("86419560004", true)
		requestBody, _ := json.Marshal(depositInput)
		req := bytes.NewReader(requestBody)
		resp, _ := json.Marshal(depositOutput)
		log := logrus.NewEntry(logrus.New())
		request := httptest.NewRequest("Post", "/deposit", req)
		ctx := context.WithValue(request.Context(), middleware.ContextID, "86419560004")
		response := httptest.NewRecorder()
		handler := a.Accounts(r, usecaseFake, log)

		//test
		http.HandlerFunc(handler.Deposit).ServeHTTP(response, request.WithContext(ctx))

		//assert
		assert.Equal(t, http.StatusCreated, response.Result().StatusCode)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
		assert.Equal(t, string(resp), strings.TrimSpace(response.Body.String()))

	})
	t.Run("the token is not valid and the deposit does not happen returning 401", func(t *testing.T) {
		//prepare
		depositInput := a.DepositInput{
			Amount: 200,
		}
		usecaseFake := new(account.UsecaseMock)
		usecaseFake.On("Deposit").Return(account.TransactionOutput{}, errors.New(""))
		usecaseFake.On("GetCPF").Return("", false)
		requestBody, _ := json.Marshal(depositInput)
		req := bytes.NewReader(requestBody)
		log := logrus.NewEntry(logrus.New())
		request := httptest.NewRequest("Post", "/deposit", req)
		ctx := context.WithValue(request.Context(), middleware.ContextID, "")
		response := httptest.NewRecorder()
		handler := a.Accounts(r, usecaseFake, log)

		//test
		http.HandlerFunc(handler.Deposit).ServeHTTP(response, request.WithContext(ctx))

		//assert
		assert.Equal(t, http.StatusUnauthorized, response.Result().StatusCode)
	})
	t.Run("the request data is invalid and error 400 is returned", func(t *testing.T) {
		//prepare
		depositInput := a.DepositInput{
			Amount: -200,
		}

		usecaseFake := new(account.UsecaseMock)
		usecaseFake.On("Deposit").Return(account.TransactionOutput{}, errors.New(""))
		usecaseFake.On("GetCPF").Return("86419560004", true)
		requestBody, _ := json.Marshal(depositInput)
		req := bytes.NewReader(requestBody)
		log := logrus.NewEntry(logrus.New())
		request := httptest.NewRequest("Post", "/deposit", req)
		ctx := context.WithValue(request.Context(), middleware.ContextID, "86419560004")
		response := httptest.NewRecorder()
		handler := a.Accounts(r, usecaseFake, log)

		//test
		http.HandlerFunc(handler.Deposit).ServeHTTP(response, request.WithContext(ctx))

		//assert
		assert.Equal(t, http.StatusInternalServerError, response.Result().StatusCode)
	})
}
