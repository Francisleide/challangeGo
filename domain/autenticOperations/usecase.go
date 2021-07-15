package autenticoperations

import "github.com/francisleide/ChallangeGo/domain/entities"

type UseCase interface {
	Deposite(cpf string, ammount float64)
	WithDraw(cpf string, ammount float64)(bool)
}


type Repository interface{
	InsertAccount(accountInput entities.AccountInput) (*entities.Account, error)
	UpdateBalance(account entities.Account)
	FindOne(cpf string) entities.Account
}