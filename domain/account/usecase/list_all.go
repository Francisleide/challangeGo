package usecase

import "github.com/francisleide/ChallengeGo/domain/entities"

func (c AccountUc) ListAll() []entities.Account {
	accounts, err := c.r.ListAllAccounts()
	if err != nil {
		c.log.WithError(err).Errorln("failed to list accounts")
	}
	return accounts
}
