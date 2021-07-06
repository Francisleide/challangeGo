package main

import (
	"database/sql"
	"fmt"
	"log"

	ac "github.com/francisleide/ChallangeGo/domain/account/usecase"
	autentic "github.com/francisleide/ChallangeGo/domain/autenticOperations/usecase"
	au "github.com/francisleide/ChallangeGo/domain/auth/usecase"
	tr "github.com/francisleide/ChallangeGo/domain/transfer/usecase"
	"github.com/francisleide/ChallangeGo/gateways"
	"github.com/francisleide/ChallangeGo/gateways/repository"
	_ "github.com/go-sql-driver/mysql"
)

var (
	host     = "localhost"
	database = "challengego"
	user     = "root"
	password = "123456"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func connect() *sql.DB {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")
	return db
}

func main() {
	db := connect()
	defer db.Close()

	r := repository.NewRepository(db)
	auc := ac.NewAccountUc(*r)
	tuc := tr.NewTransfer(*r)
	a := au.NewAuth(*r)
	aut := autentic.NewAutentic(*r)
	x := gateways.NewApi(auc, tuc, aut, a)

	x.Run("localhost", "3000")
}
