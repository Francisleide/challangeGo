package usecase

import (
	"errors"
)

func (c AccountUc) GetBalance(accountID string) (float64, error) {
	account, err := c.r.FindByID(accountID)
	if err != nil {
		//TODO: add a sentinel
		return 0, errors.New("failed to retrieve the account from repository")
	}
	return account.Balance, nil
}
