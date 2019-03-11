package search

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
	"strings"
)

var decoder = schema.NewDecoder()

type Search struct {
	Repository *Repository
}

func Init(database database.Database, router *typhon.Router) {
	repo := Repository{&database}
	prod := Search{&repo}
	prod.Routes(router)
}

func (pr *Search) Routes(router *typhon.Router) {
	router.GET("/search/:searchQuery", pr.Get)
}

func (pr *Search) Get(r typhon.Request) typhon.Response {
	response := typhon.NewResponse(r)
	searchQuery, ok := typhon.RouterForRequest(r).Params(r)["searchQuery"]

	if !ok {
		response.Error = terrors.InternalService("missing_parameter", "ProductID parameter missing in request", nil)
		return response
	}

	var urlParams URLParams
	err := decoder.Decode(&urlParams, r.URL.Query())

	if urlParams.Page == 0 {
		urlParams.Page = 1
	}

	filters, err := pr.Repository.filters()

	if err != nil {
		response.Error = terrors.InternalService("database_error", err.Error(), nil)
		return response
	}

	var validFilters []string

	if urlParams.Filter != "" {
		userFilters := strings.Split(urlParams.Filter, ",")
		for _, userFilter := range userFilters {
			for _, filter := range filters {
				if userFilter == filter.URL {
					validFilters = append(validFilters, filter.URL)
				}
			}
		}
	}

	if urlParams.Order != "" {
		switch order := urlParams.Order; order {
		case "relevance":
			fmt.Println(order)
		case "ascending":
			fmt.Println(order)
		case "descending":
			fmt.Println(order)
		}
	}

	searches, err := pr.Repository.search(searchQuery, urlParams.Page, validFilters...)

	if err != nil {
		response.Error = terrors.InternalService("database_error", err.Error(), nil)
		return response
	}

	count, err := pr.Repository.count(searchQuery)

	if err != nil {
		response.Error = terrors.InternalService("database_error", err.Error(), nil)
		return response
	}

	var marshaller = Marshaller{
		Page: urlParams.Page,
		Count: count,
		SearchProducts: searches,
		Filters: filters,
	}

	return r.Response(marshaller)
}
