package database

import "github.com/jmoiron/sqlx"

type Config struct {
	ConnectionRow string
}

func NewPostgressDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.ConnectionRow)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
