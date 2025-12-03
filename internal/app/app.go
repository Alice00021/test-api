// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/Alice00021/test_api/internal/di"
	"os"
	"os/signal"
	"syscall"

	"github.com/Alice00021/test_common/pkg/rabbitmq/rmq_rpc/client"

	"github.com/gin-gonic/gin"

	"github.com/Alice00021/test_api/config"
	amqprpc "github.com/Alice00021/test_api/internal/controller/amqp_rpc"
	"github.com/Alice00021/test_api/internal/controller/http"
	"github.com/Alice00021/test_common/pkg/httpserver"
	"github.com/Alice00021/test_common/pkg/logger"
	"github.com/Alice00021/test_common/pkg/postgres"
	"github.com/Alice00021/test_common/pkg/rabbitmq/rmq_rpc/server"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.NewMultipleWriter(
		logger.Level(cfg.Log.Level),
		logger.FileName(cfg.Log.FileName),
	)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// RabbitMQ RPC Client
	rmqClient, err := client.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, cfg.RMQ.ClientExchange, cfg.App.Name, cfg.RMQ.ClientPrefix)
	if err != nil {
		l.Fatal("RabbitMQ RPC Client - init error - client.New")
	}

	// UseCase
	uc := di.NewUseCase(rmqClient, wsServer, cfg)

	// RabbitMQ RPC Server
	rmqRouter := amqprpc.NewRouter(uc, l)

	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, cfg.App.Name, rmqRouter, l, cfg.RMQ.ClientPrefix)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	// HTTP Server
	handler := gin.New()
	http.NewRouter(handler, cfg, l, uc, wsServer)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	case err = <-rmqServer.Notify():
		l.Error(fmt.Errorf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	err = rmqServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}

	err = rmqClient.Shutdown()
	if err != nil {
		l.Fatal("RabbitMQ RPC Client - shutdown error - rmqClient.RemoteCall", err)
	}

	wsServer.Shutdown()
}
