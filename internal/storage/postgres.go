package storage

import (
	"github.com/jmoiron/sqlx"
)

func NewPostgresStorage(connStr string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
