package usecase

import (
	"errors"
	"fmt"

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

	_, errFind := c.r.FindOne(newAccount.CPF)
	if errFind != nil {
		//TODO: add a sentinel
		fmt.Println(errFind)
		return entities.Account{}, errors.New("the account already exists")
	}
	err = c.r.InsertAccount(newAccount)
	if err != nil {
		return entities.Account{}, err
	}

	return newAccount, nil
}
