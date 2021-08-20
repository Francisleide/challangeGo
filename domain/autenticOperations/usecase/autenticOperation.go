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

func (a Autentic) Deposite(CPF string, amount float64) {
	var account entities.Account
	if amount > 0 {
		account = a.r.FindOne(CPF)
		account.Balance += amount
		a.r.UpdateBalance(account)

	}

}

func (a Autentic) WithDraw(CPF string, ammount float64) bool {
	var account entities.Account
	account = a.r.FindOne(CPF)
	if account.Balance > ammount {
		account.Balance -= ammount
		a.r.UpdateBalance(account)
		return true
	}
	return false
}
