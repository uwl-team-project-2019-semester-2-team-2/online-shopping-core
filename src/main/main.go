package main

import (
	"./product"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	_"./models"
	"log"
	"net/http"
	"github.com/go-chi/chi"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "test:test@/shop")

	if err != nil {
		fmt.Print(err)
	}

	prd := product.Init(db)

	router := chi.NewRouter()
	router.Mount("/products", prd.Route())

	log.Fatal(http.ListenAndServe(":8080", router))
}
