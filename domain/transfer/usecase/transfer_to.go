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

//
func (t TransferUc) CreateTransfer(origin, destine string, amount float64) (*entities.Transfer, error) {
	var originAccount entities.Account
	var destineAccount entities.Account
	originAccount = t.r.FindOne(origin)
	destineAccount = t.r.FindOne(destine)
	if originAccount.Balance >= amount {
		var tr *entities.Transfer
		originAccount.Balance -= amount
		destineAccount.Balance += amount
		tr, err := t.r.InsertTransfer(originAccount, destineAccount, amount)
		if err != nil {
			return nil, err
		}
		return tr, nil
	} else {
		return nil, errors.New("Insufficient funds")
	}
}
