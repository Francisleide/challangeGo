package usecase

func (c AccountUc) Withdraw(CPF string, amount float64) bool {

	account, err := c.r.FindOne(CPF)
	if err != nil {
		return false
	}
	if account.Balance > amount {
		account.Balance -= amount
		c.r.UpdateBalance(account)
		return true
	}
	return false
}
