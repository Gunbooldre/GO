package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"web/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	log := setupLogger(cfg.Env)
	log.Info("Start", slog.String("env", cfg.Env))
	fmt.Println("Hello World")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case config.Envlocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
	return log
}
