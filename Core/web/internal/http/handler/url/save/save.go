package save

import (
	"log/slog"
	"net/http"
	"web/internal/lib/api/response"

	"github.com/go-chi/chi/v5/middleware"
)

type Request struct {
	URL   string `json:"url" validate:"required, url"`
	Alias string `json:"alias" validate:"required"`
}

type Responce struct {
	response.Response
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
			log.Error("failed to decode request body,", slog.Err(err))
			render.JSON(w, r, resp.Error("Failed to decode request"))

			return
		}
	}
}
