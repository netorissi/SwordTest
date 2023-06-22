package usecase

import (
	"github.com/netorissi/SwordTest/infra/broker"
	"github.com/netorissi/SwordTest/repository"
	"github.com/netorissi/SwordTest/usecase/session"
	"github.com/netorissi/SwordTest/usecase/task"
	"github.com/netorissi/SwordTest/usecase/user"
)

type Usecase struct {
	User    user.Usecase
	Task    task.Usecase
	Session session.Usecase
}

type Options struct {
	Repository *repository.Repository
	Broker     broker.MessageBroker
}

func New(opts Options) *Usecase {
	return &Usecase{
		User:    user.New(opts.Repository),
		Task:    task.New(opts.Repository, opts.Broker),
		Session: session.New(opts.Repository),
	}
}
