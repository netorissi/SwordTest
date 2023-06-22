//go:generate mockgen -source=./session.go -destination=../../mocks/session_repo_mock.go -package=mocks -mock_names=Repository=MockSessionRepository

package session

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/netorissi/SwordTest/model"
)

type Repository interface {
	GetByToken(ctx context.Context, token string) (*model.Session, error)
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
