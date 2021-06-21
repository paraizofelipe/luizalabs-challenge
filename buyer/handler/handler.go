package handler

import (
	"encoding/json"
)

type Handler struct {
	Buyer Buyer
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
