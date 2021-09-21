package repository

import (
	"github.com/francisleide/ChallengeGo/domain/entities"
)

func (r Repository) ListAllAccounts() ([]entities.Account, error) {
	var accounts []entities.Account

	rows, err := r.Db.Query("SELECT id, name, cpf, secret,balance, created_at from account;")
	defer rows.Close()
	if err != nil {
		r.log.WithError(err).Errorln("failed to perform query on account table")
		return []entities.Account{}, err
	}
	for rows.Next() {
		var account entities.Account
		err = rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		accounts = append(accounts, account)
		if err != nil {
			r.log.WithError(err).Errorln("failed to read record")
			return []entities.Account{}, err
		}

	}
	r.log.Info("all records have been read")
	return accounts, nil
}

func (r Repository) FindOne(CPF string) (entities.Account, error) {

	var accounts []entities.Account

	sql := "SELECT id, name, cpf, secret,balance, created_at from account where cpf=?"
	rows, err := r.Db.Query(sql, CPF)
	if err != nil {
		r.log.WithError(err).Errorln("Failed to perform query on account table")
		return entities.Account{}, err
	}
	for rows.Next() {
		var account entities.Account
		err = rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		if err != nil {
			r.log.WithError(err).Errorln("failed to read record")
			return entities.Account{}, err
		}
		accounts = append(accounts, account)

	}
	if len(accounts) == 0 {
		r.log.Errorln("no records found")
		return entities.Account{}, err
	}
	r.log.Info("the account was found")
	return accounts[0], nil

}

func (r Repository) UpdateBalance(ID string, balance float64) error {
	rows, err := r.Db.Exec("UPDATE account SET balance = ? WHERE id = ?", balance, ID)
	if err != nil {
		r.log.WithError(err).Errorln("failed to run update")
		return err
	}
	rowCount, err := rows.RowsAffected()
	if rowCount < 1 {
		r.log.WithError(err).Errorln("failed to update registry")
		return err
	}
	r.log.Infof("id %s account balance has been updated", ID)
	return nil

}

func (r Repository) InsertAccount(account entities.Account) error {

	_, err := r.Db.Query("insert into  account (id, name, cpf, secret,balance, created_at) values (?,?,?,?,?,? )",
		account.ID, account.Name, account.CPF, account.Secret, account.Balance, account.CreatedAt)

	if err != nil {
		r.log.WithError(err).Errorln("failed to insert a new account")
		return err
	}
	r.log.Infof("new account id %s created successfully", account.ID)
	return nil

}

func (r Repository) FindByID(accountID string) (entities.Account, error) {
	var accounts []entities.Account
	rows, err := r.Db.Query("select * from account where id=?", accountID)
	if err != nil {
		r.log.WithError(err).Errorf("failed to find account id %s", accountID)
		return entities.Account{}, err
	}
	for rows.Next() {
		var account entities.Account
		err = rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		if err != nil {
			r.log.WithError(err).Errorf("failed to find account id %s", accountID)
		}
		accounts = append(accounts, account)
	}
	if len(accounts) == 0 {
		return entities.Account{}, err
	}
	r.log.Infof("the account was found, id: %s", accountID)
	return accounts[0], nil

}
