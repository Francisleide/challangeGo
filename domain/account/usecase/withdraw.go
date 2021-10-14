package usecase

import "github.com/francisleide/ChallengeGo/domain/entities"

func (c AccountUc) Withdraw(CPF string, amount float64) (entities.TransactionOutput, error) {

	var withdrawOutput entities.TransactionOutput
	account, err := c.r.FindOne(CPF)
	withdrawOutput.ID = account.ID
	withdrawOutput.PreviousBalance = account.Balance
	if err != nil {
		return entities.TransactionOutput{}, err
	}
	if amount < 0 {
		return entities.TransactionOutput{}, ErrorInvalidValue
	}
	if account.Balance < amount {
		return entities.TransactionOutput{}, ErrorInsufficientBalance
	}
	account.Balance -= amount
	withdrawOutput.ActualBalance = account.Balance
	err = c.r.UpdateBalance(account.ID, account.Balance)
	if err != nil {
		return entities.TransactionOutput{}, err
	}
	return withdrawOutput, nil
}
