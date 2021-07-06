package usecase

import (
	"github.com/francisleide/ChallangeGo/domain/entities"
)

func (a AccountUc) List_all_accounts() []entities.Account {
	accounts := a.r.List_all_accounts()
	return accounts

}
