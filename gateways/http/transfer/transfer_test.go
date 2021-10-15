package transfer_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	ac "github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/francisleide/ChallengeGo/domain/transfer"
	"github.com/francisleide/ChallengeGo/gateways/http/middleware"
	tr "github.com/francisleide/ChallengeGo/gateways/http/transfer"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransfer(t *testing.T) {
	r := mux.NewRouter()
	t.Run("the data was passed correctly and status 201 is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())

		accountOrigin := entities.Account{
			ID:      "efb711fa-786e-4d57-9eeb-6fbaca8775a9",
			CPF:     "86419560004",
			Name:    "Ana Mariana",
			Balance: 100,
		}
		accountDestination := entities.Account{
			ID:      "8b27748e-88a8-4792-b22a-67ba8f77179f",
			CPF:     "63597331025",
			Name:    "Flávio Sábio",
			Balance: 0,
		}
		transferInput := tr.TransferInput{
			AccountDestinationID: "8b27748e-88a8-4792-b22a-67ba8f77179f",
			Amount:               50,
		}
		withdrawOutput := ac.TransactionOutput{
			ID:              "efb711fa-786e-4d57-9eeb-6fbaca8775a9",
			PreviousBalance: 100,
			ActualBalance:   50,
		}
		depositOutput := ac.TransactionOutput{
			ID:              "8b27748e-88a8-4792-b22a-67ba8f77179f",
			PreviousBalance: 0,
			ActualBalance:   50,
		}
		transferOutput := entities.Transfer{
			AccountDestinationID: "8b27748e-88a8-4792-b22a-67ba8f77179f",
			Amount:               50,
			CreatedAt:            "2021-10-07 20:36:08.43653495 +0000 UTC",
		}

		requestBody, _ := json.Marshal(transferInput)
		req := bytes.NewReader(requestBody)
		usecaseFakeAccount := new(ac.UsecaseMock)
		usecaseFakeTransfer := new(transfer.UsecaseMock)
		usecaseFakeAccount.On("GetAccountByCPF").Return(accountOrigin, nil)
		usecaseFakeAccount.On("GetAccountByID").Return(accountDestination, nil)
		usecaseFakeAccount.On("Withdraw").Return(withdrawOutput, nil)
		usecaseFakeAccount.On("Deposit").Return(depositOutput, nil)
		usecaseFakeAccount.On("GetAccountByID").Return(accountDestination, nil)
		usecaseFakeAccount.On("GetCPF").Return("86419560004", true)
		usecaseFakeTransfer.On("CreateTransfer").Return(transferOutput, nil)
		handler := tr.NewTransfer(r, usecaseFakeTransfer, usecaseFakeAccount, log)
		request := httptest.NewRequest("Post", "/transfer", req)
		response := httptest.NewRecorder()
		ctx := context.WithValue(request.Context(), middleware.ContextID, "86419560004")

		//test
		http.HandlerFunc(handler.CreateTransfer).ServeHTTP(response, request.WithContext(ctx))

		//test
		assert.Equal(t, http.StatusCreated, response.Result().StatusCode)
	})
	t.Run("token is invalid and 401 status is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		transferInput := tr.TransferInput{
			AccountDestinationID: "8b27748e-88a8-4792-b22a-67ba8f77179f",
			Amount:               50,
		}

		requestBody, _ := json.Marshal(transferInput)
		req := bytes.NewReader(requestBody)
		usecaseFakeAccount := new(ac.UsecaseMock)
		usecaseFakeTransfer := new(transfer.UsecaseMock)
		usecaseFakeAccount.On("GetAccountByCPF").Return(entities.Account{}, errors.New(""))
		usecaseFakeTransfer.On("CreateTransfer").Return()
		handler := tr.NewTransfer(r, usecaseFakeTransfer, usecaseFakeAccount, log)
		request := httptest.NewRequest("Post", "/transfer", req)
		response := httptest.NewRecorder()

		//test
		http.HandlerFunc(handler.CreateTransfer).ServeHTTP(response, request)

		//test
		assert.Equal(t, http.StatusUnauthorized, response.Result().StatusCode)
	})
	t.Run("the target account was not found and 404 is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())

		accountOrigin := entities.Account{
			ID:      "efb711fa-786e-4d57-9eeb-6fbaca8775a9",
			CPF:     "86419560004",
			Name:    "Ana Mariana",
			Balance: 100,
		}
		accountDestination := entities.Account{
			ID:      "8b27748e-88a8-4792-b22a-67ba8f77179f",
			CPF:     "63597331025",
			Name:    "Flávio Sábio",
			Balance: 0,
		}
		transferInput := tr.TransferInput{
			AccountDestinationID: "8b27748e-88a8-4792-b22a-67ba8f77179f",
			Amount:               50,
		}
		withdrawOutput := ac.TransactionOutput{
			ID:              "efb711fa-786e-4d57-9eeb-6fbaca8775a9",
			PreviousBalance: 100,
			ActualBalance:   50,
		}
		depositOutput := ac.TransactionOutput{
			ID:              "8b27748e-88a8-4792-b22a-67ba8f77179f",
			PreviousBalance: 0,
			ActualBalance:   50,
		}
		transferOutput := entities.Transfer{
			AccountDestinationID: "8b27748e-88a8-4792-b22a-67ba8f77179f",
			Amount:               50,
			CreatedAt:            "2021-10-07 20:36:08.43653495 +0000 UTC",
		}

		requestBody, _ := json.Marshal(transferInput)
		req := bytes.NewReader(requestBody)
		usecaseFakeAccount := new(ac.UsecaseMock)
		usecaseFakeTransfer := new(transfer.UsecaseMock)
		usecaseFakeAccount.On("GetAccountByCPF").Return(accountOrigin, nil)
		usecaseFakeAccount.On("GetAccountByID").Return(entities.Account{}, errors.New(""))
		usecaseFakeAccount.On("Withdraw").Return(withdrawOutput, nil)
		usecaseFakeAccount.On("Deposit").Return(depositOutput, nil)
		usecaseFakeAccount.On("GetAccountByID").Return(accountDestination, nil)
		usecaseFakeAccount.On("GetCPF").Return("86419560004", true)
		usecaseFakeTransfer.On("CreateTransfer").Return(transferOutput, nil)
		handler := tr.NewTransfer(r, usecaseFakeTransfer, usecaseFakeAccount, log)
		request := httptest.NewRequest("Post", "/transfer", req)
		response := httptest.NewRecorder()
		ctx := context.WithValue(request.Context(), middleware.ContextID, "86419560004")

		//test
		http.HandlerFunc(handler.CreateTransfer).ServeHTTP(response, request.WithContext(ctx))

		//test
		assert.Equal(t, http.StatusNotFound, response.Result().StatusCode)
	})
	t.Run("the withdrawal could not be performed and error 500 is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		transferInput := tr.TransferInput{
			AccountDestinationID: "8b27748e-88a8-4792-b22a-67ba8f77179f",
			Amount:               50,
		}

		requestBody, _ := json.Marshal(transferInput)
		req := bytes.NewReader(requestBody)
		usecaseFakeAccount := new(ac.UsecaseMock)
		usecaseFakeTransfer := new(transfer.UsecaseMock)
		usecaseFakeAccount.On("GetAccountByCPF").Return(entities.Account{}, errors.New(""))
		usecaseFakeTransfer.On("CreateTransfer").Return()
		handler := tr.NewTransfer(r, usecaseFakeTransfer, usecaseFakeAccount, log)
		request := httptest.NewRequest("Post", "/transfer", req)
		response := httptest.NewRecorder()

		//test
		http.HandlerFunc(handler.CreateTransfer).ServeHTTP(response, request)

		//test
		assert.Equal(t, http.StatusUnauthorized, response.Result().StatusCode)
	})
	t.Run("The withdrawal could not be performed and error 500 is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())

		accountOrigin := entities.Account{
			ID:      "efb711fa-786e-4d57-9eeb-6fbaca8775a9",
			CPF:     "86419560004",
			Name:    "Ana Mariana",
			Balance: 100,
		}
		accountDestination := entities.Account{
			ID:      "8b27748e-88a8-4792-b22a-67ba8f77179f",
			CPF:     "63597331025",
			Name:    "Flávio Sábio",
			Balance: 0,
		}
		transferInput := tr.TransferInput{
			AccountDestinationID: "8b27748e-88a8-4792-b22a-67ba8f77179f",
			Amount:               50,
		}
		withdrawOutput := ac.TransactionOutput{
			ID:              "efb711fa-786e-4d57-9eeb-6fbaca8775a9",
			PreviousBalance: 100,
			ActualBalance:   50,
		}
		depositOutput := ac.TransactionOutput{
			ID:              "8b27748e-88a8-4792-b22a-67ba8f77179f",
			PreviousBalance: 0,
			ActualBalance:   50,
		}
		transferOutput := entities.Transfer{
			AccountDestinationID: "8b27748e-88a8-4792-b22a-67ba8f77179f",
			Amount:               50,
			CreatedAt:            "2021-10-07 20:36:08.43653495 +0000 UTC",
		}

		requestBody, _ := json.Marshal(transferInput)
		req := bytes.NewReader(requestBody)
		usecaseFakeAccount := new(ac.UsecaseMock)
		usecaseFakeTransfer := new(transfer.UsecaseMock)
		usecaseFakeAccount.On("GetAccountByCPF").Return(accountOrigin, nil)
		usecaseFakeAccount.On("GetAccountByID").Return(accountDestination, nil)
		usecaseFakeAccount.On("Withdraw").Return(withdrawOutput, errors.New(""))
		usecaseFakeAccount.On("Deposit").Return(depositOutput, nil)
		usecaseFakeAccount.On("GetAccountByID").Return(accountDestination, nil)
		usecaseFakeAccount.On("GetCPF").Return("86419560004", true)
		usecaseFakeTransfer.On("CreateTransfer").Return(transferOutput, nil)
		handler := tr.NewTransfer(r, usecaseFakeTransfer, usecaseFakeAccount, log)
		request := httptest.NewRequest("Post", "/transfer", req)
		response := httptest.NewRecorder()
		ctx := context.WithValue(request.Context(), middleware.ContextID, "86419560004")

		//test
		http.HandlerFunc(handler.CreateTransfer).ServeHTTP(response, request.WithContext(ctx))

		//test
		assert.Equal(t, http.StatusInternalServerError, response.Result().StatusCode)
	})

	t.Run("transfer data is incorrect and 400 is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())

		accountOrigin := entities.Account{
			ID:      "efb711fa-786e-4d57-9eeb-6fbaca8775a9",
			CPF:     "86419560004",
			Name:    "Ana Mariana",
			Balance: 100,
		}
		accountDestination := entities.Account{
			ID:      "8b27748e-88a8-4792-b22a-67ba8f77179f",
			CPF:     "63597331025",
			Name:    "Flávio Sábio",
			Balance: 0,
		}
		transferInput := tr.TransferInput{
			AccountDestinationID: "8b27748e-88a8-4792-b22a-67ba8f77179f",
			Amount:               -50,
		}
		withdrawOutput := ac.TransactionOutput{
			ID:              "efb711fa-786e-4d57-9eeb-6fbaca8775a9",
			PreviousBalance: 100,
			ActualBalance:   50,
		}
		depositOutput := ac.TransactionOutput{
			ID:              "8b27748e-88a8-4792-b22a-67ba8f77179f",
			PreviousBalance: 0,
			ActualBalance:   50,
		}

		requestBody, _ := json.Marshal(transferInput)
		req := bytes.NewReader(requestBody)
		usecaseFakeAccount := new(ac.UsecaseMock)
		usecaseFakeTransfer := new(transfer.UsecaseMock)
		usecaseFakeAccount.On("GetAccountByCPF").Return(accountOrigin, nil)
		usecaseFakeAccount.On("GetAccountByID").Return(accountDestination, nil)
		usecaseFakeAccount.On("Withdraw").Return(withdrawOutput, nil)
		usecaseFakeAccount.On("Deposit").Return(depositOutput, nil)
		usecaseFakeAccount.On("GetAccountByID").Return(accountDestination, nil)
		usecaseFakeAccount.On("GetCPF").Return("86419560004", true)
		usecaseFakeTransfer.On("CreateTransfer").Return(entities.Transfer{}, errors.New(""))
		handler := tr.NewTransfer(r, usecaseFakeTransfer, usecaseFakeAccount, log)
		request := httptest.NewRequest("Post", "/transfer", req)
		response := httptest.NewRecorder()
		ctx := context.WithValue(request.Context(), middleware.ContextID, "86419560004")

		//test
		http.HandlerFunc(handler.CreateTransfer).ServeHTTP(response, request.WithContext(ctx))

		//test
		assert.Equal(t, http.StatusBadRequest, response.Result().StatusCode)
	})

}

