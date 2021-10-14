package main

import (
	"database/sql"
	"os"

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
	"github.com/sirupsen/logrus"
)

func checkError(err error, log logrus.Entry) {
	if err != nil {
		log.Fatal(err)
	}
}

func connect(mysql app.MysqlConfig, log logrus.Entry) *sql.DB {
	db, err := sql.Open("mysql", mysql.DSN())

	if err != nil {
		log.Fatal("error: ", err)
	}

	err = db.Ping()
	checkError(err, log)
	return db
}

// @title Swagger Challenge API
// @version 2.0
// @description Documentation for Challenge-Go API
// @BasePath /

func main() {
	log := logrus.New()

	logEntry := logrus.NewEntry(log)
	config := app.ReadConfig(".env")
	log.SetOutput(os.Stdout)

	db := connect(config.MysqlConfig, *logEntry)

	err := mysqldb.RunMigrations(config.MysqlConfig.URL())
	checkError(err, *logEntry)
	if err == nil {
		logEntry.Infoln("migrations ok!")
	}

	defer db.Close()

	r := repository.NewRepository(db, logEntry)
	accountUsecase := account.NewAccountUc(*r, logEntry)
	transferUseCase := transfer.NewTransferUC(*r, logEntry)
	authenticationUsecase := authentication.NewAuthenticationUC(*r)
	api := gateways.NewApi(accountUsecase, transferUseCase, authenticationUsecase, logEntry)
	docs.SwaggerInfo.Host = "localhost:8080"
	api.Run("0.0.0.0", "8080")
}
