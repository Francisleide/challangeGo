package usecase

import "errors"

func (c AccountUc) Withdraw(CPF string, amount float64) error {

	account, err := c.r.FindOne(CPF)
	if err != nil {
		return err
	}
	if account.Balance < amount {
		//TODO: add a sentinel
		return errors.New("insufficient balance")
	}
	account.Balance -= amount
	err = c.r.UpdateBalance(account.ID, account.Balance)
	if err != nil {
		return err
	}
	return nil
}
