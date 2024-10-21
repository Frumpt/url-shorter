package redirect

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"url-sorter/internal/api/response"
	"url-sorter/internal/logger"
	"url-sorter/internal/storage/database"
)

type Request struct {
	ID    int32  `json:"id" validate:"required"`
	Alias string `json:"alias"`
	Url   string `json:"url" validate:"required,url"`
}

type Response struct {
	response.Response
	Alias string `json:"body,omitempty"`
}

//go:generate go run github.com/vektra/mockery/v2@v2.40.2 --name=URLDelete
type URLDelete interface {
	DeleteURL(ctx context.Context, arg *database.DeleteURLParams) error
}

func New(log *slog.Logger, urlGetter URLDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "internal.api.redirect.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request", middleware.GetReqID(r.Context())),
		)
		fmt.Printf("%+v\n", r)
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, response.Error("not found"))

			return
		}

		requestDelete := struct {
			Alias string
		}{Alias: alias}

		err := urlGetter.DeleteURL(context.Background(), (*database.DeleteURLParams)(&requestDelete))
		if err != nil {
			log.Error("error redirect url", logger.ErrAttr(err))
			render.JSON(w, r, response.Error(err.Error()))
			return
		}
		log.Info("url deleted")

		render.JSON(w, r, http.StatusOK)
	}
}
