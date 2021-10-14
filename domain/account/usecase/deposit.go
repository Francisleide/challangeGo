package usecase

import "github.com/francisleide/ChallengeGo/domain/entities"

func (c AccountUc) Deposit(CPF string, amount float64) (entities.TransactionOutput, error) {
	var depositOutput entities.TransactionOutput
	account, err := c.r.FindOne(CPF)
	depositOutput.PreviousBalance = account.Balance
	depositOutput.ID = account.ID
	if err != nil {
		c.log.WithError(err).Errorln(ErrorRetrieveAccount)
		return entities.TransactionOutput{}, ErrorRetrieveAccount
	}
	if amount <= 0 {
		c.log.WithError(err).Errorln(ErrorInvalidValue)
		return entities.TransactionOutput{}, ErrorInvalidValue
	}
	account.Balance += amount
	err = c.r.UpdateBalance(account.ID, account.Balance)
	if err != nil {
		c.log.WithError(err).Errorln(ErrorUpdateBalance)
		return entities.TransactionOutput{}, ErrorUpdateBalance
	}
	depositOutput.ActualBalance = account.Balance
	return depositOutput, nil

}
