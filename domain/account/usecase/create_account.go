package usecase

import (
	"errors"

	"github.com/francisleide/ChallengeGo/domain/account"
	"github.com/francisleide/ChallengeGo/domain/entities"
	"github.com/sirupsen/logrus"
)

type AccountUc struct {
	r   account.Repository
	log *logrus.Entry
}

func NewAccountUc(repo account.Repository, log *logrus.Entry) AccountUc {
	return AccountUc{
		r:   repo,
		log: log,
	}
}

func (c AccountUc) CreateAccount(accountInput entities.AccountInput) (entities.Account, error) {
	account, _ := c.r.FindOne(accountInput.CPF)
	if account != (entities.Account{}) {
		//TODO: add a sentinel
		return entities.Account{}, errors.New("the account already exists")
	}

	newAccount, err := entities.NewAccount(accountInput.Name, accountInput.CPF, accountInput.Secret)
	if err != nil {
		c.log.WithError(err).Errorln("unable to create account")
		//TODO: add a sentinel
		return entities.Account{}, err
	}

	err = c.r.InsertAccount(newAccount)
	if err != nil {
		return entities.Account{}, err
	}

	return newAccount, nil
}
