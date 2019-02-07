package controller

import (
	"../common"
	"net/http"
	"github.com/go-chi/chi"
	"encoding/json"
)

type MarketingController struct {}

func (controller *MarketingController) Route() chi.Router {
	router := chi.NewRouter()
	router.Get("/", controller.index)
	return router
}

func (controller *MarketingController) index(w http.ResponseWriter, r *http.Request) {
	webpage := common.Page { "Marketing" }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(webpage)
}
