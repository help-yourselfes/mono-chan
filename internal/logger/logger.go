package logger

import (
	"context"
	"log/slog"

	"github.com/helpyourselfes/mono-chan/internal/router/middleware"
)

func GetLogger() *slog.Logger {
	return slog.Default()
}

func FromContext(ctx context.Context) *slog.Logger {
	if log, ok := ctx.Value(middleware.LoggerKey).(*slog.Logger); ok {
		return log
	}

	return slog.Default()
}
