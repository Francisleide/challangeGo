package usecase

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

func (c AccountUc) ListAll() ([]entities.Account, error) {
	accounts, err := c.r.ListAllAccounts()
	if err != nil {
		c.log.WithError(err).Errorln(ErrorListAccounts)
		return []entities.Account{}, ErrorListAccounts
	}
	return accounts, nil
}
