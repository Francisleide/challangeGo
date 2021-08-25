package transfer

import (
	"time"

	"github.com/francisleide/ChallengeGo/domain/entities"
)

type UseCase interface {
	CreateTransfer(origin, destine string, amount float64) (*entities.Transfer, error)
}

type TransferInput struct {
	DestinationCPF string
	Amount         float64
}

type Repository interface {
	InsertTransfer(accountOrigin, DestinationAccount entities.Account, amount float64) (*entities.Transfer, error)
	FindOne(CPF string) entities.Account
}

func NewTransferInput(accountOriginID, accountDestinationID string, amount float64) entities.Transfer {
	return entities.Transfer{
		ID:               entities.GenerateID(),
		OriginAccountID:  accountOriginID,
		DestineAccountID: accountDestinationID,
		Amount:           amount,
		CreatedAt:        time.Now().Format(time.RFC822),
	}
}
