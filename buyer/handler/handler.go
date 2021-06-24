package handler

import (
	"encoding/json"
	"net/http"

	"github.com/paraizofelipe/luizalabs-challenge/router"
)

type Handler interface {
	create(*router.Context)
	remove(*router.Context)
	update(*router.Context)
	detail(*router.Context)
	addFavoriteProduct(*router.Context)
	Router(http.ResponseWriter, *http.Request)
}

type ErrorResponse struct {
	Error string `json:"errors"`
}

type SuccessResponse struct {
	Msg string `json:"msg"`
}

func (r ErrorResponse) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}

func (r SuccessResponse) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}
