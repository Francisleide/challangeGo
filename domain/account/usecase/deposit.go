package usecase

import (
	"errors"
)

func (c AccountUc) Deposit(CPF string, amount float64) error {

	account, err := c.r.FindOne(CPF)
	if err != nil {
		c.log.WithError(err).Errorln("the account does not exist")
		//TODO add a new sentinel
		return errors.New("the account does not exist")
	}
	if amount <= 0 {
		c.log.WithError(err).Errorln("invalid value")
		//TODO add a new sentinel
		return errors.New("invalid value")
	}
	account.Balance += amount
	err = c.r.UpdateBalance(account.ID, account.Balance)
	if err != nil {
		c.log.WithError(err).Errorf("failed to update balance")
		return err
	}
	return nil

}
