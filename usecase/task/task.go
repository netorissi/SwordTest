//go:generate mockgen -source=./task.go -destination=../../mocks/task_case_mock.go -package=mocks -mock_names=Usecase=MockTaskUsecase

package task

import (
	"context"

	"github.com/netorissi/SwordTest/infra/broker"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/repository"
)

type Usecase interface {
	Get(ctx context.Context) ([]model.Task, error)
	GetByUserID(ctx context.Context, userID int) ([]model.Task, error)
	Create(ctx context.Context, task model.Task) error
	Update(ctx context.Context, task model.Task) error
	Complete(ctx context.Context, taskID, userID int) error
	Del(ctx context.Context, taskID int) error
}

type impl struct {
	repository *repository.Repository
	broker     broker.MessageBroker
}

func New(repository *repository.Repository, broker broker.MessageBroker) Usecase {
	return &impl{repository, broker}
}
