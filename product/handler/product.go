package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/luizalabs-challenge/api"
	"github.com/paraizofelipe/luizalabs-challenge/product/service"
)

type Product struct {
	Logger  *log.Logger
	Service service.Service
}

func NewHandler(db *sqlx.DB, logger *log.Logger) Product {
	return Product{
		Logger:  logger,
		Service: service.NewService(db),
	}
}

func (h Product) list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(SuccessResponse{Msg: "Tudo OK!"})
		if err != nil {
			http.Error(w, ErrorResponse{Error: "failed to parse json"}.String(), http.StatusInternalServerError)
		}
	}
}

func (h Product) detail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (h Product) Router(w http.ResponseWriter, r *http.Request) {
	router := api.NewRouter(h.Logger)

	router.AddRoute(`product\/\?page=\d`, http.MethodGet, h.list())
	router.AddRoute(`product\/?$`, http.MethodGet, h.detail())

	router.ServeHTTP(w, r)
}
