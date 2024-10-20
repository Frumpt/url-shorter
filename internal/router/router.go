package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"url-sorter/internal/logger"
)

func NewRouter(log *slog.Logger) *chi.Mux {
	router := chi.NewRouter()

	loggerMiddleware := logger.NewMiddleware(log)
	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.URLFormat,
		loggerMiddleware,
	)
	return router
}
