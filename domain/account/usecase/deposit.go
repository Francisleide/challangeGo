package usecase

import (
	ac "github.com/francisleide/ChallengeGo/domain/account"
)

func (c AccountUc) Deposit(CPF string, amount float64) (ac.TransactionOutput, error) {
	var depositOutput ac.TransactionOutput
	account, err := c.r.FindOne(CPF)
	if err != nil {
		c.log.WithError(err).Errorln(ErrorRetrieveAccount)
		return ac.TransactionOutput{}, ErrorRetrieveAccount
	}
	depositOutput.PreviousBalance = account.Balance
	depositOutput.ID = account.ID
	if amount <= 0 {
		return ac.TransactionOutput{}, ErrorInvalidValue
	}
	account.Balance += amount
	err = c.r.UpdateBalance(account.ID, account.Balance)
	if err != nil {
		return ac.TransactionOutput{}, ErrorUpdateBalance
	}
	depositOutput.ActualBalance = account.Balance
	return depositOutput, nil

}
