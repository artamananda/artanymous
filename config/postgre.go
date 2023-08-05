package config

import (
	"database/sql"

	"github.com/artamananda/artanymous/helper"
)

func NewDB() *sql.DB {
	connStr := "user=postgres dbname=test_db_arta password=postgres host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)
	return db
}
