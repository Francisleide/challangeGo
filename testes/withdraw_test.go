package tests

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	ac "github.com/francisleide/ChallangeGo/domain/account/usecase"
	autentic "github.com/francisleide/ChallangeGo/domain/autenticOperations/usecase"
	au "github.com/francisleide/ChallangeGo/domain/auth/usecase"
	tr "github.com/francisleide/ChallangeGo/domain/transfer/usecase"
	"github.com/francisleide/ChallangeGo/gateways/db/repository"
	gateways "github.com/francisleide/ChallangeGo/gateways/http"
	_ "github.com/go-sql-driver/mysql"
)

func testWithdraw(t *testing.T) {
	var (
		host     = "localhost"
		database = "challengego"
		user     = "root"
		password = "123456"
	)

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully created connection to database.")

	defer db.Close()

	r := repository.NewRepository(db)
	auc := ac.NewAccountUc(*r)
	tuc := tr.NewTransfer(*r)
	a := au.NewAuth(*r)
	aut := autentic.NewAutentic(*r)
	x := gateways.NewApi(auc, tuc, aut, a)

	x.Run("localhost", "3000")

	
}
