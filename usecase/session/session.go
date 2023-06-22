//go:generate mockgen -source=./session.go -destination=../../mocks/session_case_mock.go -package=mocks -mock_names=Usecase=MockSessionUsecase

package session

import (
	"context"

	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/repository"
)

type Usecase interface {
	GetByToken(ctx context.Context, token string) (*model.Session, error)
}

type impl struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) Usecase {
	return &impl{repository}
}
