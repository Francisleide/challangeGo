package usecase_test

import (
	"errors"
	"testing"

	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/francisleide/ChallengeGo/domain/transfer"
	"github.com/francisleide/ChallengeGo/domain/transfer/usecase"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransfer(t *testing.T) {
	t.Run("the accounts exist and there is balance for transfer so it occurs", func(t *testing.T) {
		//prepare
		mockRepo := new(transfer.MockRepository)
		transferUC := usecase.NewTransferUC(mockRepo, nil)
		accountOrigin := entities.Account{
			ID:      "2ab7195f-222a-45c3-9189-4f5da5cd745f",
			CPF:     "47708141001",
			Secret:  "abc123",
			Name:    "Paulo Saulo",
			Balance: 2000,
		}
		accountDestination := entities.Account{
			ID:      "34add062-ccf4-4530-976d-da7b193a4db4",
			CPF:     "47708141001",
			Name:    "Pereira Silveira",
			Secret:  "abc123",
			Balance: 0,
		}
		transferExpected := entities.Transfer{
			AccountOriginID:      "2ab7195f-222a-45c3-9189-4f5da5cd745f",
			AccountDestinationID: "34add062-ccf4-4530-976d-da7b193a4db4",
			Amount:               200,
		}
		mockRepo.On("InsertTransfer").Return(nil)

		//test
		transferReceived, err := transferUC.CreateTransfer(accountOrigin, accountDestination, 200)

		//assert
		assert.Equal(t, transferExpected.AccountDestinationID, transferReceived.AccountDestinationID)
		assert.Equal(t, transferExpected.AccountOriginID, transferReceived.AccountOriginID)
		assert.Equal(t, transferExpected.Amount, transferReceived.Amount)
		assert.NotEmpty(t, transferReceived.ID)
		assert.NotEmpty(t, transferReceived.CreatedAt)
		assert.NoError(t, err)

	})
	t.Run("transfer should not occur due to database problem", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(transfer.MockRepository)
		transferUC := usecase.NewTransferUC(mockRepo, log)
		accountOrigin := entities.Account{
			ID:      "2ab7195f-222a-45c3-9189-4f5da5cd745f",
			CPF:     "47708141001",
			Secret:  "abc123",
			Name:    "Paulo Saulo",
			Balance: 2000,
		}
		accountDestination := entities.Account{
			ID:      "34add062-ccf4-4530-976d-da7b193a4db4",
			CPF:     "47708141001",
			Name:    "Pereira Silveira",
			Secret:  "abc123",
			Balance: 0,
		}
		mockRepo.On("InsertTransfer").Return(errors.New(""))

		//test
		transferReceived, err := transferUC.CreateTransfer(accountOrigin, accountDestination, 200)

		//assert
		assert.Equal(t, usecase.ErrorSaveTransfer, err)
		assert.Equal(t, entities.Transfer{}, transferReceived)

	})
	t.Run("the source account does not have enough balance and the transfer does not take place", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(transfer.MockRepository)
		transferUC := usecase.NewTransferUC(mockRepo, log)
		accountOrigin := entities.Account{
			ID:      "2ab7195f-222a-45c3-9189-4f5da5cd745f",
			CPF:     "47708141001",
			Secret:  "abc123",
			Name:    "Paulo Saulo",
			Balance: 10,
		}
		accountDestination := entities.Account{
			ID:      "34add062-ccf4-4530-976d-da7b193a4db4",
			CPF:     "47708141001",
			Name:    "Pereira Silveira",
			Secret:  "abc123",
			Balance: 0,
		}
		mockRepo.On("InsertTransfer").Return(nil)

		//test
		transferReceived, err := transferUC.CreateTransfer(accountOrigin, accountDestination, 200)

		//assert
		assert.Equal(t, entities.Transfer{}, transferReceived)
		assert.Equal(t, usecase.ErrorInsufficientFunds, err)
	})
}

func TestListUserTransfers(t *testing.T) {
	t.Run("a valid account is fetched and the list of all transfers is returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(transfer.MockRepository)
		transferUC := usecase.NewTransferUC(mockRepo, log)
		transferExpected := []entities.Transfer{
			{
				AccountOriginID:      "2ab7195f-222a-45c3-9189-4f5da5cd745f",
				AccountDestinationID: "34add062-ccf4-4530-976d-da7b193a4db4",
				Amount:               200,
			},
			{
				AccountOriginID:      "2ab7195f-222a-45c3-9189-4f5da5cd745f",
				AccountDestinationID: "a3067ef2-74af-4df0-a379-48a81b256ba9",
				Amount:               100,
			},
		}

		mockRepo.On("ListUserTransfers").Return(transferExpected, nil)

		//test
		transfersReceived, err := transferUC.ListUserTransfers("2ab7195f-222a-45c3-9189-4f5da5cd745f")

		//assert
		assert.Equal(t, transferExpected, transfersReceived)
		assert.NoError(t, err)
	})
	t.Run("an invalid account is fetched and the list of all transfers is not returned", func(t *testing.T) {
		//prepare
		log := logrus.NewEntry(logrus.New())
		mockRepo := new(transfer.MockRepository)
		transferUC := usecase.NewTransferUC(mockRepo, log)
		mockRepo.On("ListUserTransfers").Return([]entities.Transfer{}, errors.New(""))

		//test
		transfersReceived, err := transferUC.ListUserTransfers("2ab7195f-222a-45c3-9189-4f5da5cd745f")

		//assert
		assert.Equal(t, usecase.ErrorSaveTransfer, err)
		assert.Equal(t, []entities.Transfer{}, transfersReceived)
	})
}
