package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/francisleide/ChallengeGo/docs"
	account "github.com/francisleide/ChallengeGo/domain/account/usecase"
	authentication "github.com/francisleide/ChallengeGo/domain/auth/usecase"
	transfer "github.com/francisleide/ChallengeGo/domain/transfer/usecase"
	"github.com/francisleide/ChallengeGo/gateways/db/repository"
	gateways "github.com/francisleide/ChallengeGo/gateways/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/francisleide/ChallengeGo/app"
	mysqldb "github.com/francisleide/ChallengeGo/gateways/db/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func connect(mysql app.MysqlConfig) *sql.DB {
	db, err := sql.Open("mysql", mysql.DSN())

	if err != nil {
		log.Fatal("error: ", err)
	}

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
	if err == nil {
		fmt.Println("Migrations ok!")
	}
	if err != nil {
		log.Fatal("Error in db migrations! ", err)
	}
	defer db.Close()

	r := repository.NewRepository(db)
	accountUsecase := account.NewAccountUc(*r)
	transferUseCase := transfer.NewTransferUC(*r)
	authenticationUsecase := authentication.NewAuthenticationUC(*r)
	api := gateways.NewApi(accountUsecase, transferUseCase, authenticationUsecase)
	docs.SwaggerInfo.Host = "localhost:8080"
	api.Run("0.0.0.0", "8080")
}
