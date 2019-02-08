package product

import (
	"net/http"
)

func product(w http.ResponseWriter, r *http.Request) {
	//productID := r.Context().Value("productID")

	/*
		prod, err := NewProductRepository.product(productID.(string))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err != nil {
			panic(err)
			json.NewEncoder(w).Encode(`error`)
		} else {
			json.NewEncoder(w).Encode(prod)
		}
	*/
}

func brand(w http.ResponseWriter, r *http.Request) {
	/*brands := controller.Repository.Brands()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(brands)*/
}

func search(w http.ResponseWriter, r *http.Request) {
	/*searchQuery := r.Context().Value("searchQuery")

	brands, err := controller.Repository.Search(searchQuery.(string))

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(brands)*/
}
