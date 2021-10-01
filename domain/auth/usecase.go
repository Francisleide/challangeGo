package auth

import "github.com/francisleide/ChallengeGo/domain/entities"

type UseCase interface {
	CreateToken(CPF string, secret string) (string, error)
	Login(CPF, secret string) error
}

type Repository interface {
	FindOne(CPF string) (entities.Account, error)
}