func TestListUserTransfers(t *testing.T) {
	r := mux.NewRouter()
	t.Run("the data to list transfers is correct and 201 is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		transfers := []entities.Transfer{
			{
				ID:                   "1f7e1200-ec15-4666-bd7d-d842bfc87685",
				AccountOriginID:      "287cb5ca-5d64-4920-9c98-d3b76eaa9ac0",
				AccountDestinationID: "6ce37b8d-a035-49c4-b5a4-d1bb5009bc16",
				Amount:               100,
				CreatedAt:            "2021-09-14 15:16:50.542672248 +0000 UTC",
			},
			{
				ID:                   "77af82ed-2b68-446f-a062-60a3c1786544",
				AccountOriginID:      "287cb5ca-5d64-4920-9c98-d3b76eaa9ac0",
				AccountDestinationID: "7d6816af-9725-47ce-ab45-80455662cefe",
				Amount:               100,
				CreatedAt:            "2021-09-14 15:16:58.542672248 +0000 UTC",
			},
		}
		accountOrigin := entities.Account{
			ID:      "8b27748e-88a8-4792-b22a-67ba8f77179f",
			CPF:     "86419560004",
			Name:    "Flávio Sábio",
			Balance: 0,
		}
		usecaseFake := new(transfer.UsecaseMock)
		usecaseFakeAccount := new(ac.UsecaseMock)
		usecaseFakeAccount.On("GetAccountByCPF").Return(accountOrigin, nil)
		usecaseFake.On("ListUserTransfers").Return(transfers, nil)
		usecaseFakeAccount.On("GetCPF").Return("86419560004", true)

		handler := tr.NewTransfer(r, usecaseFake, usecaseFakeAccount, log)
		request := httptest.NewRequest("Post", "/transfer", nil)
		response := httptest.NewRecorder()
		ctx := context.WithValue(request.Context(), middleware.ContextID, "86419560004")

		//test
		http.HandlerFunc(handler.ListUserTransfers).ServeHTTP(response, request.WithContext(ctx))

		//test
		assert.Equal(t, http.StatusOK, response.Result().StatusCode)

	})
	t.Run("the token is incorrect or does not exist and error 401 is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		transfers := []entities.Transfer{
			{
				ID:                   "1f7e1200-ec15-4666-bd7d-d842bfc87685",
				AccountOriginID:      "287cb5ca-5d64-4920-9c98-d3b76eaa9ac0",
				AccountDestinationID: "6ce37b8d-a035-49c4-b5a4-d1bb5009bc16",
				Amount:               100,
				CreatedAt:            "2021-09-14 15:16:50.542672248 +0000 UTC",
			},
			{
				ID:                   "77af82ed-2b68-446f-a062-60a3c1786544",
				AccountOriginID:      "287cb5ca-5d64-4920-9c98-d3b76eaa9ac0",
				AccountDestinationID: "7d6816af-9725-47ce-ab45-80455662cefe",
				Amount:               100,
				CreatedAt:            "2021-09-14 15:16:58.542672248 +0000 UTC",
			},
		}
		accountOrigin := entities.Account{
			ID:      "8b27748e-88a8-4792-b22a-67ba8f77179f",
			CPF:     "86419560004",
			Name:    "Flávio Sábio",
			Balance: 0,
		}
		usecaseFake := new(transfer.UsecaseMock)
		usecaseFakeAccount := new(ac.UsecaseMock)
		usecaseFakeAccount.On("GetAccountByCPF").Return(accountOrigin, nil)
		usecaseFake.On("ListUserTransfers").Return(transfers, nil)
		usecaseFakeAccount.On("GetCPF").Return("", false)

		handler := tr.NewTransfer(r, usecaseFake, usecaseFakeAccount, log)
		request := httptest.NewRequest("Post", "/transfer", nil)
		response := httptest.NewRecorder()
		ctx := context.WithValue(request.Context(), middleware.ContextID, "")

		//test
		http.HandlerFunc(handler.ListUserTransfers).ServeHTTP(response, request.WithContext(ctx))

		//test
		assert.Equal(t, http.StatusUnauthorized, response.Result().StatusCode)
	})
	t.Run("a problem occurs to retrieve data from transfers and error 500 is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		transfers := []entities.Transfer{
			{
				ID:                   "1f7e1200-ec15-4666-bd7d-d842bfc87685",
				AccountOriginID:      "287cb5ca-5d64-4920-9c98-d3b76eaa9ac0",
				AccountDestinationID: "6ce37b8d-a035-49c4-b5a4-d1bb5009bc16",
				Amount:               100,
				CreatedAt:            "2021-09-14 15:16:50.542672248 +0000 UTC",
			},
			{
				ID:                   "77af82ed-2b68-446f-a062-60a3c1786544",
				AccountOriginID:      "287cb5ca-5d64-4920-9c98-d3b76eaa9ac0",
				AccountDestinationID: "7d6816af-9725-47ce-ab45-80455662cefe",
				Amount:               100,
				CreatedAt:            "2021-09-14 15:16:58.542672248 +0000 UTC",
			},
		}
		accountOrigin := entities.Account{
			ID:      "8b27748e-88a8-4792-b22a-67ba8f77179f",
			CPF:     "86419560004",
			Name:    "Flávio Sábio",
			Balance: 0,
		}
		usecaseFake := new(transfer.UsecaseMock)
		usecaseFakeAccount := new(ac.UsecaseMock)
		usecaseFakeAccount.On("GetAccountByCPF").Return(accountOrigin, nil)
		usecaseFake.On("ListUserTransfers").Return(transfers, errors.New(""))
		usecaseFakeAccount.On("GetCPF").Return("86419560004", true)

		handler := tr.NewTransfer(r, usecaseFake, usecaseFakeAccount, log)
		request := httptest.NewRequest("Post", "/transfer", nil)
		response := httptest.NewRecorder()
		ctx := context.WithValue(request.Context(), middleware.ContextID, "86419560004")

		//test
		http.HandlerFunc(handler.ListUserTransfers).ServeHTTP(response, request.WithContext(ctx))

		//test
		assert.Equal(t, http.StatusInternalServerError, response.Result().StatusCode)
	})
}
