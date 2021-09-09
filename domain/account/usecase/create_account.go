package usecase

import (
	"errors"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/entities"
)

type AccountUc struct {
	r account.Repository
}

func NewAccountUc(repo account.Repository) AccountUc {
	return AccountUc{
		r: repo,
	}
}

func (c AccountUc) CreateAccount(accountInput entities.AccountInput) (entities.Account, error) {
	_, errFind := c.r.FindOne(accountInput.CPF)
	if errFind != nil {
		//TODO: add a sentinel
		return entities.Account{}, errors.New("the account already exists")
	}
	newAccount, err := entities.NewAccount(accountInput.Name, accountInput.CPF, accountInput.Secret)
	if err != nil {
		//TODO: add a sentinel
		return entities.Account{}, err
	}

	err = c.r.InsertAccount(newAccount)
	if err != nil {
		return entities.Account{}, err
	}

	return newAccount, nil
}
