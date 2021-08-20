package repository

import (
	"fmt"

	"github.com/francisleide/ChallangeGo/domain/entities"
	"github.com/francisleide/ChallangeGo/domain/transfer"
)

var (
	IDTransfer           string
	accountOriginID      string
	accountDestinationID string
	amount               float64
	transferCreatedAt    string
)
var tranfers []entities.Transfer

func (r Repository) ListAllTransfers() []entities.Transfer {

	rows, err := r.Db.Query("SELECT id, account_origin_id, account_destination_id, amount,transfer_created_at, from transfer;")
	defer rows.Close()
	checkError(err)
	for rows.Next() {
		err = rows.Scan(&IDTransfer, &accountOriginID, &accountDestinationID, &amount, &transferCreatedAt)
		tranfer := entities.Transfer{IDTransfer, accountOriginID, accountDestinationID, amount, transferCreatedAt}
		tranfers = append(tranfers, tranfer)

	}
	checkError(err)
	err = rows.Err()
	return tranfers
}

func (r Repository) InsertTransfer(accountOrigin, accountDestine entities.Account, ammount float64) (*entities.Transfer, error) {
	//Atualizar o balance da conta de origem
	//isso é regra de negócio, deveria estar no usecase
	r.UpdateBalance(accountOrigin)
	//Atualizar o balance da conta de destino
	//isso é regra de negócio, deveria estar no usecase
	r.UpdateBalance(accountDestine)
	//inserir a transferência

	t := transfer.NewTransferInput(accountOrigin.ID, accountDestine.ID, ammount)
	fmt.Printf(t.ID)
	_, err := r.Db.Query("insert into  transfer (id, account_origin_id, account_destination_id,amount,created_at) values (?,?,?,?,?)",
		t.ID, t.AccountOriginID, t.AccountDestinationID, t.Amount, t.CreatedAt)

	if err != nil {
		checkError(err)
		return nil, err
	}
	return &t, nil

}
