package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/paraizofelipe/luizalabs-challenge/config"
	"github.com/paraizofelipe/luizalabs-challenge/router"
)

type CustomClaims struct {
	router.Authorization
	jwt.StandardClaims
}

func BasicAuth(handler router.Handler, ctx *router.Context) {
	var (
		err          error
		ok           bool
		claims       *CustomClaims
		mySigningKey = []byte(config.Secret)
	)

	ctx.ResponseWriter.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

	s := strings.SplitN(ctx.Request.Header.Get("Authorization"), " ", 2)
	if len(s) != 1 {
		http.Error(ctx.ResponseWriter, "Not authorized!", http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(s[0], &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		http.Error(ctx.ResponseWriter, "Not authorized!", http.StatusUnauthorized)
		return
	}

	if claims, ok = token.Claims.(*CustomClaims); !ok && !token.Valid {
		http.Error(ctx.ResponseWriter, "Not authorized!", http.StatusUnauthorized)
		return
	}

	if !claims.Read && !claims.Write {
		http.Error(ctx.ResponseWriter, "Not authorized!", http.StatusUnauthorized)
		return
	}

	ctx.Authorization = claims.Authorization

	handler(ctx)
}
