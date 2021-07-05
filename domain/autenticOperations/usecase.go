package autenticoperations

type UseCase interface {
	Deposite(cpf string, ammount float64)
	WithDraw(cpf string, ammount float64)(bool)
}
