package usecase

import (
	"fmt"

	"github.com/francisleide/ChallangeGo/domain/entities"
	"github.com/francisleide/ChallangeGo/gateways/repository"
)

type AccountUc struct {
	r repository.Repository
}

func NewAccountUc(repo repository.Repository) AccountUc {
	return AccountUc{
		r: repo,
	}
}

//
func (c AccountUc) Create_account(account entities.AccountInput) (*entities.Account, error) {
	fmt.Println("CPF no Usecase: ", account.Cpf)
	account2, err := c.r.InsertAccount(account)
	if err != nil {
		return nil, err
	}
	return account2, nil
}
