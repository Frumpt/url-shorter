package save

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"url-sorter/internal/api/response"
	"url-sorter/internal/storage/database"
	"url-sorter/lib/random"
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

// TODO: move to config
const aliasLength = 6

//go:generate go run github.com/vektra/mockery/v2@v2.40.2 --name=URLSaver
type URLSaver interface {
	SaveURL(ctx context.Context, arg *database.SaveURLParams) (*database.Url, error)
}

func ErrAttr(err error) slog.Attr {
	return slog.Any("error", err)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "internal.api.save.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			//TODO: remake error to pretty in add func
			log.Error("failed to decode request", ErrAttr(err))

			render.JSON(w, r, response.Error("failed to decode request"))

			return
		}

		log.Info(
			"request body decoded",
			slog.Any("request", req),
		)

		if err = validator.New().Struct(req); err != nil {
			validErrors, ok := err.(validator.ValidationErrors)
			if !ok {
				render.JSON(w, r, response.Error("wrong type validator error"))
				return
			}
			//TODO: remake error to pretty in add func
			validError := response.ValidationError(validErrors)
			validErrorJSON, _ := json.Marshal(validError)
			log.Error("invalid request", ErrAttr(errors.New(string(validErrorJSON))))

			render.JSON(w, r, validError)
			return
		}
		alias := req.Alias
		if alias == "" {
			// TODO: If generated alias exist, make checking
			alias = random.NewRandomString(aliasLength)
		}

		res, err := urlSaver.SaveURL(context.Background(), (*database.SaveURLParams)(&req))
		if err != nil {
			//TODO: remake error to pretty in add func
			log.Error("error save url", ErrAttr(err))
			render.JSON(w, r, response.Error(err.Error()))
			return
		}
		log.Info("url added")

		render.JSON(w, r, Response{
			Response: response.OK(),
			Alias:    res.Alias,
		})
	}
}
