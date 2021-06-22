package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/paraizofelipe/luizalabs-challenge/product/domain"
	"github.com/paraizofelipe/luizalabs-challenge/product/service"
	"github.com/paraizofelipe/luizalabs-challenge/router"
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

func (h Product) create(ctx *router.Context) {
	var (
		err     error
		product domain.Product
	)

	if err = json.NewDecoder(ctx.Body).Decode(&product); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusBadRequest, "Error when creating the product!")
		return
	}

	if err = h.Service.Add(product); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Error when creating the product!")
		return
	}

	ctx.Text(http.StatusCreated, "Product created")
}

func (h Product) update(ctx *router.Context) {
	var (
		err     error
		id      string = ctx.Params["id"]
		product domain.Product
	)

	if err = json.NewDecoder(ctx.Body).Decode(&product); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusBadRequest, "Error when updating the product!")
		return
	}
	if product.ID, err = uuid.Parse(id); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Invalid ID!")
		return
	}
	if err = h.Service.Update(product); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Error when creating the product!")
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h Product) remove(ctx *router.Context) {
	var (
		err error
		id  string = ctx.Params["id"]
	)

	if _, err = uuid.Parse(id); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Invalid ID!")
		return
	}
	if err = h.Service.RemoveByID(id); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusBadRequest, "Error when remmoving the product!")
		return
	}

	ctx.Text(http.StatusOK, "Product removed")
}

func (h Product) list(ctx *router.Context) {
	var (
		err          error
		page         int
		listProducts []domain.Product
	)

	if page, err = strconv.Atoi(ctx.QueryString.Get("page")); err != nil || page <= 0 {
		ctx.Text(http.StatusInternalServerError, "Invalid page value")
		h.Logger.Println(err)
		return
	}
	if listProducts, err = h.Service.ListByPage(page - 1); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Error when listing the products!")
		return
	}

	ctx.JSON(http.StatusOK, listProducts)
}

func (h Product) detail(ctx *router.Context) {
	var (
		err     error
		id      string = ctx.Params["id"]
		product domain.Product
	)

	if _, err = uuid.Parse(id); err != nil {
		h.Logger.Println(err)
		ctx.Text(http.StatusInternalServerError, "Invalid ID!")
		return
	}
	if product, err = h.Service.FindByID(id); err != nil {
		ctx.Text(http.StatusInternalServerError, "Error when fetching the product!")
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h Product) Router(w http.ResponseWriter, r *http.Request) {
	router := router.NewRouter(h.Logger)

	router.AddRoute(`product/(?P<id>[\w|-]+)/?`, http.MethodGet, h.detail)
	router.AddRoute(`product/(?P<id>[\w|-]+)/?`, http.MethodDelete, h.remove)
	router.AddRoute(`product/(?P<id>[\w|-]+)/?`, http.MethodPatch, h.update)
	router.AddRoute(`product/?`, http.MethodGet, h.list)
	router.AddRoute(`product/?`, http.MethodPost, h.create)

	router.ServeHTTP(w, r)
}
