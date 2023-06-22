//go:generate mockgen -source=./user.go -destination=../../mocks/user_case_mock.go -package=mocks -mock_names=Usecase=MockUserUsecase

package user

import (
	"context"

	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/repository"
)

type Usecase interface {
	GetByID(ctx context.Context, userID int) (*model.User, error)
}

type impl struct {
	repository *repository.Repository
}

func New(repository *repository.Repository) Usecase {
	return &impl{repository}
}
