package middleware

import (
	"context"
	"log/slog"
	"net/http"
)

type key int

var LoggerKey key

func InjectLogger(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), LoggerKey, log)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
