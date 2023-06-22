package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func NewMySQL() (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	return sqlx.NewDb(db, "sqlmock"), mock
}
