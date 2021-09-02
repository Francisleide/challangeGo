package usecase

import (
	"errors"

	"github.com/francisleide/ChallengeGo/domain/entities"
)

func (c AccountUc) GetBalance(accountID string) (entities.Account, error){
	account, ok := c.r.FindByID(accountID)
	if ! ok{
		//TODO: add a sentinel
		return entities.Account{}, errors.New("account not found")
	}
	return account, nil
}