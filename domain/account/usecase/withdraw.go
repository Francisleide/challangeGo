package usecase

import (
	ac "github.com/francisleide/ChallengeGo/domain/account"
)

func (c AccountUc) Withdraw(CPF string, amount float64) (ac.TransactionOutput, error) {

	var withdrawOutput ac.TransactionOutput
	account, err := c.r.FindOne(CPF)
	withdrawOutput.ID = account.ID
	withdrawOutput.PreviousBalance = account.Balance
	if err != nil {
		return ac.TransactionOutput{}, err
	}
	if amount < 0 {
		return ac.TransactionOutput{}, ErrorInvalidValue
	}
	if account.Balance < amount {
		return ac.TransactionOutput{}, ErrorInsufficientBalance
	}
	account.Balance -= amount
	withdrawOutput.ActualBalance = account.Balance
	err = c.r.UpdateBalance(account.ID, account.Balance)
	if err != nil {
		return ac.TransactionOutput{}, err
	}
	return withdrawOutput, nil
}
