package usecase

import (
	autenticoperations "github.com/francisleide/ChallangeGo/domain/autenticOperations"
	"github.com/francisleide/ChallangeGo/domain/entities"
)

type Autentic struct {
	r autenticoperations.Repository
}

func NewAutentic(repo autenticoperations.Repository) Autentic {
	return Autentic{
		r: repo,
	}
}

func (a Autentic) Deposite(cpf string, amount float64) {
	var account entities.Account
	if amount > 0 {
		account = a.r.FindOne(cpf)
		account.Balance += amount
		a.r.UpdateBalance(account)

	}

}

func (a Autentic) WithDraw(cpf string, ammount float64) bool {
	var account entities.Account
	account = a.r.FindOne(cpf)
	if account.Balance > ammount {
		account.Balance -= ammount
		a.r.UpdateBalance(account)
		return true
	}
	return false
}
