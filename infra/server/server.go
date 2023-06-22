package server

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/netorissi/SwordTest/api/swagger"
	"github.com/netorissi/SwordTest/config"
	"github.com/netorissi/SwordTest/shared"
)

type Server interface {
	Start()
	Stop()
	GetInstance() *echo.Echo
}

type serverImpl struct {
	instance *echo.Echo
}

func New() Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())

	e.Validator = NewValidator()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		if err := c.JSON(shared.GetHTTPCode(err), shared.Response{Err: err}); err != nil {
			shared.Logger.Errorf("Error handling error: %v", err)
		}
	}

	return &serverImpl{
		instance: e,
	}
}

func (s *serverImpl) Start() {
	swagger.New(s.instance.Group("/swagger"))

	shared.Logger.Infof("Server started on %s.", config.Global.Server.Port)

	go func() {
		if err := s.instance.Start(config.Global.Server.Port); err != nil && err != http.ErrServerClosed {
			shared.Logger.Fatalf("Server start error: %v", err)
		}
	}()
}

func (s *serverImpl) Stop() {
	shared.Logger.Info("Server stoping...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.instance.Shutdown(ctx); err != nil {
		shared.Logger.Errorf("Server shutdown error: %v", err)
	}
}

func (s *serverImpl) GetInstance() *echo.Echo {
	return s.instance
}
