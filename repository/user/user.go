//go:generate mockgen -source=./user.go -destination=../../mocks/user_repo_mock.go -package=mocks -mock_names=Repository=MockUserRepository

package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/netorissi/SwordTest/model"
)

type Repository interface {
	GetByID(ctx context.Context, userID int) (*model.User, error)
}

type impl struct {
	mySQLReader *sqlx.DB
	mySQLWriter *sqlx.DB
}

type Options struct {
	MySQLReader *sqlx.DB
	MySQLWriter *sqlx.DB
}

func New(opts Options) Repository {
	return &impl{
		mySQLReader: opts.MySQLReader,
		mySQLWriter: opts.MySQLWriter,
	}
}
