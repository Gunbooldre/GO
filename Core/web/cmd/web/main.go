package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"web/internal/config"
	"web/internal/http/handler/delete"
	"web/internal/http/handler/redirect"
	"web/internal/http/handler/url/save"
	"web/internal/storage/sqllite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	res, err := storage.GetUrl(1)
	if err != nil {
		log.Info("storage.GetUrl", "err", err)
	} else {
		log.Info("storage.GetUrl", "res", res)
	}

	_ = storage

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", save.New(log, storage))
	router.Get("/{alias}", redirect.New(log, storage))
	router.Delete("/{alias}", delete.New(log, storage))
	log.Info("Start", slog.String("env", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:              cfg.HTTPServer.Address,
		Handler:           router,
		ReadHeaderTimeout: cfg.HTTPServer.Timeout,
		WriteTimeout:      cfg.HTTPServer.Timeout,
		IdleTimeout:       cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("srv.ListenAndServe", "err", err)
	}
	log.Error("Server Stopped")
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
