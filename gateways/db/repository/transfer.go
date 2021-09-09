package repository

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

func (r Repository) InsertTransfer(transfer entities.Transfer) (entities.Transfer, error) {

	_, err := r.Db.Query("insert into  transfer (id, account_origin_id, account_destination_id,amount,created_at) values (?,?,?,?,?)",
		transfer.ID, transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount, transfer.CreatedAt)

	if err != nil {
		checkError(err)
		return entities.Transfer{}, err
	}
	return transfer, nil

}

func (r Repository) ListUserTransfers(accountID string) ([]entities.Transfer, error) {
	var transfers []entities.Transfer
	rows, err := r.Db.Query("select * from transfer where account_origin_id=?", accountID)
	if err != nil {
		return []entities.Transfer{}, err
	}
	for rows.Next() {
		var transfer entities.Transfer
		err = rows.Scan(&transfer.ID, &transfer.AccountOriginID, &transfer.AccountDestinationID, &transfer.Amount, &transfer.CreatedAt)
		transfers = append(transfers, transfer)
	}
	if err != nil {
		return []entities.Transfer{}, err
	}
	return transfers, nil

}
