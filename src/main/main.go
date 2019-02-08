package main

import (
	_ "./models"
	"./product"
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	mysqlConfig := mysql.NewConfig()
	mysqlConfig.User = ""
	mysqlConfig.Passwd = ""
	mysqlConfig.Addr = "127.0.0.1"

	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	prd := product.Init(db)

	router := chi.NewRouter()
	router.Mount("/products", prd.Route())

	log.Fatal(http.ListenAndServe(":8080", router))
}
