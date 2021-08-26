package auth

type UseCase interface {
	CreateToken(CPF string, secret string) (string, error)
	
}
