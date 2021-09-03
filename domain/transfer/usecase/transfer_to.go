package usecase

import (
	"errors"

	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/francisleide/ChallengeGo/domain/transfer"
)

type TransferUc struct {
	r transfer.Repository
}

func NewTransferUC(repo transfer.Repository) TransferUc {
	return TransferUc{
		r: repo,
	}
}

func (t TransferUc) CreateTransfer(accountOriginID, accountDestinationID string, amount float64) (entities.Transfer, error) {
	accountOrigin, errFind := t.r.FindOne(accountOriginID)
	if errFind != nil {
		//TODO: add a sentinel
		return entities.Transfer{}, errors.New("origin account not found")
	}
	accountDestination, errFindID := t.r.FindByID(accountDestinationID)
	if errFindID != nil {
		//TODO: add a sentinel
		return entities.Transfer{}, errors.New("origin destine not found")
	}

	if accountOrigin.Balance >= amount {
		transfer, err := entities.NewTransfer(accountOrigin.ID, accountDestinationID, amount)
		if err != nil {
			//TODO: add a sentinel
			return entities.Transfer{}, errors.New("invalid transfer")
		}
		accountOrigin.Balance -= amount
		accountDestination.Balance += amount
		t.r.UpdateBalance(accountOrigin)
		t.r.UpdateBalance(accountDestination)

		tr, err := t.r.InsertTransfer(transfer)
		if err != nil {
			return entities.Transfer{}, err
		}
		return tr, nil
	} else {
		//TODO add a sentinel
		return entities.Transfer{}, errors.New("insufficient funds")
	}

}

func (t TransferUc) ListUserTransfers(CPF string) ([]entities.Transfer, error) {
	transfer, error := t.r.FindOne(CPF)
	if error != nil {
		//TODO: add a sentinel
		return []entities.Transfer{}, errors.New("account not found")
	}
	transfers, err := t.r.ListUserTransfers(transfer.ID)
	if err != nil {
		return []entities.Transfer{}, err
	}
	return transfers, nil
}
