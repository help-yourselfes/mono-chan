package router

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(log *slog.Logger, storage *sql.DB) *chi.Mux {
	api := chi.NewRouter()

	api.Use(middleware.RequestID)
	api.Use(middleware.Logger)

	api.Mount("/boards", boardRoutes(log, storage))
	api.Mount("/posts", postRoutes(log, storage))
	// api.Route("/view")

	return api
}
