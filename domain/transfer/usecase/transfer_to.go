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

//conta de origem entities.Account
//conta de destino entities.Account
func (t TransferUc) CreateTransfer(accountOrigin, accountDestination entities.Account, amount float64) (entities.Transfer, error) {
	if accountOrigin.Balance >= amount {
		transfer, err := entities.NewTransfer(accountOrigin.ID, accountDestination.ID, amount)
		if err != nil {
			//TODO: add a sentinel
			return entities.Transfer{}, errors.New("invalid transfer")
		}
		//accountOrigin.Balance -= amount
		//accountDestination.Balance += amount
		//ja foi feito no handler
		//t.r.UpdateBalance(accountOrigin.ID, accountOrigin.Balance)
		//t.r.UpdateBalance(accountDestination.ID, accountDestination.Balance)

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

func (t TransferUc) ListUserTransfers(accountID string) ([]entities.Transfer, error) {
	//transfer, error := t.r.FindOne(CPF)
/*	if error != nil {
		//TODO: add a sentinel
		return []entities.Transfer{}, errors.New("account not found")
	}*/
	transfers, err := t.r.ListUserTransfers(accountID)
	if err != nil {
		return []entities.Transfer{}, err
	}
	return transfers, nil
}
