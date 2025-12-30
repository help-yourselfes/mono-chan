package router

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(storage *sql.DB) *chi.Mux {
	api := chi.NewRouter()

	api.Use(middleware.RequestID)
	api.Use(middleware.Logger)

	api.Mount("/boards", boardRoutes(storage))
	api.Mount("/posts", boardRoutes(storage))
	// api.Route("/view")

	return api
}
