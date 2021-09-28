package usecase

import (
	"errors"

	"github.com/francisleide/ChallengeGo/domain/entities"
)

func (c AccountUc) ListAll() ([]entities.Account, error) {
	accounts, err := c.r.ListAllAccounts()
	if err != nil {
		c.log.WithError(err).Errorln("failed to list accounts")
		//TODO: add a new sentinel
		return []entities.Account{}, errors.New("failed to list accounts")
	}
	return accounts, nil
}
