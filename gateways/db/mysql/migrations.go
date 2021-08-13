package mysql

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

//go:embed migrations
var _migrations embed.FS

func RunMigrations(dbURL string) error {
	source, err := httpfs.New(http.FS(_migrations), "migrations")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("httpfs", source, "mysql://"+dbURL)
	if err != nil {
		return err
	}

	defer m.Close()

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		return err
	}

	return nil

}
