package router

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/paraizofelipe/luizalabs-challenge/config"
)

type Handler func(*Context)

type Route struct {
	Pattern       *regexp.Regexp
	Method        string
	ActionHandler Handler
}

type Router struct {
	http.Handler
	debug        bool
	logger       *log.Logger
	Routes       []Route
	DefaultRoute Handler
}

func NewRouter(logger *log.Logger) *Router {
	logger.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	return &Router{
		Routes: make([]Route, 0),
		logger: logger,
		debug:  config.Debug,
		DefaultRoute: func(ctx *Context) {
			ctx.Text(http.StatusNotFound, "URL Not found!")
		},
	}
}

// AddRoute ---
func (r *Router) AddRoute(pattern string, method string, handler Handler) {
	re := regexp.MustCompile(pattern)
	route := Route{Pattern: re, Method: method, ActionHandler: handler}

	r.Routes = append(r.Routes, route)
}

func ParsePathParams(matches []string, re *regexp.Regexp) (params map[string]string) {
	params = make(map[string]string)
	for index, key := range re.SubexpNames() {
		if key == "" {
			continue
		}
		params[key] = matches[index]
	}
	return
}

// ServerHTTP ---
func (r *Router) ServeHTTP(w http.ResponseWriter, resp *http.Request) {
	ctx := &Context{Request: resp, ResponseWriter: w}

	for _, rt := range r.Routes {
		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 && rt.Method == ctx.Method {
			ctx.Params = ParsePathParams(matches, rt.Pattern)
			if ctx.URL.Query() != nil {
				ctx.QueryString = ctx.URL.Query()
			}
			if r.debug {
				r.trace(resp)
			}
			rt.ActionHandler(ctx)
			return
		}
	}

	r.DefaultRoute(ctx)
}

// trace ---
func (r *Router) trace(req *http.Request) {
	debugLine := fmt.Sprintf("%v %v %v", req.RemoteAddr, req.Method, req.URL.Path)
	r.logger.Println(debugLine)
}
