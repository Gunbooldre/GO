package save

import (
	"errors"
	"log/slog"
	"net/http"
	"web/internal/config"
	resp "web/internal/lib/api/response"
	"web/internal/lib/random"
	"web/internal/storage"

	"github.com/go-chi/render"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	URL   string `json:"url" validate:"required, url"`
	Alias string `json:"alias" validate:"required"`
}

type Response struct {
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

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(config.AliasLength)
		}

		id, err := urlSaver.SaveUrl(req.URL, alias)
		if errors.Is(err, storage.ErrUrlExist) {
			log.Info("url already exists", slog.Any("url", req.URL))
			render.JSON(w, r, resp.Error("URL already exists"))
			return
		}
		if err != nil {
			log.Error("Failed to add URL", slog.Any("err", err))
			render.JSON(w, r, resp.Error("Failed to add URL"))
			return
		}
		log.Info("URL added", slog.Int64("ID", id))
		responseOK(w, r, alias)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, alias string) {
	render.JSON(w, r, Response{
		Response: resp.OK(),
		Alias:    alias,
	})
}
