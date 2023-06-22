package repository

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/netorissi/SwordTest/config"
	"github.com/netorissi/SwordTest/infra/database"
	"github.com/netorissi/SwordTest/repository/session"
	"github.com/netorissi/SwordTest/repository/task"
	"github.com/netorissi/SwordTest/repository/user"
	"github.com/netorissi/SwordTest/shared"
)

type Repository struct {
	User    user.Repository
	Task    task.Repository
	Session session.Repository
}

func New() *Repository {
	var (
		mysqlReader = database.NewMySQLConn(config.Global.MySQL.ReaderURL)
		mysqlWriter = database.NewMySQLConn(config.Global.MySQL.WriterURL)
	)

	// using migrate only for tests
	driver, err := mysql.WithInstance(mysqlWriter.DB, &mysql.Config{})
	if err != nil {
		shared.Logger.DPanic(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./repository/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		shared.Logger.DPanic(err.Error())
	}

	m.Up()

	return &Repository{
		User:    user.New(user.Options{MySQLReader: mysqlReader, MySQLWriter: mysqlWriter}),
		Task:    task.New(task.Options{MySQLReader: mysqlReader, MySQLWriter: mysqlWriter}),
		Session: session.New(session.Options{MySQLReader: mysqlReader, MySQLWriter: mysqlWriter}),
	}
}
