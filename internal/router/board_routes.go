package router

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/helpyourselfes/mono-chan/internal/app/board/handlers"
	"github.com/helpyourselfes/mono-chan/internal/app/board/service"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/repo"
)

func boardRoutes(log *slog.Logger, storage *sql.DB) chi.Router {
	repo := repo.NewSQLiteBoardRepo(storage)
	service := service.NewBoardService(repo)
	handler := handlers.NewBoardHandler(service)

	r := chi.NewRouter()

	r.Get("/{key}", handler.GetBoardByKey(log))
	r.Get("/list", handler.GetBoardsList(log))

	r.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth("mono-chan", map[string]string{"admin": "admin"}))
		r.Post("/", handler.CreateBoard(log))
		r.Delete("/{key}", handler.DeleteBoard(log))
		r.Put("/", handler.UpdateBoard(log))
	})

	return r
}
