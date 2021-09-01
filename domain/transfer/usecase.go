package transfer

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

type UseCase interface {
	CreateTransfer(origin, destine string, amount float64) (entities.Transfer, error)
}

type Repository interface {
	InsertTransfer(transfer entities.Transfer) (entities.Transfer, error)
	FindOne(CPF string) (entities.Account, bool)
	UpdateBalance(account entities.Account) bool
	FindByID(accountID string) (entities.Account, bool)
}
