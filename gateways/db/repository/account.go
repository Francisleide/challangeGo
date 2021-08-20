package repository

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/francisleide/ChallangeGo/domain/entities"
)

var (
	ID        string
	nome      string
	CPF       string
	secret    string
	balance   float64
	createdAt string
)
var accounts []entities.Account

func (r Repository) ListAllAccounts() []entities.Account {

	rows, err := r.Db.Query("SELECT id, nome, cpf, secret,balance, created_at from account;")
	defer rows.Close()
	checkError(err)
	fmt.Println("Reading data:")
	for rows.Next() {
		err = rows.Scan(&ID, &nome, &CPF, &secret, &balance, &createdAt)
		account := entities.Account{ID, nome, CPF, secret, balance, createdAt}
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
	sql = "SELECT id, nome, cpf, secret,balance, created_at from account where cpf=?"
	rows, err := r.Db.Query(sql, CPF)
	fmt.Println(sql, CPF)
	fmt.Println("Quantidade de linhas: ", len(accounts))
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&ID, &nome, &CPF, &secret, &balance, &createdAt)
		account := entities.Account{ID, nome, CPF, secret, balance, createdAt}
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
	account = entities.NewAccount(accountInput.Nome, accountInput.CPF, accountInput.Secret)
	fmt.Println("CPF no Repository: ", account.CPF)
	fmt.Printf(account.ID)
	account_exist := r.FindOne(accountInput.CPF)
	if !reflect.DeepEqual(account_exist, entities.Account{}) {
		return nil, errors.New("Já exite este CPF no banco.")
	}
	_, err := r.Db.Query("insert into  account (id, nome, cpf, secret,balance, created_at) values (?,?,?,?,?,? )",
		account.ID, account.Nome, account.CPF, account.Secret, account.Balance, account.CreatedAt)

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
