package redirect

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

type URLGetter interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.redirect.New"
		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)
		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, resp.Error("Not FOUND"))
			return
		}
		resURL, err := urlGetter.GetURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("Not FOUND", alias)
			render.JSON(w, r, resp.Error("Not FOUND"))
			return
		}
		if err != nil {
			log.Error(err.Error(), alias)
			render.JSON(w, r, resp.Error("Internal Error"))
			return
		}
		log.Info("Redirect", alias, resURL)
		http.Redirect(w, r, res]чсURL, http.StatusFound)
	}
}
