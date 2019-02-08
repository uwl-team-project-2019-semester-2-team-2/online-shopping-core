package product

import (
	"net/http"
	"encoding/json"
)

// type ProductController struct {
// 	Repository *ProductRepository
	
// }

func (controller *ProductController) product(w http.ResponseWriter, r *http.Request) {
	productID := r.Context().Value("productID")

	prod := controller.Repository.product(productID.(string))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	json.NewEncoder(w).Encode(prod)
}

func (controller *ProductController) brand(w http.ResponseWriter, r *http.Request) {
	brands := controller.Repository.Brands()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	
	json.NewEncoder(w).Encode(brands)
}