package main

import (
	"context"
	cfg "efmo-test/internal/config"
	controller "efmo-test/internal/controller/http/fiber"
	"efmo-test/internal/service"
	"efmo-test/internal/storage/pgx"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()

	config := cfg.LoadConfig("config/config.env", logger)
	config.Storage.SetURI(logger)

	storage := pgx.NewPGX(logger, config.Storage.GetURI())

	if err := storage.Ping(context.Background()); err != nil {
		logger.Debug("failed to ping to database", zap.Error(err))
	}

	srv := service.NewService(logger, storage)
	wApp := fiber.New()
	ctrl := controller.NewController(logger, srv, wApp)
	ctrl.ConfigureRoutes()

	quit := make(chan os.Signal, 1)

	go func() {
		if err := wApp.Listen(config.Service.Port); err != nil {
			logger.Fatal("Can't shutdown service", zap.Error(err))
		}
	}()

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	storage.Disconnect()
	logger.Info("Database disconnected")

	if err := wApp.Shutdown(); err != nil {
		logger.Info("Failed to stop server")
		return
	}

	logger.Info("Server stopped")
}
