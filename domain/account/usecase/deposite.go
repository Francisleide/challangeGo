package usecase

import "github.com/francisleide/ChallangeGo/domain/entities"



func (c AccountUc) Deposite(CPF string, amount float64) {
	var account entities.Account
	if amount > 0 {
		account = c.r.FindOne(CPF)
		account.Balance += amount
		c.r.UpdateBalance(account)

	}

}
