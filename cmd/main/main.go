package main

import (
	"log/slog"
	"net/http"

	boardHandlers "github.com/helpyourselfes/mono-chan/internal/app/board/handlers"
	boardService "github.com/helpyourselfes/mono-chan/internal/app/board/service"
	postHandlers "github.com/helpyourselfes/mono-chan/internal/app/post/handlers"
	postService "github.com/helpyourselfes/mono-chan/internal/app/post/service"
	threadHandlers "github.com/helpyourselfes/mono-chan/internal/app/thread/handlers"
	threadService "github.com/helpyourselfes/mono-chan/internal/app/thread/service"
	"github.com/helpyourselfes/mono-chan/internal/config"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/repo"
	"github.com/helpyourselfes/mono-chan/internal/infrastructure/storage"
	"github.com/helpyourselfes/mono-chan/internal/logger"
	"github.com/helpyourselfes/mono-chan/internal/logger/sl"
	"github.com/helpyourselfes/mono-chan/internal/router"
)

func main() {
	cfg := config.Load()

	_ = cfg

	log := logger.GetLogger()

	db, err := storage.InitSQLiteStorage(
		cfg.StoragePath,
	)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		return
	}

	txManager := storage.NewSqlTxManager(db)

	boardRepo := repo.NewSQLiteBoardRepo(db)
	postRepo := repo.NewSQLitePostRepo(db)
	threadRepo := repo.NewSQLiteThreadRepo(db)

	boardService := boardService.NewBoardService(boardRepo)
	threadService := threadService.NewThreadService(boardRepo, threadRepo, postRepo, txManager)
	postService := postService.NewPostService(boardRepo, threadRepo, postRepo, txManager)

	boardHandler := boardHandlers.NewBoardHandler(boardService)
	threadHandler := threadHandlers.NewThreadHandler(threadService)
	postHandler := postHandlers.NewPostHandler(postService)

	r := router.SetupRouter(log, db, *boardHandler, *threadHandler, *postHandler)

	log.Info("server started", slog.String("address", cfg.HTTPServer.Address))

	srv := &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
		return
	}

	log.Error("server stopped")
}
