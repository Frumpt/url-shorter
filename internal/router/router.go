package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log/slog"
	"url-sorter/internal/api/handlers/url/redirect"
	"url-sorter/internal/api/handlers/url/save"
	"url-sorter/internal/logger"
	"url-sorter/internal/storage"
)

func NewRouter(log *slog.Logger, stg *storage.Storage) *chi.Mux {
	router := chi.NewRouter()

	loggerMiddleware := logger.NewMiddleware(log)
	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.URLFormat,
		loggerMiddleware,
	)

	router.Post("/url", save.New(log, stg))
	router.Get("/{alias}", redirect.New(log, stg))

	return router
}
