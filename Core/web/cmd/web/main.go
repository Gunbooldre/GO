package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"web/internal/config"
	"web/internal/storage/sqllite"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	log := setupLogger(cfg.Env)
	log.Info("Start", slog.String("env", cfg.Env))
	fmt.Println("Hello World")

	storage, err := sqllite.New(cfg.StoragePath)
	if err != nil {
		log.Error("sqllite.New", "err", err)
	}

	//id, err := storage.SaveUrl("https://google.com", "google")
	//if err != nil {
	//	log.Error("storage.SaveUrl", "err", err, "id", id)
	//} else {
	//	log.Info("URL SAVED", id)
	//}

	res, err := storage.GetUrl(1)
	if err != nil {
		log.Info("storage.GetUrl", "err", err)
	} else {
		log.Info("storage.GetUrl", "res", res)
	}

	_ = storage
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
