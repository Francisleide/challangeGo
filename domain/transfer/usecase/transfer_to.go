package usecase

import (
	"errors"
	"fmt"

	"github.com/francisleide/ChallangeGo/domain/entities"
	"github.com/francisleide/ChallangeGo/domain/transfer"
)

type TransferUc struct {
	r transfer.Repository
}

func NewTransfer(repo transfer.Repository) TransferUc {
	return TransferUc{
		r: repo,
	}
}

//
func (t TransferUc) CreateTransfer(origem, destino string, ammount float64) (*entities.Transfer, error) {
	var accountOrigem entities.Account
	var accountDestine entities.Account
	accountOrigem = t.r.FindOne(origem)
	fmt.Println("CPF da conta de origem: ", accountOrigem.CPF)
	fmt.Println("Saldo anterior na conta de origem: ", accountOrigem.Balance)
	accountDestine = t.r.FindOne(destino)
	fmt.Println("Saldo anterior na conta de destino: ", accountDestine.Balance)
	if accountOrigem.Balance >= ammount {
		var tr *entities.Transfer
		accountOrigem.Balance -= ammount
		fmt.Println("Valor no usecase origem: ", accountOrigem.Balance)
		accountDestine.Balance += ammount
		tr, err := t.r.InsertTransfer(accountOrigem, accountDestine, ammount)
		//a transfer vai ser montada lá no repository e enviada pra o banco
		if err != nil {
			return nil, err
		}
		return tr, nil
	} else {
		return nil, errors.New("Saldo insuficiente")
	}
}
