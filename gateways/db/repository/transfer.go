package repository

import (
	"fmt"

	"github.com/francisleide/ChallangeGo/domain/entities"
	"github.com/francisleide/ChallangeGo/domain/transfer"
)

var (
	id_transfer            string
	account_origin_id      string
	account_destination_id string
	amount                 float64
	transfer_created_at    string
)
var tranfers []entities.Transfer

func (r Repository) List_all_transfers() []entities.Transfer {

	rows, err := r.Db.Query("SELECT id, account_origin_id, account_destination_id, amount,transfer_created_at, from transfer;")
	defer rows.Close()
	checkError(err)
	for rows.Next() {
		err = rows.Scan(&id_transfer, &account_origin_id, &account_destination_id, &amount, &transfer_created_at)
		tranfer := entities.Transfer{id_transfer, account_origin_id, account_destination_id, amount, transfer_created_at}
		tranfers = append(tranfers, tranfer)

	}
	checkError(err)
	err = rows.Err()
	return tranfers
}

func (r Repository) InsertTransfer(account_origem, account_destino entities.Account, ammount float64) (*entities.Transfer, error) {
	//Atualizar o balance da conta de origem
	//isso é regra de negócio, deveria estar no usecase
	r.UpdateBalance(account_origem)
	//Atualizar o balance da conta de destino
	//isso é regra de negócio, deveria estar no usecase
	r.UpdateBalance(account_destino)
	//inserir a transferência

	t := transfer.NewTransferInput(account_origem.Id, account_destino.Id, ammount)
	fmt.Printf(t.Id)
	_, err := r.Db.Query("insert into  transfer (id, account_origin_id, account_destination_id,amount,created_at) values (?,?,?,?,?)",
		t.Id, t.Account_origin_id, t.Account_destination_id, t.Amount, t.Created_at)

	if err != nil {
		checkError(err)
		return nil, err
	}
	return &t, nil

}
