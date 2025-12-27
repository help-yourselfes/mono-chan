package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/helpyourselfes/mono-chan/internal/config"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/storage"
	"github.com/helpyourselfes/mono-chan/internal/logger"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
)

func main() {
	cfg := config.Load()

	_ = cfg

	log := logger.GetLogger()

	storage, err := storage.InitSQLiteStorage(
		cfg.StoragePath,
	)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		return
	}

	_ = storage

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	log.Info("server started", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
		return
	}

	log.Error("server stopped")
}
