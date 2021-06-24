package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/domain"
	"github.com/paraizofelipe/luizalabs-challenge/buyer/service"
	"github.com/paraizofelipe/luizalabs-challenge/router"
)

type Buyer struct {
	Logger  *log.Logger
	Service service.Service
}

func NewHandler(db *sqlx.DB, logger *log.Logger) Handler {
	return &Buyer{
		Logger:  logger,
		Service: service.NewService(db),
	}
}

func (h Buyer) create(ctx *router.Context) {
	var (
		err   error
		buyer domain.Buyer
	)

	if err = json.NewDecoder(ctx.Body).Decode(&buyer); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusBadRequest, "Error when creating the buyer!")
		return
	}

	if err = h.Service.Add(buyer); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Error when creating the buyer!")
		return
	}

	ctx.Text(http.StatusCreated, "Buyer created")
}

func (h Buyer) remove(ctx *router.Context) {
	var (
		err error
		id  string = ctx.Params["id"]
	)

	if err = h.Service.RemoveByID(id); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusBadRequest, "Error when removing the buyer!")
		return
	}

	ctx.Text(http.StatusOK, "Buyer removed")
}

func (h Buyer) update(ctx *router.Context) {
	var (
		err   error
		id    string = ctx.Params["id"]
		buyer domain.Buyer
	)

	if err = json.NewDecoder(ctx.Body).Decode(&buyer); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusBadRequest, "Error when updating the buyer!")
		return
	}
	if buyer.ID, err = uuid.Parse(id); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusBadRequest, "Invalid ID!")
		return
	}
	if err = h.Service.Update(buyer); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Error when creating the buyer!")
		return
	}

	ctx.JSON(http.StatusOK, buyer)
}

func (h Buyer) detail(ctx *router.Context) {
	var (
		err   error
		id    string = ctx.Params["id"]
		buyer domain.Buyer
	)

	if _, err = uuid.Parse(id); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Invalid ID!")
		return
	}
	if buyer, err = h.Service.FindByID(id); err != nil {
		ctx.Text(http.StatusInternalServerError, "Error when fetching the buyer!")
		return
	}

	ctx.JSON(http.StatusOK, buyer)
}

func (h Buyer) addFavoriteProduct(ctx *router.Context) {
	var (
		err       error
		id        string = ctx.Params["id"]
		productID string = ctx.Params["id_product"]
	)

	if _, err = uuid.Parse(id); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Invalid ID of buyer!")
		return
	}
	if _, err = uuid.Parse(productID); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Invalid ID of product!")
		return
	}
	if err = h.Service.AddFavoriteProduct(id, productID); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Invalid ID of product!")
		return
	}
}

func (h Buyer) Router(w http.ResponseWriter, r *http.Request) {
	router := router.NewRouter(h.Logger)

	router.AddRoute(`buyer/(?P<id>[\w|-]+)/?`, http.MethodGet, h.detail)
	router.AddRoute(`buyer/(?P<id>[\w|-]+)/?`, http.MethodDelete, h.remove)
	router.AddRoute(`buyer/(?P<id>[\w|-]+)/?`, http.MethodPatch, h.update)
	router.AddRoute(`buyer/(?P<id>[\w|-]+)/product/(?P<id_product>[\w|-]+)?`, http.MethodPost, h.addFavoriteProduct)
	router.AddRoute(`buyer/?`, http.MethodPost, h.create)

	router.ServeHTTP(w, r)
}
