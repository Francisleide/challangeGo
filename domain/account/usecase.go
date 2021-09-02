package account

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

type UseCase interface {
	CreateAccount(account entities.AccountInput) (entities.Account, error)
	ListAll() []entities.Account
	Deposit(CPF string, amount float64) error
	Withdraw(CPF string, amount float64) bool
	GetBalance(accountID string) (entities.Account, error)
}

type Repository interface {
	FindOne(CPF string) (entities.Account, bool)
	UpdateBalance(account entities.Account) bool
	InsertAccount(entities.Account) error
	ListAllAccounts() []entities.Account
	FindByID(accountID string) (entities.Account, bool)
}
