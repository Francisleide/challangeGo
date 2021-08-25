package usecase

import "github.com/francisleide/ChallengeGo/domain/entities"

func (c AccountUc) Withdraw(CPF string, amount float64) bool {
	var account entities.Account
	account = c.r.FindOne(CPF)
	if account.Balance > amount {
		account.Balance -= amount
		c.r.UpdateBalance(account)
		return true
	}
	return false
}
