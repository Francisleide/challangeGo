package transfer_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/francisleide/ChallengeGo/domain/transfer"
	tr "github.com/francisleide/ChallengeGo/gateways/http/transfer"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransfer(t *testing.T) {
	r := mux.NewRouter()
	t.Run("the data was passed correctly and status 200 is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		accountOrigin := entities.Account{
			ID:      "efb711fa-786e-4d57-9eeb-6fbaca8775a9",
			Name:    "Ana Mariana",
			Balance: 100,
		}
		accountDestination := entities.Account{
			ID:      "8b27748e-88a8-4792-b22a-67ba8f77179f",
			Name:    "Flávio Sábio",
			Balance: 0,
		}
		transferInput := tr.TransferInput{
			AccountDestinationID: "280160d0-5c66-4faf-97c8-1512f90da152",
			Amount:               200,
		}

		requestBody, _ := json.Marshal(transferInput)
		req := bytes.NewReader(requestBody)
		usecaseFakeAccount := new(account.UsecaseMock)
		usecaseFakeTransfer := new(transfer.UsecaseMock)
		usecaseFakeAccount.On("GetAccountByCPF").Return(accountOrigin, nil)
		usecaseFakeAccount.On("GetAccountByCPF").Return(accountDestination, nil)
		usecaseFakeTransfer.On("CreateTransfer").Return()
		handler := tr.NewTransfer(r, usecaseFakeTransfer, usecaseFakeAccount, log)
		request := httptest.NewRequest("Post", "/transfer", req)
		response := httptest.NewRecorder()

		//test
		http.HandlerFunc(handler.CreateTransfer).ServeHTTP(response, request)

		//test
		assert.Equal(t, http.StatusOK, response.Result().StatusCode)
	})
}
