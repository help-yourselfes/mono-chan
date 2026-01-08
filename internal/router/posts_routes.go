package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/helpyourselfes/mono-chan/internal/app"
	"github.com/helpyourselfes/mono-chan/internal/app/post/handlers"
	"github.com/helpyourselfes/mono-chan/internal/app/post/service"
)

func postRoutes(repos *app.Repos, tx app.TransactionManager) chi.Router {
	srv := service.NewPostService(repos, tx)
	handler := handlers.NewPostHandler(srv)

	r := chi.NewRouter()

	r.Post("/", handler.CreatePost)
	r.Get("/post/{boardKey}/{id}", handler.GetById)
	r.Put("/", handler.Update)
	r.Delete("/", handler.UserDelete)
	r.Get("/list/{boardKey}/{threadId}", handler.List)

	r.Route("/admin", func(r chi.Router) {
		r.Use(middleware.BasicAuth("mono-chan", map[string]string{
			"admin": "admin", // TODO: make autorisation
		}))
		r.Delete("/delete", handler.AdminDelete)
	})
	return r
}
