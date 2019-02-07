package controller

import (
	_"../common"
	"net/http"
	"github.com/go-chi/chi"
	"encoding/json"
)

type ProductController struct {}

type Product struct {
	Id int 				`json:"id"`
	Name string			`json:"name"`
	Brand string		`json:"brand"`
	Colour string		`json:"colour"`
	Description string	`json:"description"`
	Sizes []Size		`json:"sizes"`
	Pictures []string	`json:"pictures"`
	Related []Related	`json:"related"`
}

type Brand struct {
	Id int				`json:"id"`
	Name string			`json:"name"`
}

type Related struct {
	Id int				`json:"id"`
	Title string		`json:"title"`
}

type Size struct {
	Size string			`json:"size"`
	InStock bool		`json:"in_stock"`
}

func (controller *ProductController) Route() chi.Router {
	router := chi.NewRouter()
	router.Get("/", controller.index)
	return router
}

func (controller *ProductController) index(w http.ResponseWriter, r *http.Request) {
	products := []Product {
		Product {
			Id: 0,
			Name: "Nike Air Max 90",
			Brand: "Nike",
			Colour: "Black",
			Description: "The Nike Air Max 90 released in 1990 and is considered to be the second flagship sneaker of the Air Max legacy. The most popular colorway is the 'Infrared', which was the original colorway in 1990. The Air Max 90 was designed by legendary Nike architect Tinker Hatfield.",
			Sizes: []Size {
				Size {
					Size: "6",
					InStock: true,
				},
			},
			Pictures: []string {
				"pic.jpg",
				"pic2.jpg",
			},
			Related: []Related {
				Related {
					Id: 1,
					Title: "White",
				},
			},
		},
	}


	// webpage := common.Page { "API v1" }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	json.NewEncoder(w).Encode(products)
}
