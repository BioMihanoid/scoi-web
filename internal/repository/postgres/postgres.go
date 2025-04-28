package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable = "users"
)

func NewPostgresDb(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
