package repository

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	Db  *sql.DB
	log *logrus.Entry
}

func NewRepository(db *sql.DB, log *logrus.Entry) *Repository {
	return &Repository{
		Db:  db,
		log: log,
	}

}
