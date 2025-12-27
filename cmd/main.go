package main

import (
	"log/slog"
	"net/http"

	"github.com/helpyourselfes/mono-chan/internal/config"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/storage"
	"github.com/helpyourselfes/mono-chan/internal/logger"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	"github.com/helpyourselfes/mono-chan/internal/router"
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
	r := router.SetupRouter(storage)

	log.Info("server started", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      r,
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
