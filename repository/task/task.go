//go:generate mockgen -source=./task.go -destination=../../mocks/task_repo_mock.go -package=mocks -mock_names=Repository=MockTaskRepository

package task

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/netorissi/SwordTest/model"
)

type Repository interface {
	Get(ctx context.Context) ([]model.Task, error)
	GetByID(ctx context.Context, taskID int) (*model.Task, error)
	GetByUserID(ctx context.Context, userID int) ([]model.Task, error)
	Create(ctx context.Context, task model.Task) error
	Update(ctx context.Context, task model.Task) error
	Del(ctx context.Context, taskID int) error
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
