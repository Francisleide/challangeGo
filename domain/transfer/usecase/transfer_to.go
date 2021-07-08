package usecase

import (
	"errors"
	"fmt"

	"github.com/francisleide/ChallangeGo/domain/entities"
	"github.com/francisleide/ChallangeGo/gateways/db/repository"
)

type TransferUc struct {
	r repository.Repository
}

func NewTransfer(repo repository.Repository) TransferUc {
	return TransferUc{
		r: repo,
	}
}

//
func (t TransferUc) Create_transfer(origem, destino string, ammount float64) (*entities.Transfer, error) {
	var account_origem entities.Account
	var account_destino entities.Account
	account_origem = t.r.FindOne(origem)
	fmt.Printf("CPF da conta de origem: ", account_origem.Cpf)
	fmt.Println("Saldo anterior na conta de origem: ", account_origem.Balance)
	account_destino = t.r.FindOne(destino)
	fmt.Println("Saldo anterior na conta de destino: ", account_destino.Balance)
	if account_origem.Balance >= ammount {
		var tr *entities.Transfer
		account_origem.Balance -= ammount
		fmt.Println("Valor no usecase origem: ", account_origem.Balance)
		account_destino.Balance += ammount
		tr, err := t.r.InsertTransfer(account_origem, account_destino, ammount)
		//a transfer vai ser montada lá no repository e enviada pra o banco
		if err != nil {
			return nil, err
		}
		return tr, nil
	} else {
		return nil, errors.New("Saldo insuficiente")
	}
}
