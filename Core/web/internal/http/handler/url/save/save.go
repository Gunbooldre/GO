package save

import (
	"log/slog"
	"net/http"
	resp "web/internal/lib/api/response"

	"github.com/go-chi/render"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	URL   string `json:"url" validate:"required, url"`
	Alias string `json:"alias" validate:"required"`
}

type Resposce struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveUrl(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.url.save.New"
		log = log.With(
			slog.String("operation", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		var req Request
		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error("failed to decode request body,", slog.Any("reg", err))
			render.JSON(w, r, resp.Error("Failed to decode request"))

			return
		}
		log.Info("Request body decoded", slog.Any("body", req))
		if err := validator.New().Struct(req); err != nil {
			validatorErr := err.(validator.ValidationErrors)
			log.Error("failed to validate request body", slog.Any("body", req))
			render.JSON(w, r, resp.ValidateErrors(validatorErr))
			return
		}
	}
}
