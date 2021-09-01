package repository

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

var (
	transferID           string
	accountOriginID      string
	accountDestinationID string
	amount               float64
	transferCreatedAt    string
)
var transfers []entities.Transfer

func (r Repository) ListAllTransfers() []entities.Transfer {

	rows, err := r.Db.Query("select id, account_origin_id, account_destination_id, amount,transfer_created_at, from transfer;")
	defer rows.Close()
	checkError(err)
	for rows.Next() {
		err = rows.Scan(&transferID, &accountOriginID, &accountDestinationID, &amount, &transferCreatedAt)
		tranfer := entities.Transfer{transferID, accountOriginID, accountDestinationID, amount, transferCreatedAt}
		transfers = append(transfers, tranfer)

	}
	checkError(err)
	err = rows.Err()
	return transfers
}

func (r Repository) InsertTransfer(transfer entities.Transfer) (entities.Transfer, error) {
	//r.UpdateBalance(accountOrigin) chamar isso no UC
	//r.UpdateBalance(accountDestine) chamar isso no UC

	//t := transfer.NewTransferInput(accountOrigin.ID, accountDestine.ID, amount)
	_, err := r.Db.Query("insert into  transfer (id, account_origin_id, account_destination_id,amount,created_at) values (?,?,?,?,?)",
		transfer.ID, transfer.OriginAccountID, transfer.DestinationAccountID, transfer.Amount, transfer.CreatedAt)

	if err != nil {
		checkError(err)
		return entities.Transfer{}, err
	}
	return transfer, nil

}

