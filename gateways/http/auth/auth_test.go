package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/auth"
	a "github.com/francisleide/ChallengeGo/gateways/http/auth"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAuthentication(t *testing.T) {
	r := mux.NewRouter()
	t.Run("login and password data are passed correctly and 200 is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		account := a.Login{
			CPF:    "86594861026",
			Secret: "123abc",
		}
		requestBody, _ := json.Marshal(account)
		req := bytes.NewReader(requestBody)
		usecaseFake := new(auth.UsecaseMock)
		usecaseFake.On("CreateToken").Return("valid-token", nil)
		handler := a.Auth(r, usecaseFake, log)
		request := httptest.NewRequest("Post", "/login", req)
		response := httptest.NewRecorder()

		//test
		http.HandlerFunc(handler.Authentication).ServeHTTP(response, request)

		//assert
		assert.Equal(t, http.StatusOK, response.Result().StatusCode)
		assert.NotEmpty(t, response)
	})
}
