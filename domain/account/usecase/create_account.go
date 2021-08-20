package usecase

import (
	"fmt"

	"github.com/francisleide/ChallangeGo/domain/account"
	"github.com/francisleide/ChallangeGo/domain/entities"
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
	fmt.Println("CPF no Usecase: ", account.CPF)
	account2, err := c.r.InsertAccount(account)
	if err != nil {
		return nil, ErrorAccountAlreadyExists
	}
	return account2, nil
}

func (c AccountUc) ListAllAccounts() []entities.Account {
	accounts := c.r.ListAllAccounts()
	return accounts

}
