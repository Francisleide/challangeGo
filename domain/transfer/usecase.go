package transfer

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

type UseCase interface {
	CreateTransfer(accountOrigin, accountDestination entities.Account, amount float64) (entities.Transfer, error)
	ListUserTransfers(CPF string) ([]entities.Transfer, error)
}

type Repository interface {
	InsertTransfer(transfer entities.Transfer) (entities.Transfer, error)
	FindOne(CPF string) (entities.Account, error) //RETIRAR
	UpdateBalance(ID string, balance float64) error
	FindByID(accountID string) (entities.Account, error) //RETIRAR
	ListUserTransfers(CPF string) ([]entities.Transfer, error)
}
