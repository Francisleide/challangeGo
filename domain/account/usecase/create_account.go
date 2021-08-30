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
	newAccount, err := entities.NewAccount(accountInput.Name, accountInput.CPF, accountInput.Secret)
	if err != nil {
		//TODO: add a sentinel
		return entities.Account{}, err
	}

	_, ok := c.r.FindOne(newAccount.CPF)
	if ok {
		//TODO: add a sentinel
		return entities.Account{}, errors.New("the account already exists")
	}
	errInsert := c.r.InsertAccount(newAccount)
	if errInsert != nil {
		return entities.Account{}, errInsert
	}

	return newAccount, nil
}
