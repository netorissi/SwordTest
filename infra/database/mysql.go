package database

import (
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
	"github.com/netorissi/SwordTest/shared"
)

func NewMySQLConn(url string) *sqlx.DB {
	db := sqlx.MustConnect("mysql", url)

	shared.Logger.Info("MySQL connected.")

	return db
}
