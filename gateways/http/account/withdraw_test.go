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
	"github.com/francisleide/ChallengeGo/domain/entities"
	a "github.com/francisleide/ChallengeGo/gateways/http/account"
	"github.com/francisleide/ChallengeGo/gateways/http/middleware"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestWithdraw(t *testing.T) {
	r := mux.NewRouter()
	t.Run("The withdrawal amount is valid and 200 is returned", func(t *testing.T) {
		//prepare
		withdrawInput := a.Withdraw{
			Amount: 200,
		}
		withdrawOutput := entities.TransactionOutput{
			ID:              "34add062-ccf4-4530-976d-da7b193a4db4",
			PreviousBalance: 300,
			ActualBalance:   100,
		}
		resp, _ := json.Marshal(withdrawOutput)
		requestBody, _ := json.Marshal(withdrawInput)
		req := bytes.NewReader(requestBody)
		usecaseFake := new(account.UsecaseMock)
		usecaseFake.On("Withdraw").Return(withdrawOutput, nil)
		usecaseFake.On("GetCPF").Return("86419560004", true)
		log := logrus.NewEntry(logrus.New())
		handler := a.Accounts(r, usecaseFake, log)
		request := httptest.NewRequest("Post", "/withdraw", req)
		ctx := context.WithValue(request.Context(), middleware.ContextID, "86419560004")
		response := httptest.NewRecorder()

		//test
		http.HandlerFunc(handler.Withdraw).ServeHTTP(response, request.WithContext(ctx))

		//assert
		assert.Equal(t, http.StatusOK, response.Result().StatusCode)
		assert.Equal(t, string(resp), strings.TrimSpace(response.Body.String()))
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))

	})

	t.Run("he token is not valid and the withdrawal does not happen returning 401", func(t *testing.T) {
		//prepare
		withdrawInput := a.Withdraw{
			Amount: 200,
		}
		requestBody, _ := json.Marshal(withdrawInput)
		req := bytes.NewReader(requestBody)
		usecaseFake := new(account.UsecaseMock)
		usecaseFake.On("Withdraw").Return(entities.TransactionOutput{}, errors.New(""))
		usecaseFake.On("GetCPF").Return("", false)
		log := logrus.NewEntry(logrus.New())
		handler := a.Accounts(r, usecaseFake, log)
		request := httptest.NewRequest("Post", "/withdraw", req)
		ctx := context.WithValue(request.Context(), middleware.ContextID, "")
		response := httptest.NewRecorder()

		//test
		http.HandlerFunc(handler.Withdraw).ServeHTTP(response, request.WithContext(ctx))

		//assert
		assert.Equal(t, http.StatusUnauthorized, response.Result().StatusCode)
	})
	t.Run("the request data is invalid and error 400 is returned", func(t *testing.T) {
		//prepare
		withdrawInput := a.Withdraw{
			Amount: -200,
		}
		requestBody, _ := json.Marshal(withdrawInput)
		req := bytes.NewReader(requestBody)
		usecaseFake := new(account.UsecaseMock)
		usecaseFake.On("Withdraw").Return(entities.TransactionOutput{}, errors.New(""))
		usecaseFake.On("GetCPF").Return("86419560004", true)
		log := logrus.NewEntry(logrus.New())
		handler := a.Accounts(r, usecaseFake, log)
		request := httptest.NewRequest("Post", "/withdraw", req)
		ctx := context.WithValue(request.Context(), middleware.ContextID, "86419560004")
		response := httptest.NewRecorder()

		//test
		http.HandlerFunc(handler.Withdraw).ServeHTTP(response, request.WithContext(ctx))

		//assert
		assert.Equal(t, http.StatusInternalServerError, response.Result().StatusCode)
	})
}
