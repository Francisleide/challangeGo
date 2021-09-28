package usecase

import (
	"errors"

	"github.com/francisleide/ChallengeGo/domain/entities"
)

func (c AccountUc) GetAccountByID(ID string) (entities.Account, error) {
	account, err := c.r.FindByID(ID)
	if err != nil {
		c.log.WithError(err).Errorln("failed to retrieve the account from repository")
		//TODO add a new sentinel
		return entities.Account{}, errors.New("failed to retrieve the account from repository")
	}

	return account, nil
}
func (c AccountUc) GetAccountByCPF(CPF string) (entities.Account, error) {
	account, err := c.r.FindOne(CPF)
	if err != nil {
		c.log.WithError(err).Errorln("failed to retrieve the account from repository")
		return entities.Account{}, err
	}
	return account, nil

}
