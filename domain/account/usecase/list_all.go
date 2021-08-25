package usecase

import "github.com/francisleide/ChallengeGo/domain/entities"

func (c AccountUc) ListAll() []entities.Account {
	accounts := c.r.ListAllAccounts()
	return accounts
}
