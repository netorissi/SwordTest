package api

import (
	"github.com/labstack/echo/v4"
	"github.com/netorissi/SwordTest/api/middleware"
	v1 "github.com/netorissi/SwordTest/api/v1"
	"github.com/netorissi/SwordTest/config"
	"github.com/netorissi/SwordTest/shared"
	"github.com/netorissi/SwordTest/usecase"
)

type Options struct {
	Server  *echo.Echo
	Usecase *usecase.Usecase
}

func New(opts Options) {
	var (
		basePath = opts.Server.Group(config.Global.Server.BasePath)
		middle   = middleware.New(*opts.Usecase)
	)

	v1.New(v1.Options{
		API:        basePath,
		Usecase:    opts.Usecase,
		Middleware: middle,
	})

	shared.Logger.Info("API started")
}
