package product

import (
	"fmt"
	"github.com/monzo/terrors"
	"github.com/monzo/typhon"
	"github.com/uwl-team-project-2019-semester-2-team-2/online-shopping-core/database"
	"log"
	"strconv"
)

type Product struct {
	Repository *Repository
}

func Init(database database.Database, router *typhon.Router) {
	repo := Repository{&database}
	prod := Product{&repo}
	prod.Routes(router)
}

func (pr *Product) Routes(router *typhon.Router) {
	router.GET("/product/:productId", pr.Get)
}

func (pr *Product) Get(r typhon.Request) typhon.Response {
	response := typhon.NewResponse(r)

	productIdStr, ok := typhon.RouterForRequest(r).Params(r)["productId"]

	if !ok {
		response.Error = terrors.InternalService("missing_parameter", "ProductID parameter missing in request", nil)
		return response
	}

	prod, err := pr.Repository.product(productIdStr)

	if err != nil {
		response.Error = terrors.InternalService("database_error", err.Error(), nil)
		return response
	}

	stock, err := pr.Repository.stock(productIdStr)

	if err != nil {
		response.Error = terrors.InternalService("database_error", err.Error(), nil)
		return response
	}

	related, err := pr.Repository.related(productIdStr, strconv.Itoa(prod.ProductLineId))

	if err != nil {
		response.Error = terrors.InternalService("database_error", err.Error(), nil)
		return response
	}

	prod.Stock = stock
	prod.Related = related

	log.Print(fmt.Sprintf("processing get request for product %s", productIdStr))

	return r.Response(prod)
}
