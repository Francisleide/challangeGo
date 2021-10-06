package transfer

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

type UseCase interface {
	CreateTransfer(accountOrigin, accountDestination entities.Account, amount float64) (entities.Transfer, error)
	ListUserTransfers(CPF string) ([]entities.Transfer, error)
}

type Repository interface {
	InsertTransfer(transfer entities.Transfer) error
	UpdateBalance(ID string, balance float64) error
	ListUserTransfers(CPF string) ([]entities.Transfer, error)
}
