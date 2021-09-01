package entities

import (
	"errors"
	"time"

	"github.com/satori/uuid.go"
)

type Transfer struct {
	ID                   string
	OriginAccountID      string
	DestinationAccountID string
	Amount               float64
	CreatedAt            string
}

type TransferInput struct {
	DestinationAccountID string
	Amount               float64
}

func AmountValidation(amount float64) bool {
	return amount > 0

}

func NewTransfer(accountOriginID, accountDestinationID string, amount float64) (Transfer, error) {
	if !AmountValidation(amount) {
		//TODO: add a sentinel
		return Transfer{}, errors.New("invalid amount")
	}
	return Transfer{
		ID:                   uuid.NewV4().String(),
		OriginAccountID:      accountOriginID,
		DestinationAccountID: accountDestinationID,
		Amount:               amount,
		CreatedAt:            time.Now().Format(time.RFC822),
	}, nil

}
