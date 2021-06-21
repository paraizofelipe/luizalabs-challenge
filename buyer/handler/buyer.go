package handler

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/luizalabs-challenge/api"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/service"
)

type Buyer struct {
	Logger  *log.Logger
	Service service.Service
}

func NewHandler(db *sqlx.DB, logger *log.Logger) Buyer {
	return Buyer{
		Logger:  logger,
		Service: service.NewService(db),
	}
}

func (h Buyer) Router(w http.ResponseWriter, r *http.Request) {
	router := api.NewRouter(h.Logger)

	router.ServeHTTP(w, r)
}
