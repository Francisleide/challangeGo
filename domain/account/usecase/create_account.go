package usecase

import (
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

func (c AccountUc) CreateAccount(account entities.AccountInput) (*entities.Account, error) {
	account2, err := c.r.InsertAccount(account)
	if err != nil {
		return nil, ErrorAccountAlreadyExists
	}
	return account2, nil
}
