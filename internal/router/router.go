package router

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	boardHandler "github.com/helpyourselfes/mono-chan/internal/app/board/handlers"
	postHandler "github.com/helpyourselfes/mono-chan/internal/app/post/handlers"
	threadHandler "github.com/helpyourselfes/mono-chan/internal/app/thread/handlers"
	mw "github.com/helpyourselfes/mono-chan/internal/router/middleware"
)

func SetupRouter(log *slog.Logger, db *sql.DB, boards boardHandler.BoardHandler, threads threadHandler.ThreadHandler, posts postHandler.PostHandler) *chi.Mux {
	api := chi.NewRouter()

	api.Use(middleware.Logger)
	api.Use(middleware.RequestID)
	api.Use(mw.InjectLogger(log))

	api.Route("/boards", func(r chi.Router) {
		r.Get("/", boards.GetBoardsList)
		r.Get("/{key}", boards.GetBoardByKey)

		r.Group(func(r chi.Router) {
			r.Use(middleware.BasicAuth("mono-chan", map[string]string{"admin": "admin"}))
			r.Post("/", boards.CreateBoard)
			r.Put("/{key}", boards.UpdateBoard)
			r.Delete("/{key}", boards.DeleteBoard)
		})
	})

	api.Route("/boards/{boardKey}/threads", func(r chi.Router) {
		r.Get("/", threads.List)
		r.Get("/{threadId}", threads.GetByPostID)
		r.Get("/{threadId}", threads.GetWithPost)
		r.Post("/", threads.Create)
		r.Put("/{threadId}", threads.Update)
		r.Delete("/{threadId}", threads.DeleteByUser)

		r.Group(func(r chi.Router) {
			r.Use(middleware.BasicAuth("mono-chan", map[string]string{
				"admin": "admin", // TODO: make autorisation
			}))

			r.Delete("/delete", threads.DeleteByAdmin)
		})
	})

	api.Route("/boards/{boardKey}/threads/{threadId}/posts", func(r chi.Router) {

		r.Post("/", posts.CreatePost)
		r.Get("/post/{boardKey}/{id}", posts.GetById)
		r.Put("/", posts.Update)
		r.Delete("/", posts.UserDelete)
		r.Get("/list", posts.List)

		r.Route("/admin", func(r chi.Router) {
			r.Use(middleware.BasicAuth("mono-chan", map[string]string{
				"admin": "admin", // TODO: make autorisation
			}))
			r.Delete("/delete", posts.AdminDelete)
		})
	})

	return api
}
