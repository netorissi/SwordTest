package task

import (
	"github.com/labstack/echo/v4"
	"github.com/netorissi/SwordTest/api/middleware"
	"github.com/netorissi/SwordTest/usecase"
)

type impl struct {
	usecase *usecase.Usecase
}

type Options struct {
	API        *echo.Group
	Usecase    *usecase.Usecase
	Middleware middleware.Middleware
}

func New(opts Options) {
	var (
		api = opts.API.Group("/tasks")
		h   = &impl{usecase: opts.Usecase}
	)

	// tech permissions
	api.POST("", h.create, opts.Middleware.Tech)
	api.PUT("/:id", h.update, opts.Middleware.Tech)
	api.PATCH("/:id/complete", h.complete, opts.Middleware.Tech)
	api.GET("/me", h.me, opts.Middleware.Tech)

	// manager permissions
	api.GET("", h.getAll, opts.Middleware.Manager)
	api.DELETE("/:id", h.del, opts.Middleware.Manager)
}
