package usecase

import "errors"

func (c AccountUc) Deposit(CPF string, amount float64) error {

	account, err := c.r.FindOne(CPF)
	if err != nil {
		//TODO add a new sentinel
		return errors.New("the account does not exist")
	}
	account.Balance += amount
	c.r.UpdateBalance(account)
	return nil

}
