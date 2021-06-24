package handler

import (
	"net/http"

	"github.com/paraizofelipe/luizalabs-challenge/router"
)

type Handler interface {
	create(*router.Context)
	remove(*router.Context)
	update(*router.Context)
	detail(*router.Context)
	Router(http.ResponseWriter, *http.Request)
}
