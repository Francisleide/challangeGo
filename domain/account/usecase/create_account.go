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

func (c AccountUc) Create_account(account entities.AccountInput) (*entities.Account, error) {
	fmt.Println("CPF no Usecase: ", account.Cpf)
	account2, err := c.r.InsertAccount(account)
	if err != nil {
		return nil, ErrorAccountAlreadyExists
	}
	return account2, nil
}

func (c AccountUc) List_all_accounts() []entities.Account {
	accounts := c.r.List_all_accounts()
	return accounts

}
