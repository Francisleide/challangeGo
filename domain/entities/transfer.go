package entities

import "time"

type Transfer struct {
	Id                     string  `json: "transfer_id"`
	Account_origin_id      string  `json: "account_origin_id"`
	Account_destination_id string  `json: "account_destination_id"`
	Ammount                float64 `json: "ammount"`
	Created_at             string  `json: "account_destination_id"`
}

type TransferInput struct {
	Cpf_destino string  `json: "cpf_destino"`
	Ammount     float64 `json: "ammount"`
}

func NewTransferInput(account_origin_id, account_destination_id string, ammount float64) Transfer {
	return Transfer{
		Id:                     GenerateId(),
		Account_origin_id:      account_origin_id,
		Account_destination_id: account_destination_id,
		Ammount:                ammount,
		Created_at:             time.Now().Format(time.RFC822),
	}
}
