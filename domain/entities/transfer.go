package entities

import (
	"errors"
	"time"

	"github.com/satori/uuid.go"
)

type Transfer struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               float64
	CreatedAt            string
}

type TransferInput struct {
	AccountDestinationID string
	Amount               float64
}

func ValidateAmount(amount float64) bool {
	return amount > 0

}

func NewTransfer(accountOriginID, accountDestinationID string, amount float64) (Transfer, error) {
	if !ValidateAmount(amount) {
		//TODO: add a sentinel
		return Transfer{}, errors.New("invalid amount")
	}
	return Transfer{
		ID:                   uuid.NewV4().String(),
		AccountOriginID:      accountOriginID,
		AccountDestinationID: accountDestinationID,
		Amount:               amount,
		CreatedAt:            time.Now().Format(time.RFC822),
	}, nil

}
