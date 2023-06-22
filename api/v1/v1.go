package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/netorissi/SwordTest/api/middleware"
	"github.com/netorissi/SwordTest/api/v1/task"
	"github.com/netorissi/SwordTest/usecase"
)

type Options struct {
	API        *echo.Group
	Usecase    *usecase.Usecase
	Middleware middleware.Middleware
}

func New(opts Options) {
	api := opts.API.Group("/v1")

	task.New(task.Options{
		API:        api,
		Usecase:    opts.Usecase,
		Middleware: opts.Middleware,
	})
}
