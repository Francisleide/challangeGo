package autenticoperations

import "github.com/francisleide/ChallangeGo/domain/entities"

type UseCase interface {
	Deposite(CPF string, ammount float64)
	WithDraw(CPF string, ammount float64)(bool)
}


type Repository interface{
	InsertAccount(accountInput entities.AccountInput) (*entities.Account, error)
	UpdateBalance(account entities.Account)
	FindOne(CPF string) entities.Account
}