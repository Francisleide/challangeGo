package auth

type UseCase interface {
	CreateToken(cpf string, secret string) (string, error)
}
