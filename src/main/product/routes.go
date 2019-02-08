package product

import (
	"github.com/go-chi/chi"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"net/http"
	"context"
)

type ProductController struct {
	Repository *Repository
}

func Init(db *sql.DB) ProductController {
	return ProductController { NewProductRepository(db) }
}

func (controller *ProductController) Route() chi.Router {
	router := chi.NewRouter()
	
	router.Get("/brands", controller.brand)
	
	router.Route("/{productID}", func (r chi.Router) {
		r.Use(ProductContext)
		r.Get("/", controller.product)
	}) 
	
	router.Route("/search/{searchQuery}", func (r chi.Router) {
		r.Use(SearchContext)
		r.Get("/", controller.search)
	}) 
	return router
}

func ProductContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  productID := chi.URLParam(r, "productID")
	  ctx := context.WithValue(r.Context(), "productID", productID)
	  next.ServeHTTP(w, r.WithContext(ctx))
	})
  }
  
func SearchContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		searchQuery := chi.URLParam(r, "searchQuery")
	  ctx := context.WithValue(r.Context(), "searchQuery", searchQuery)
	  next.ServeHTTP(w, r.WithContext(ctx))
	})
  }
  