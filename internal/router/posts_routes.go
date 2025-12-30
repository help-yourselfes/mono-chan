package router

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/helpyourselfes/mono-chan/internal/app/post/handlers"
	"github.com/helpyourselfes/mono-chan/internal/app/post/service"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/repo"
)

func postRoutes(storage *sql.DB) chi.Router {
	repo := repo.NewSQLitePostRepo(storage)
	srv := service.NewPostService(repo)
	handler := handlers.NewPostHandler(srv)

	r := chi.NewRouter()

	r.Post("/", handler.CreatePost)

	return r
}
