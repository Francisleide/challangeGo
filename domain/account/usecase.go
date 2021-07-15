package account

import (
	"github.com/francisleide/ChallangeGo/domain/entities"
)

type UseCase interface {
	Create_account(account entities.AccountInput) (*entities.Account, error)
	List_all_accounts() []entities.Account
}

type Repository interface {
	FindOne(cpf string) entities.Account
	UpdateBalance(account entities.Account)
	InsertAccount(accountInput entities.AccountInput) (*entities.Account, error)
	List_all_accounts() []entities.Account
}
