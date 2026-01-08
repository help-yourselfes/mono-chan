package router

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/helpyourselfes/mono-chan/internal/app"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/repo"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/storage"
	mw "github.com/helpyourselfes/mono-chan/internal/router/middleware"
)

func SetupRouter(log *slog.Logger, db *sql.DB) *chi.Mux {
	api := chi.NewRouter()

	api.Use(middleware.Logger)
	api.Use(middleware.RequestID)
	api.Use(mw.InjectLogger(log))

	boardRepo := repo.NewSQLiteBoardRepo(db)
	postRepo := repo.NewSQLitePostRepo(db)
	threadRepo := repo.NewSQLiteThreadRepo(db)

	var repos *app.Repos = &app.Repos{
		Boards:  boardRepo,
		Threads: threadRepo,
		Posts:   postRepo,
	}

	txManager := storage.NewSqlTxManager(db)

	api.Mount("/boards", boardRoutes(log, boardRepo))
	api.Mount("/posts", postRoutes(log, repos, txManager))
	api.Mount("/threads", threadRoutes(repos, txManager))
	// api.Route("/view")

	return api
}
