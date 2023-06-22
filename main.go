package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/netorissi/SwordTest/api"
	"github.com/netorissi/SwordTest/config"
	"github.com/netorissi/SwordTest/event"
	"github.com/netorissi/SwordTest/infra/broker"
	"github.com/netorissi/SwordTest/infra/server"
	"github.com/netorissi/SwordTest/repository"
	"github.com/netorissi/SwordTest/usecase"
)

var srv server.Server

func init() {
	config.LoadConfig()
	srv = server.New()
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Esta é uma rota privada
func main() {
	complete := make(chan struct{})

	// start repositories
	repository := repository.New()

	// start message broker
	messagebroker := broker.New()

	// start events subscribers
	event.New(messagebroker)

	// start usecases
	usecase := usecase.New(usecase.Options{
		Repository: repository,
		Broker:     messagebroker,
	})

	// start api http
	api.New(api.Options{
		Server:  srv.GetInstance(),
		Usecase: usecase,
	})

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit

		srv.Stop()

		close(complete)
	}()

	// start server
	srv.Start()

	<-complete
}
