package account

import (
	"github.com/francisleide/ChallangeGo/domain/entities"
)

type UseCase interface {
	CreateAccount(account entities.AccountInput) (*entities.Account, error)
	ListAllAccounts() []entities.Account
	Deposite(CPF string, amount float64)
	WithDraw(CPF string, amount float64) bool
}

type Repository interface {
	FindOne(CPF string) entities.Account
	UpdateBalance(account entities.Account)
	InsertAccount(accountInput entities.AccountInput) (*entities.Account, error)
	ListAllAccounts() []entities.Account
}
