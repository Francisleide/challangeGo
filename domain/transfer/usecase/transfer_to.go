package usecase

import (
	"errors"

	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/francisleide/ChallengeGo/domain/transfer"
	"github.com/sirupsen/logrus"
)

type TransferUc struct {
	r   transfer.Repository
	log *logrus.Entry
}

func NewTransferUC(repo transfer.Repository, log *logrus.Entry) TransferUc {
	return TransferUc{
		r:   repo,
		log: log,
	}
}

func (t TransferUc) CreateTransfer(accountOrigin, accountDestination entities.Account, amount float64) (entities.Transfer, error) {
	if accountOrigin.Balance >= amount {
		transfer, err := entities.NewTransfer(accountOrigin.ID, accountDestination.ID, amount)
		if err != nil {
			//TODO: add a sentinel
			return entities.Transfer{}, errors.New("invalid transfer")
		}

		err = t.r.InsertTransfer(transfer)
		if err != nil {
			t.log.WithError(err).Error("unable to save transfer")
			return entities.Transfer{}, errors.New("unable to save transfer")
		}
		return transfer , nil
	} else {
		//TODO add a sentinel
		return entities.Transfer{}, errors.New("insufficient funds")
	}

}

func (t TransferUc) ListUserTransfers(accountID string) ([]entities.Transfer, error) {

	transfers, err := t.r.ListUserTransfers(accountID)
	if err != nil {
		return []entities.Transfer{}, errors.New("unable to save transfer")
	}
	return transfers, nil
}
