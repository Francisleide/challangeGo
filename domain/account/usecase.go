package account

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

type UseCase interface {
	CreateAccount(account entities.AccountInput) (entities.Account, error)
	ListAll() []entities.Account
	Deposit(CPF string, amount float64) error
	Withdraw(CPF string, amount float64) error
	GetBalance(accountID string) (float64, error)
	GetAccountByID(ID string) (entities.Account, error)
	GetAccountByCPF(CPF string) (entities.Account, error)
}

type Repository interface {
	FindOne(CPF string) (entities.Account, error)
	UpdateBalance(ID string, balance float64) error
	InsertAccount(entities.Account) error
	ListAllAccounts() []entities.Account
	FindByID(accountID string) (entities.Account, error)
}
