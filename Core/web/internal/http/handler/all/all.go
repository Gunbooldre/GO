package all

import (
	"log/slog"
	"net/http"
	resp "web/internal/lib/api/response"
	"web/internal/storage/sqllite"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type GetAllData interface {
	GetAll() ([]sqllite.Web, error)
}

type Response struct {
	resp.Response
	Urls []sqllite.Web `json:"urls"`
}

func New(logger *slog.Logger, getAllData GetAllData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.redirect.New"
		log := logger.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		res, err := getAllData.GetAll()
		if err != nil {
			render.JSON(w, r, resp.Error("Internal Error"))
			return
		}
		log.Info("GetALl")
		render.JSON(w, r, Response{
			Response: resp.OK(),
			Urls:     res,
		})
	}
}
