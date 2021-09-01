package repository

import (
	"fmt"
	"log"

	"github.com/francisleide/ChallengeGo/domain/entities"
)

func (r Repository) ListAllAccounts() []entities.Account {
	var accounts []entities.Account

	rows, err := r.Db.Query("SELECT id, name, cpf, secret,balance, created_at from account;")
	defer rows.Close()
	checkError(err)
	fmt.Println("Reading data:")
	for rows.Next() {
		var account entities.Account
		err = rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		accounts = append(accounts, account)

	}
	checkError(err)
	err = rows.Err()
	fmt.Println("Done.")
	return accounts
}

func (r Repository) FindOne(CPF string) (entities.Account, bool) {

	var accounts []entities.Account

	sql := "SELECT id, name, cpf, secret,balance, created_at from account where cpf=?"
	rows, err := r.Db.Query(sql, CPF)
	checkError(err)
	for rows.Next() {
		var account entities.Account
		err = rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		accounts = append(accounts, account)
		checkError(err)
	}
	if len(accounts) == 0 {
		return entities.Account{}, false
	}
	return accounts[0], true

}

func (r Repository) UpdateBalance(account entities.Account) bool {

	rows, err := r.Db.Exec("UPDATE account SET balance = ? WHERE id = ?", account.Balance, account.ID)
	checkError(err)
	rowCount, err := rows.RowsAffected()
	if err != nil || rowCount < 1 {
		return false
	}
	return true

}

func (r Repository) InsertAccount(account entities.Account) error {

	_, err := r.Db.Query("insert into  account (id, name, cpf, secret,balance, created_at) values (?,?,?,?,?,? )",
		account.ID, account.Name, account.CPF, account.Secret, account.Balance, account.CreatedAt)

	if err != nil {
		checkError(err)
		return err
	}
	return nil

}

func (r Repository) FindByID(accountID string) (entities.Account, bool) {
	var accounts []entities.Account
	rows, err := r.Db.Query("select * from account where id=?", accountID)
	if err != nil {
		checkError(err)
		return entities.Account{}, false
	}
	for rows.Next() {
		var account entities.Account
		err = rows.Scan(&account.ID, &account.Name, &account.CPF, &account.Secret, &account.Balance, &account.CreatedAt)
		accounts = append(accounts, account)
		checkError(err)
	}
	if len(accounts) == 0 {
		return entities.Account{}, false
	}
	return accounts[0], true

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
