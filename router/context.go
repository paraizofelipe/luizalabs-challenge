package router

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Context struct {
	http.ResponseWriter
	*http.Request
	Authorization Authorization
	QueryString   url.Values
	Params        map[string]string
}

func (c *Context) Text(code int, body string) {
	var err error

	c.ResponseWriter.Header().Set("Content-Type", "text/plain")
	c.WriteHeader(code)

	if _, err = io.WriteString(c.ResponseWriter, body); err != nil {
		panic(err)
	}
}

func (c *Context) JSON(code int, body interface{}) {
	var (
		err error
		j   []byte
	)

	if j, err = json.Marshal(body); err != nil {
		c.ResponseWriter.Header().Set("Content-Type", "application/json")
		c.WriteHeader(http.StatusInternalServerError)
		_, err = io.WriteString(c.ResponseWriter, ErrorResponse{Error: "error converting response to JSON!"}.String())
		if err != nil {
			panic(err)
		}
		return
	}

	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.WriteHeader(code)
	_, err = io.WriteString(c.ResponseWriter, string(j))
	if err != nil {
		panic(err)
	}
}
