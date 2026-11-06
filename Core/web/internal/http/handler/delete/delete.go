package delete

import (
	"errors"
	"log/slog"
	"net/http"
	resp "web/internal/lib/api/response"
	"web/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type UrlDelete interface {
	DeleteUrl(alias string) (string, error)
}

func New(log *slog.Logger, deleteUrl UrlDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "http.handler.delete.UrlDelete"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, resp.Error("Not FOUND"))
			return
		}
		resURL, err := deleteUrl.DeleteUrl(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", slog.String("alias", alias))
			render.JSON(w, r, resp.Error("Not FOUND"))
			return
		}
		if err != nil {
			log.Error("failed to delete url", slog.String("alias", alias), slog.String("err", err.Error()))
			render.JSON(w, r, resp.Error("Internal Error"))
			return
		}
		log.Info("Delete", slog.String("alias", alias), slog.String("url", resURL))
		render.JSON(w, r, resp.OK())
	}
}
