package usecase

import "github.com/francisleide/ChallangeGo/domain/entities"

func (c AccountUc) WithDraw(CPF string, ammount float64) bool {
	var account entities.Account
	account = c.r.FindOne(CPF)
	if account.Balance > ammount {
		account.Balance -= ammount
		c.r.UpdateBalance(account)
		return true
	}
	return false
}
