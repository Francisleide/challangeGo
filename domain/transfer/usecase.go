package transfer

import "github.com/francisleide/ChallangeGo/domain/entities"

type UseCase interface {
	Create_transfer(origem, destino string, ammount float64)  (*entities.Transfer, error)
}