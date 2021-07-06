package repository

import (
	"fmt"
	"log"
	"reflect"

	"github.com/francisleide/ChallangeGo/domain/entities"
)

var (
	id         string
	nome       string
	cpf        string
	secret     string
	balance    float64
	created_at string
)
var accounts []entities.Account

func (r Repository) List_all_accounts() []entities.Account {

	rows, err := r.Db.Query("SELECT id, nome, cpf, secret,balance, created_at from account;")
	defer rows.Close()
	checkError(err)
	fmt.Println("Reading data:")
	for rows.Next() {
		err = rows.Scan(&id, &nome, &cpf, &secret, &balance, &created_at)
		account := entities.Account{id, nome, cpf, secret, balance, created_at}
		accounts = append(accounts, account)

	}
	checkError(err)
	err = rows.Err()
	fmt.Println("Done.")
	return accounts
}


func (r Repository) FindOne(cpf string) entities.Account {

	var accounts []entities.Account
	var sql string
	sql = "SELECT id, nome, cpf, secret,balance, created_at from account where cpf=?"
	rows, err := r.Db.Query(sql, cpf)
	fmt.Println(sql, cpf)
	
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&id, &nome, &cpf, &secret, &balance, &created_at)
		account := entities.Account{id, nome, cpf, secret, balance, created_at}
		accounts = append(accounts, account)

		checkError(err)
	}
	fmt.Println("CPF da conta : ", accounts[0].Cpf)
	if(len(accounts)== 0){
		return entities.Account{}
	}
	return accounts[0]

}

func (r Repository) UpdateBalance(account entities.Account) {

	rows, err := r.Db.Exec("UPDATE account SET balance = ? WHERE id = ?", account.Balance, account.Id)
	checkError(err)
	rowCount, err := rows.RowsAffected()
	fmt.Println(rowCount)
}

func (r Repository) InsertAccount(accountInput entities.AccountInput) (*entities.Account, error) {
	var account entities.Account
	account = entities.NewAccount(accountInput.Nome, accountInput.Cpf, accountInput.Secret)
	fmt.Println("CPF no Repository: ", account.Cpf)
	fmt.Printf(account.Id)
	account_exist := r.FindOne(accountInput.Cpf)
	if !reflect.DeepEqual(account_exist, entities.Account{}){
		log.Fatal("Já existe uma conta com este CPF")
	}
	_, err := r.Db.Query("insert into  account (id, nome, cpf, secret,balance, created_at) values (?,?,?,?,?,? )",
		account.Id, account.Nome, account.Cpf, account.Secret, account.Balance, account.Created_at)

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
