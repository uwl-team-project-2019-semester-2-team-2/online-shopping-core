package search

import (
	"fmt"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
	"log"
)

type Handler struct {
	Repository *Repository
}

func Init(database database.Database, router *typhon.Router) {
	repo := Repository{&database}
	prod := Handler{&repo}
	prod.Routes(router)
}

func (pr *Handler) Routes(router *typhon.Router) {
	router.GET("/search/:searchQuery", pr.Get)
}


func (pr *Handler) Get(r typhon.Request) typhon.Response {
	response := typhon.NewResponse(r)
	searchQuery, ok := typhon.RouterForRequest(r).Params(r)["searchQuery"]
	log.Print(fmt.Sprintf("processing get request for product %s", searchQuery))

	if !ok {
		response.Error = terrors.InternalService("missing_parameter", "ProductID parameter missing in request", nil)
		return response
	}

	searches, err := pr.Repository.search(searchQuery)

	if err != nil {
		response.Error = terrors.InternalService("Database Error", err.Error(), nil)
		return response
	}

	return r.Response(searches)
}
