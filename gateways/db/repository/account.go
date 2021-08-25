package repository

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/francisleide/ChallengeGo/domain/entities"
)

var (
	ID        string
	name      string
	CPF       string
	secret    string
	balance   float64
	createdAt string
)
var accounts []entities.Account

func (r Repository) ListAllAccounts() []entities.Account {

	rows, err := r.Db.Query("SELECT id, name, cpf, secret,balance, created_at from account;")
	defer rows.Close()
	checkError(err)
	fmt.Println("Reading data:")
	for rows.Next() {
		err = rows.Scan(&ID, &name, &CPF, &secret, &balance, &createdAt)
		account := entities.Account{ID, name, CPF, secret, balance, createdAt}
		accounts = append(accounts, account)

	}
	checkError(err)
	err = rows.Err()
	fmt.Println("Done.")
	return accounts
}

func (r Repository) FindOne(CPF string) entities.Account {

	var accounts []entities.Account
	var sql string
	sql = "SELECT id, name, cpf, secret,balance, created_at from account where cpf=?"
	rows, err := r.Db.Query(sql, CPF)
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&ID, &name, &CPF, &secret, &balance, &createdAt)
		account := entities.Account{ID, name, CPF, secret, balance, createdAt}
		accounts = append(accounts, account)
		checkError(err)
	}
	if len(accounts) == 0 {
		return entities.Account{}
	}
	return accounts[0]

}

///retornar erro (tratar)
func (r Repository) UpdateBalance(account entities.Account) {

	rows, err := r.Db.Exec("UPDATE account SET balance = ? WHERE id = ?", account.Balance, account.ID)
	checkError(err)
	rowCount, err := rows.RowsAffected()
	fmt.Println(rowCount)

}

func (r Repository) InsertAccount(accountInput entities.AccountInput) (*entities.Account, error) {
	var account entities.Account
	account = entities.NewAccount(accountInput.Name, accountInput.CPF, accountInput.Secret)

	account_exist := r.FindOne(accountInput.CPF)
	if !reflect.DeepEqual(account_exist, entities.Account{}) {
		return nil, errors.New("The CPF already exists.")
	}
	_, err := r.Db.Query("insert into  account (id, name, cpf, secret,balance, created_at) values (?,?,?,?,?,? )",
		account.ID, account.Name, account.CPF, account.Secret, account.Balance, account.CreatedAt)

	if err != nil {
		checkError(err)
		return nil, err
	}
	return &account, nil

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
