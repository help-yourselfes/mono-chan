package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/helpyourselfes/mono-chan/internal/app/board/handlers"
	"github.com/helpyourselfes/mono-chan/internal/app/board/repo"
	"github.com/helpyourselfes/mono-chan/internal/app/board/service"
)

func boardRoutes(repo repo.BoardRepo) chi.Router {
	service := service.NewBoardService(repo)
	handler := handlers.NewBoardHandler(service)

	r := chi.NewRouter()

	r.Get("/{key}", handler.GetBoardByKey)
	r.Get("/list", handler.GetBoardsList)

	r.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth("mono-chan", map[string]string{"admin": "admin"}))
		r.Post("/", handler.CreateBoard)
		r.Delete("/{key}", handler.DeleteBoard)
		r.Put("/", handler.UpdateBoard)
	})

	return r
}
