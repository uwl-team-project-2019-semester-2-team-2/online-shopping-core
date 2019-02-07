package controller

import (
	"github.com/go-chi/chi"
	"../product"
)

func Init(router chi.Router) {
	router.Mount("/products", new(product.ProductController).Route())
	router.Mount("/marketing", new(MarketingController).Route())
}