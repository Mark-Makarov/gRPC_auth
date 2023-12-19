package main

import (
	"log/slog"
	"os"
	"sso/internal/app"
	"sso/internal/config"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("app started on port", slog.Int("port", cfg.GRPC.Port))

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	application.GRPCSrv.MustRun()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envDev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
