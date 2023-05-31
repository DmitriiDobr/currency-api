package repository

import (
	"fmt"
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
	//db, err := sqlx.Open("postgres", "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable")
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.Username, c.DbName, c.Password, c.SslMode))
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
