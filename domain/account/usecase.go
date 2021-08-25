package account

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

type UseCase interface {
	CreateAccount(account entities.AccountInput) (*entities.Account, error)
	ListAll() []entities.Account
	Deposit(CPF string, amount float64)
	Withdraw(CPF string, amount float64) bool
}

type Repository interface {
	FindOne(CPF string) entities.Account
	UpdateBalance(account entities.Account)
	InsertAccount(accountInput entities.AccountInput) (*entities.Account, error)
	ListAllAccounts() []entities.Account
}
