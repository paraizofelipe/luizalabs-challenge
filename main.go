package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	buyer "github.com/paraizofelipe/luizalabs-challenge/buyer/handler"
	"github.com/paraizofelipe/luizalabs-challenge/config"
	product "github.com/paraizofelipe/luizalabs-challenge/product/handler"
)

func main() {
	var err error
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	storage, err := sqlx.Open("postgres", config.Storage)
	if err != nil {
		log.Panic(err)
	}

	buyerHandler := buyer.NewHandler(storage, logger)
	productHandler := product.NewHandler(storage, logger)

	http.HandleFunc("/api/buyer/", buyerHandler.Router)
	http.HandleFunc("/api/product/", productHandler.Router)

	url := fmt.Sprintf("%s:%s", config.Host, os.Getenv("PORT"))

	log.Printf("Server listening in %s", url)

	if err = http.ListenAndServe(url, nil); err != nil {
		logger.Fatal(err)
	}
}
