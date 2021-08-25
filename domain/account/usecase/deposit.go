package usecase

import "github.com/francisleide/ChallengeGo/domain/entities"



func (c AccountUc) Deposit(CPF string, amount float64) {
	var account entities.Account
	if amount > 0 {
		account = c.r.FindOne(CPF)
		account.Balance += amount
		c.r.UpdateBalance(account)

	}

}
