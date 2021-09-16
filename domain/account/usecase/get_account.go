package usecase

import (
	"errors"

	"github.com/francisleide/ChallengeGo/domain/entities"
)

func (c AccountUc) GetAccountByID(ID string) (entities.Account, error) {
	account, err := c.r.FindByID(ID)
	if err != nil {
		c.log.WithError(err).Errorln("unable to find account")
		return entities.Account{}, err
	}
	if account == (entities.Account{}) {
		return entities.Account{}, errors.New("account not found")
	}
	return account, nil
}
func (c AccountUc) GetAccountByCPF(CPF string) (entities.Account, error) {
	account, err := c.r.FindOne(CPF)
	if err != nil {
		c.log.WithError(err).Errorln("account not found")
		return entities.Account{}, err
	}
	return account, nil

}
