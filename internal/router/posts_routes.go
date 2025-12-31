package router

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/helpyourselfes/mono-chan/internal/app/post/handlers"
	"github.com/helpyourselfes/mono-chan/internal/app/post/service"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/repo"
)

func postRoutes(log *slog.Logger, storage *sql.DB) chi.Router {
	repo := repo.NewSQLitePostRepo(storage)
	srv := service.NewPostService(repo)
	handler := handlers.NewPostHandler(srv)

	r := chi.NewRouter()

	r.Post("/", handler.CreatePost(log))
	r.Get("/post/{boardKey}/{id}", handler.GetById(log))
	r.Put("/", handler.Update(log))
	r.Delete("/", handler.UserDelete(log))
	r.Get("/list/{boardKey}/{threadId}", handler.List(log))

	r.Route("/admin", func(r chi.Router) {
		r.Use(middleware.BasicAuth("mono-chan", map[string]string{
			"admin": "admin", // TODO: make autorisation
		}))
		r.Delete("/", handler.AdminDelete(log))
	})
	return r
}
