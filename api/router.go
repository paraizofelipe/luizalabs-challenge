package api

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/paraizofelipe/luizalabs-challenge/config"
)

type Route struct {
	Pattern        string
	ActionHandlers map[string]http.Handler
}

type Router struct {
	http.Handler
	routes []Route
	logger *log.Logger
	debug  bool
}

func NewRouter(logger *log.Logger) *Router {
	logger.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	return &Router{
		routes: make([]Route, 0),
		logger: logger,
		debug:  config.Debug,
	}
}

// AddRoute ---
func (rt *Router) AddRoute(pattern string, method string, handler http.Handler) {
	var found = false
	for _, route := range rt.routes {
		if route.Pattern == pattern {
			found = true
			route.ActionHandlers[method] = handler
		}
	}

	if !found {
		rt.routes = append(rt.routes, Route{
			Pattern: pattern,
			ActionHandlers: map[string]http.Handler{
				method: handler,
			},
		})
	}
}

// ServerHTTP ---
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range rt.routes {
		if matched, _ := regexp.MatchString(route.Pattern, r.URL.Path); matched {
			if h, registered := route.ActionHandlers[r.Method]; registered {
				if rt.debug {
					rt.trace(r)
				}
				h.ServeHTTP(w, r)
			} else {
				http.NotFound(w, r)
			}
			return
		}
	}
}

// trace ---
func (rt *Router) trace(req *http.Request) {
	debugLine := fmt.Sprintf("%v %v %v", req.RemoteAddr, req.Method, req.URL.Path)
	rt.logger.Println(debugLine)
}
