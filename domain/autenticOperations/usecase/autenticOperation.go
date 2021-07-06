package usecase

import (
	"github.com/francisleide/ChallangeGo/domain/entities"
	"github.com/francisleide/ChallangeGo/gateways/repository"
)

type Autentic struct {
	r repository.Repository
}

func NewAutentic(repo repository.Repository) Autentic {
	return Autentic{
		r: repo,
	}
}

func (a Autentic) Deposite(cpf string, ammount float64) {
	var account entities.Account
	account = a.r.FindOne(cpf)
	account.Balance += ammount
	a.r.UpdateBalance(account)

}


//withdraw aqui!
func (a Autentic)WithDraw(cpf string, ammount float64) bool {
	var account entities.Account
	account = a.r.FindOne(cpf)
	if account.Balance > ammount {
		account.Balance -= ammount
		a.r.UpdateBalance(account)
		return true
	}
	return false
}