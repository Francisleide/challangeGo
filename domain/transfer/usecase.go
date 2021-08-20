package transfer

import (
	"time"

	"github.com/francisleide/ChallangeGo/domain/entities"
)

type UseCase interface {
	CreateTransfer(origem, destino string, ammount float64) (*entities.Transfer, error)
}

type TransferInput struct {
	CPFDestino string
	Amount     float64
}

type Repository interface {
	InsertTransfer(accountOrigem, accountDestino entities.Account, ammount float64) (*entities.Transfer, error)
	FindOne(CPF string) entities.Account
}

func NewTransferInput(accountOriginId, accountDestinationId string, ammount float64) entities.Transfer {
	return entities.Transfer{
		ID:                   entities.GenerateID(),
		AccountOriginID:      accountOriginId,
		AccountDestinationID: accountDestinationId,
		Amount:               ammount,
		CreatedAt:            time.Now().Format(time.RFC822),
	}
}
