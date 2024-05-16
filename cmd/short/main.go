package main

import (
	"cmd/short/main.go/internal/app"
	"cmd/short/main.go/internal/config"
	"cmd/short/main.go/internal/lib/logger/sl"
	"cmd/short/main.go/internal/storage/pg"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	log.Info("starting short")

	storage, err := pg.New(cfg.DB)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	application := app.New(log, storage)

	application.InitApp()
	go application.StartApp()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	s := <-stop
	log.Info("stopping application", slog.String("signal", s.String()))

	application.StopApp()
	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				AddSource:   false,
				Level:       slog.LevelDebug,
				ReplaceAttr: nil,
			}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource:   false,
				Level:       slog.LevelDebug,
				ReplaceAttr: nil,
			}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				AddSource:   false,
				Level:       slog.LevelInfo,
				ReplaceAttr: nil,
			}),
		)
	}

	return log
}
