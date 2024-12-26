package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Bakhram74/gw-exchanger/config"
	"github.com/Bakhram74/gw-exchanger/internal/grpc"

	"github.com/Bakhram74/gw-exchanger/internal/service"
	"github.com/Bakhram74/gw-exchanger/internal/storage/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.NewConfig()
	log := setupLogger(cfg.Env)
	slog.SetDefault(log)

	dbUrl := url(cfg)

	postgres, err := postgres.New(dbUrl)
	if err != nil {
		panic(err)
	}

	service := service.NewService(postgres)

	server := grpc.New(service, cfg.Port)
	go func() {
		server.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	server.Stop()
	log.Info("Gracefully stopped")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func url(cfg config.Config) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Storage.PostgresUsername,
		cfg.Storage.PostgresPassword,
		cfg.Storage.PostgresHost,
		cfg.Storage.PostgresPort,
		cfg.Storage.PostgresDatabase,
		cfg.Storage.PostgresSslMode)
}
