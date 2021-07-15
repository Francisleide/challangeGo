package transfer

import (
	"time"

	"github.com/francisleide/ChallangeGo/domain/entities"
)

type UseCase interface {
	Create_transfer(origem, destino string, ammount float64)  (*entities.Transfer, error)
}

type TransferInput struct {
	Cpf_destino string  `json: "cpf_destino"`
	Ammount     float64 `json: "ammount"`
}

type Repository interface{
	InsertTransfer(account_origem, account_destino entities.Account, ammount float64)(*entities.Transfer, error)
	FindOne(cpf string) entities.Account
}

func NewTransferInput(account_origin_id, account_destination_id string, ammount float64) entities.Transfer {
	return entities.Transfer{
		Id:                     entities.GenerateId(),
		Account_origin_id:      account_origin_id,
		Account_destination_id: account_destination_id,
		Amount:                ammount,
		Created_at:             time.Now().Format(time.RFC822),
	}
}