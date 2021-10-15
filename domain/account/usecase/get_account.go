package usecase

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

func (c AccountUc) GetAccountByID(ID string) (entities.Account, error) {
	account, err := c.r.FindByID(ID)
	if err != nil {
		return entities.Account{}, ErrorRetrieveAccount
	}

	return account, nil
}
func (c AccountUc) GetAccountByCPF(CPF string) (entities.Account, error) {
	account, err := c.r.FindOne(CPF)
	if err != nil {
		return entities.Account{}, ErrorRetrieveAccount
	}
	return account, nil

}
