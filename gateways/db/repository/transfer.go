package repository

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

func (r Repository) InsertTransfer(transfer entities.Transfer) error {

	_, err := r.Db.Query("insert into transfer (id, account_origin_id, account_destination_id,amount,created_at) values (?,?,?,?,?)",
		transfer.ID, transfer.AccountOriginID, transfer.AccountDestinationID, transfer.Amount, transfer.CreatedAt)

	if err != nil {
		r.log.WithError(err).Errorln("failed to insert transfer")
		return err
	}
	r.log.Info("transfer registered successfully")
	return nil

}

func (r Repository) ListUserTransfers(accountID string) ([]entities.Transfer, error) {
	var transfers []entities.Transfer
	rows, err := r.Db.Query("select * from transfer where account_origin_id=?", accountID)
	if err != nil {
		r.log.WithError(err).Errorf("unable to find transfers from id %s", accountID)
		return []entities.Transfer{}, err
	}
	for rows.Next() {
		var transfer entities.Transfer
		err = rows.Scan(&transfer.ID, &transfer.AccountOriginID, &transfer.AccountDestinationID, &transfer.Amount, &transfer.CreatedAt)
		if err != nil {
			r.log.WithError(err).Errorln("failed to read record")
			return []entities.Transfer{}, err
		}
		transfers = append(transfers, transfer)
	}
	r.log.Info("")
	return transfers, nil

}
