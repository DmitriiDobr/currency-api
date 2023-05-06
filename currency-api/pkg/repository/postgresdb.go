package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Password string
	Username string
	DbName   string
	SslMode  string
}

func (c *Config) InitDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
