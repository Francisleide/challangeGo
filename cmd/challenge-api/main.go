package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/francisleide/ChallangeGo/docs"
	ac "github.com/francisleide/ChallangeGo/domain/account/usecase"
	autentic "github.com/francisleide/ChallangeGo/domain/autenticOperations/usecase"
	au "github.com/francisleide/ChallangeGo/domain/auth/usecase"
	tr "github.com/francisleide/ChallangeGo/domain/transfer/usecase"
	"github.com/francisleide/ChallangeGo/gateways/db/repository"
	gateways "github.com/francisleide/ChallangeGo/gateways/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/francisleide/ChallangeGo/app"
	mysqldb "github.com/francisleide/ChallangeGo/gateways/db/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func connect(mysql app.MysqlConfig) *sql.DB {
	//var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", user, password, host, database)
	//var connectionString = fmt.Sprint("root:123456@tcp(localhost:3306)/challengego?multiStatements=true")

	// Initialize connection object.

	db, err := sql.Open("mysql", mysql.DSN())

	if err != nil {
		log.Fatal("O erro está aqui: ", err)
	}
	/*driver, errm := mysql.WithInstance(db, &mysql.Config{})
	fmt.Println(errm)
	m, err := migrate.NewWithDatabaseInstance(
		"file://../../francisleide/ChallangeGo/gateways/db/mysql",
		"mysql",
		driver,
	)
	fmt.Println(err)

	m.Steps(2)
	checkError(err) */

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database.")
	return db
}

// @title Swagger Challenge API
// @version 2.0
// @description Documentation for Challenge-Go API

// TODO edit basepath !!!
// @BasePath /

func main() {

	config := app.ReadConfig(".env")

	db := connect(config.MysqlConfig)

	err := mysqldb.RunMigrations(config.MysqlConfig.URL())
	if(err == nil){
		fmt.Println("Migrations ok!")
	}
	if err != nil {
		log.Fatal("Error in db migrations! ", err)
	}
	defer db.Close()
	

	r := repository.NewRepository(db)
	auc := ac.NewAccountUc(*r)
	tuc := tr.NewTransfer(*r)
	a := au.NewAuth(*r)
	aut := autentic.NewAutentic(*r)
	x := gateways.NewApi(auc, tuc, aut, a)
	docs.SwaggerInfo.Host = "localhost:8080"
	x.Run("0.0.0.0", "8080")
}
