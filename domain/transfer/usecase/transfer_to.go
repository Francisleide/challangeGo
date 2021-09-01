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
	account, ok := t.r.FindOne(accountOriginID)
	if !ok {
		return entities.Transfer{}, errors.New("account not found")
	}
	accountOrigin, okOrigin := t.r.FindByID(account.ID)
	if !okOrigin {
		//TODO: add a sentinel
		return entities.Transfer{}, errors.New("origin account not found")
	}
	accountDestination, okDestine := t.r.FindByID(accountDestinationID)
	if !okDestine {
		//TODO: add a sentinel
		return entities.Transfer{}, errors.New("origin destine not found")
	}

	if accountOrigin.Balance >= amount {
		//chamar update
		//chamar o update da outra conta
		//create_transfer do entities
		transfer, err := entities.NewTransfer(accountOriginID, accountDestinationID, amount)
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
