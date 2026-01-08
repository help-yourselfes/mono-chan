package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/helpyourselfes/mono-chan/internal/app"
	"github.com/helpyourselfes/mono-chan/internal/app/thread/handlers"
	"github.com/helpyourselfes/mono-chan/internal/app/thread/service"
)

func threadRoutes(repos *app.Repos, tx app.TransactionManager) chi.Router {
	srv := service.NewThreadService(repos, tx)
	handler := handlers.NewThreadHandler(srv)

	r := chi.NewRouter()

	r.Post("/", handler.Create)
	r.Get("/get", handler.GetByPostID)
	r.Get("/getPost", handler.GetWithPost)
	r.Put("/", handler.Update)
	r.Delete("/", handler.DeleteByUser)
	r.Get("/list", handler.List)

	r.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth("mono-chan", map[string]string{
			"admin": "admin", // TODO: make autorisation
		}))

		r.Delete("/delete", handler.DeleteByAdmin)
	})
	return r
}
